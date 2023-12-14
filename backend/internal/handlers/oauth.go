package handlers

import (
    "backend/internal/scanner"
    "encoding/json"
    "fmt"
    "net/http"
    "os"

    "github.com/zmb3/spotify"
)

const (
    SPOTIFY_REDIRECT = "http://127.0.0.1:8080/api/oauth/callback/spotify"
    DEEZER_REDIRECT = "http://127.0.0.1:8080/api/oauth/callback/deezer"
)

type Url struct {
    Url string `json:"url"`
}

type Tokens struct {
    AccessToken  string `json:"accessToken"`
    RefreshToken string `json:"refreshToken"`
}

func DeezerURL(w http.ResponseWriter, r *http.Request) {
    perms := "basic_access,email,offline_access,manage_library,manage_community,delete_library,listening_history"
    id := os.Getenv("DEEZER_CLIENT_ID")
    url := &Url{
        Url: "https://connect.deezer.com/oauth/auth.php?app_id=" + id + "&redirect_uri=" + DEEZER_REDIRECT + "&perms=" + perms,
    }
    json, err := json.Marshal(url)
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

func DeezerCallback(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "URL: %s\n", r.URL)
}
