package handlers

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"

    "github.com/zmb3/spotify"
)

const (
    REDIRECT = "http://127.0.0.1/redirect"
)

type Url struct {
    Url string `json:"url"`
}

func OauthSpotify(w http.ResponseWriter, r *http.Request) {
    // TODO: Get CSRF Token
    csrf := "abc123"
    // TODO: Proper redirect
    auth := spotify.NewAuthenticator(
        REDIRECT,
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
    url := &Url{
        Url: auth.AuthURL(csrf),
    }
    json, err := json.Marshal(url)
    if err != nil {
        fmt.Printf("Error: %s", err.Error())
    }
    fmt.Fprintf(w, "%s", json)
}

func OauthDeezer(w http.ResponseWriter, r *http.Request) {
    perms := "basic_access,email,offline_access,manage_library,manage_community,delete_library,listening_history"
    id := os.Getenv("DEEZER_CLIENT_ID")
    url := &Url{
        Url: "https://connect.deezer.com/oauth/auth.php?app_id=" + id + "&redirect_uri=" + REDIRECT + "&perms=" + perms,
    }
    json, err := json.Marshal(url)
    if err != nil {
        fmt.Printf("Error: %s", err.Error())
    }
    fmt.Fprintf(w, "%s", json)
}
