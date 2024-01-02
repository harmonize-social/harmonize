package platforms

import (
    "backend/internal/auth"
    "backend/internal/models"
    "backend/internal/repositories"
    "context"
    "fmt"
    "net/http"

    "github.com/google/uuid"
    spotify "github.com/zmb3/spotify/v2"
    spotifyauth "github.com/zmb3/spotify/v2/auth"
    "golang.org/x/oauth2"
)

func SpotifyClientId(userId *uuid.UUID) (*spotify.Client, error) {
    var token oauth2.Token
    err := repositories.Pool.QueryRow(context.Background(), "SELECT access_token, refresh_token, expiry FROM connections JOIN libraries ON connections.id = libraries.connection_id WHERE user_id = $1 AND platform_id = $2", userId, "spotify").Scan(&token.AccessToken, &token.RefreshToken, &token.Expiry)
    if err != nil {
        return nil, err
    }
    auth := spotifyauth.New(
        spotifyauth.WithClientID(repositories.SpotifyClientId),
        spotifyauth.WithClientSecret(repositories.SpotifySecret),
    )
    newToken, err := auth.RefreshToken(context.Background(), &token)
    if err != nil {
        return nil, err
    }
    newAuth := spotifyauth.New()
    httpClient := newAuth.Client(context.Background(), newToken)
    client := spotify.New(httpClient)

    return client, nil
}

func SpotifyURL(csrf string) (string, error) {
    url := GetSpotifyAuthenticator(csrf).AuthURL(csrf)
    return url, nil
}

func GetSpotifyAuthenticator(csrf string) spotifyauth.Authenticator {
    auth := spotifyauth.New(
        spotifyauth.WithRedirectURL(repositories.SpotifyRedirect),
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
        spotifyauth.WithClientID(repositories.SpotifyClientId),
        spotifyauth.WithClientSecret(repositories.SpotifySecret),
    )
    return *auth
}

func SpotifyCallback(w http.ResponseWriter, r *http.Request) {
    session := r.URL.Query().Get("state")
    user, err := auth.GetUserFromSession(uuid.MustParse(session))
    if err != nil {
        fmt.Println(err)
        models.Error(w, http.StatusInternalServerError, "Internal server error")
        return
    }
    auth := GetSpotifyAuthenticator(session)
    token, err := auth.Token(r.Context(), session, r)
    if err != nil {
        fmt.Println(err)
        models.Error(w, http.StatusInternalServerError, "Internal server error")
        return
    }
    err = repositories.CreateConnectionAndLibrary(user.ID, "spotify", token.AccessToken, token.RefreshToken, token.Expiry)
    if err != nil {
        fmt.Println(err)
        models.Error(w, http.StatusInternalServerError, "Internal server error")
        return
    }
    models.Result(w, "Ok")
}
