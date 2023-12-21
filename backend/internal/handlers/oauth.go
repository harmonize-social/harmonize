package handlers

import (
    "backend/internal/models"
    "backend/internal/repositories"
    "backend/internal/scanner"
    "context"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "os"
    "time"

    "github.com/google/uuid"
    "github.com/markbates/goth/providers/deezer"
    deezer2 "github.com/stayradiated/deezer"
    "github.com/zmb3/spotify/v2"
    "github.com/zmb3/spotify/v2/auth"
)

const (
    SPOTIFY_REDIRECT = "http://127.0.0.1:8080/api/v1/oauth/callback/spotify"
    DEEZER_REDIRECT  = "http://127.0.0.1:8080/api/v1/oauth/callback/deezer"
    TEST_SESSION     = "df8d8816-3280-41aa-9e27-ec60ba297c9e"
)

type Url struct {
    Url string `json:"url"`
}

type Tokens struct {
    AccessToken  string `json:"accessToken"`
    RefreshToken string `json:"refreshToken"`
}

func DeezerProvider() *deezer.Provider {
    id := os.Getenv("DEEZER_CLIENT_ID")
    secret := os.Getenv("DEEZER_SECRET")
    provider := deezer.New(id, secret, DEEZER_REDIRECT, "basic_access", "email", "offline_access", "manage_library", "manage_community", "delete_library", "listening_history")
    return provider
}

func DeezerURL(w http.ResponseWriter, r *http.Request) {
    provider := DeezerProvider()
    session, err := provider.BeginAuth("abc123")
    if err != nil {
        fmt.Printf("Error: %s", err.Error())
    }
    url, err := session.GetAuthURL()
    if err != nil {
        fmt.Printf("Error: %s", err.Error())
    }
    urlStruct := &Url{
        Url: url,
    }
    json, err := json.Marshal(urlStruct)
    if err != nil {
        fmt.Printf("Error: %s", err.Error())
    }
    fmt.Fprintf(w, "%s", json)
}

func SpotifyURL(w http.ResponseWriter, r *http.Request) {
    state := r.Header.Get("id")
    auth := GetSpotifyAuthenticator(state)
    url := &Url{
        Url: auth.AuthURL(state),
    }
    json, err := json.Marshal(url)
    if err != nil {
        fmt.Printf("Error: %s", err.Error())
    }
    fmt.Fprintf(w, "%s", json)
}

func GetSpotifyAuthenticator(csrf string) spotifyauth.Authenticator {
    auth := spotifyauth.New(
        spotifyauth.WithRedirectURL(SPOTIFY_REDIRECT),
        spotifyauth.WithScopes(
            spotifyauth.ScopeImageUpload,
            spotifyauth.ScopePlaylistReadPrivate,
            spotifyauth.ScopePlaylistModifyPublic,
            spotifyauth.ScopePlaylistModifyPrivate,
            spotifyauth.ScopePlaylistReadCollaborative,
            spotifyauth.ScopeUserFollowModify,
            spotifyauth.ScopeUserFollowRead,
            spotifyauth.ScopeUserLibraryModify,
            spotifyauth.ScopeUserLibraryRead,
            spotifyauth.ScopeUserReadPrivate,
            spotifyauth.ScopeUserReadEmail,
            spotifyauth.ScopeUserReadCurrentlyPlaying,
            spotifyauth.ScopeUserReadPlaybackState,
            spotifyauth.ScopeUserModifyPlaybackState,
            spotifyauth.ScopeUserReadRecentlyPlayed,
            spotifyauth.ScopeUserTopRead,
            spotifyauth.ScopeStreaming,
        ),
        spotifyauth.WithClientID(os.Getenv("SPOTIFY_CLIENT_ID")),
        spotifyauth.WithClientSecret(os.Getenv("SPOTIFY_SECRET")),
    )
    return *auth
}

func SpotifyCallback(w http.ResponseWriter, r *http.Request) {
    session := r.URL.Query().Get("state")
    user, err := getUserFromSession(uuid.MustParse(session))
    if err != nil {
        fmt.Printf("Error: %s", err.Error())
        models.Error(w, http.StatusUnauthorized, "Invalid token")
        return
    }
    auth := GetSpotifyAuthenticator(session)
    token, err := auth.Token(r.Context(), session, r)
    if err != nil {
        http.Error(w, "Couldn't get token", http.StatusNotFound)
        fmt.Println(err)
        return
    }
    client := spotify.New(auth.Client(r.Context(), token), spotify.WithRetry(true))
    connection := &models.Connection{
        ID:           uuid.New(),
        UserID:       user.ID,
        AccessToken:  token.AccessToken,
        RefreshToken: token.RefreshToken,
        Expiry:       token.Expiry,
    }
    sqlStatement := `INSERT INTO connections (id, user_id, access_token, refresh_token, expiry) VALUES ($1, $2, $3, $4, $5) RETURNING id`
    var connectionID uuid.UUID
    err2 := repositories.Pool.QueryRow(context.Background(),
        sqlStatement,
        connection.ID,
        connection.UserID,
        connection.AccessToken,
        connection.RefreshToken,
        connection.Expiry.Format(time.RFC3339)).Scan(&connectionID)
    if err2 != nil {
        fmt.Fprintf(w, "connection: %s\n\r", connection.UserID)
        fmt.Fprintf(w, "Unable to execute the query. %s", err2)
    }
    go scanner.Spotify(*client, connectionID)
}

type DeezerAccessToken struct {
    AccessToken string `json:"access_token"`
}

func DeezerCallback(w http.ResponseWriter, r *http.Request) {
    provider := DeezerProvider()
    code := r.URL.Query().Get("code")
    url := "https://connect.deezer.com/oauth/access_token.php?app_id=" + provider.ClientKey + "&secret=" + provider.Secret + "&code=" + code + "&output=json"
    response, err := provider.Client().Get(url)
    if err != nil {
        fmt.Fprintf(w, "err: %s", err)
        return
    }
    defer response.Body.Close()
    body, err := io.ReadAll(response.Body)
    var deezerToken DeezerAccessToken
    json.Unmarshal(body, &deezerToken)
    expiresAt := time.Now().Add(time.Hour * 24 * 365 * 100)
    sessionActual := &deezer.Session{
        AuthURL:     "",
        AccessToken: deezerToken.AccessToken,
        ExpiresAt:   expiresAt,
    }
    response2, err := provider.Client().Get("https://api.deezer.com/user/me?access_token=" + sessionActual.AccessToken)
    if err != nil {
        fmt.Fprintf(w, "err: %s", err)
        return
    }
    defer response.Body.Close()
    body2, err := io.ReadAll(response2.Body)
    if err != nil {
        fmt.Fprintf(w, "%s", err)
        return
    }
    var user deezer2.User
    json.Unmarshal(body2, &user)
    if err != nil {
        fmt.Fprintf(w, "%s", err)
        return
    }
    go scanner.ScanDeezer(user.ID)
}
