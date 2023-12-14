package handlers

import (
    "backend/internal/scanner"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "os"
    "time"

    "github.com/markbates/goth/providers/deezer"
    "github.com/zmb3/spotify"
)

const (
    SPOTIFY_REDIRECT = "http://127.0.0.1:8080/api/oauth/callback/spotify"
    DEEZER_REDIRECT  = "http://127.0.0.1:8080/api/oauth/callback/deezer"
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
    // TODO: Get CSRF Token
    csrf := "abc123"
    auth := GetSpotifyAuthenticator(csrf)
    url := &Url{
        Url: auth.AuthURL(csrf),
    }
    json, err := json.Marshal(url)
    if err != nil {
        fmt.Printf("Error: %s", err.Error())
    }
    fmt.Fprintf(w, "%s", json)
}

func GetSpotifyAuthenticator(csrf string) spotify.Authenticator {
    auth := spotify.NewAuthenticator(
        SPOTIFY_REDIRECT,
        spotify.ScopeImageUpload,
        spotify.ScopePlaylistReadPrivate,
        spotify.ScopePlaylistModifyPublic,
        spotify.ScopePlaylistModifyPrivate,
        spotify.ScopePlaylistReadCollaborative,
        spotify.ScopeUserFollowModify,
        spotify.ScopeUserFollowRead,
        spotify.ScopeUserLibraryModify,
        spotify.ScopeUserLibraryRead,
        spotify.ScopeUserReadPrivate,
        spotify.ScopeUserReadEmail,
        spotify.ScopeUserReadCurrentlyPlaying,
        spotify.ScopeUserReadPlaybackState,
        spotify.ScopeUserModifyPlaybackState,
        spotify.ScopeUserReadRecentlyPlayed,
        spotify.ScopeUserTopRead,
        spotify.ScopeStreaming,
    )
    id := os.Getenv("SPOTIFY_CLIENT_ID")
    secret := os.Getenv("SPOTIFY_SECRET")
    auth.SetAuthInfo(id, secret)
    return auth
}

func SpotifyCallback(w http.ResponseWriter, r *http.Request) {
    // TODO: Get CSRF Token
    csrf := "abc123"
    auth := GetSpotifyAuthenticator(csrf)
    token, err := auth.Token(csrf, r)
    if err != nil {
        http.Error(w, "Couldn't get token", http.StatusNotFound)
        return
    }
    // TODO: Save these
    fmt.Fprintf(w, "AccessToken: %s\n", token.AccessToken)
    fmt.Fprintf(w, "RefreshToken: %s", token.RefreshToken)
    client := auth.NewClient(token)
    go scanner.ScanSpotify(client)
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
    var user DeezerUser
    json.Unmarshal(body2, &user)
    if err != nil {
        fmt.Fprintf(w, "%s", err)
        return
    }
    go scanner.ScanDeezer(user.ID)
}

type DeezerUser struct {
    ID int `json:"id"`
}
