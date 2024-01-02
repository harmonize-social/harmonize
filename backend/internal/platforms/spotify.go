package platforms

import (
    "backend/internal/models"
    "backend/internal/repositories"
    "backend/internal/auth"
    "context"
    "fmt"
    "net/http"
    "time"

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
    sqlStatement = `
    INSERT INTO libraries (platform_id, id, connection_id) VALUES ('spotify', uuid_generate_v4(), $1) RETURNING id;
    `
    tag, err := repositories.Pool.Exec(context.Background(),
        sqlStatement,
        connectionID)
    if err != nil {
        fmt.Printf("error: %v", err)
        models.Error(w, http.StatusInternalServerError, "Internal server error")
    }

    if tag.RowsAffected() == 0 {
        models.Error(w, http.StatusInternalServerError, "Internal server error")
        return
    }
    models.Result(w, "Ok")
}
