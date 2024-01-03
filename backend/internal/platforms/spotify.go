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

func GetSpotifySongs(userId *uuid.UUID, limit int, offset int) ([]models.PlatformSong, error) {
    client, err := SpotifyClientId(userId)
    if err != nil {
        return nil, err
    }
    tracks, err := client.CurrentUsersTracks(context.Background(), spotify.Limit(limit), spotify.Offset(offset))
    if err != nil {
        return nil, err
    }

    songs := make([]models.PlatformSong, len(tracks.Tracks))
    for i, track := range tracks.Tracks {
        artists := make([]models.PlatformArtist, len(track.Artists))
        for j, artist := range track.Artists {
            artists[j] = models.PlatformArtist{
                Platform: "spotify",
                ID:       artist.ID.String(),
                Name:     artist.Name,
                MediaURL: "",
            }
        }
        songs[i] = models.PlatformSong{
            Platform: "spotify",
            ID:       track.ID.String(),
            Title:    track.Name,
            Album: models.PlatformAlbum{
                Platform: "spotify",
                ID:       track.Album.ID.String(),
                Title:    track.Album.Name,
                Artists:  artists,
                MediaURL: track.Album.Images[0].URL,
            },
            MediaURL:   track.Album.Images[0].URL,
            PreviewURL: track.PreviewURL,
        }
    }
    return songs, nil
}

func GetSpotifyArtists(userId *uuid.UUID, limit int, offset int) ([]models.PlatformArtist, error) {
    client, err := SpotifyClientId(userId)
    if err != nil {
        return nil, err
    }
    artistsPage, err := client.CurrentUsersFollowedArtists(context.Background(), spotify.Limit(limit), spotify.Offset(offset))
    if err != nil {
        return nil, err
    }

    artists := make([]models.PlatformArtist, 0)
    for _, artist := range artistsPage.Artists {
        fullArtist, err := client.GetArtist(context.Background(), artist.ID)

        if err != nil {
            return nil, err
        }

        platformArtist := models.PlatformArtist{
            Platform: "spotify",
            ID:       fullArtist.ID.String(),
            Name:     fullArtist.Name,
            MediaURL: fullArtist.Images[0].URL,
        }

        artists = append(artists, platformArtist)
    }
    return artists, nil
}

func GetSpotifyAlbums(userId *uuid.UUID, limit int, offset int) ([]models.PlatformAlbum, error) {
    client, err := SpotifyClientId(userId)
    if err != nil {
        return nil, err
    }
    albumsPage, err := client.CurrentUsersAlbums(context.Background(), spotify.Limit(limit), spotify.Offset(offset))
    if err != nil {
        return nil, err
    }

    albums := make([]models.PlatformAlbum, len(albumsPage.Albums))
    for i, album := range albumsPage.Albums {
        fullAlbum, err := client.GetAlbum(context.Background(), album.ID)
        if err != nil {
            return nil, err
        }
        artists := make([]models.PlatformArtist, len(album.Artists))
        for j, artist := range album.Artists {
            artists[j] = models.PlatformArtist{
                Platform: "spotify",
                ID:       artist.ID.String(),
                Name:     artist.Name,
                MediaURL: "",
            }
        }
        songs := make([]models.PlatformSong, len(fullAlbum.Tracks.Tracks))
        for j, track := range fullAlbum.Tracks.Tracks {
            songs[j] = models.PlatformSong{
                Platform:   "spotify",
                ID:         track.ID.String(),
                Title:      track.Name,
                PreviewURL: track.PreviewURL,
            }
        }

        albums[i] = models.PlatformAlbum{
            Platform: "spotify",
            ID:       album.ID.String(),
            Title:    album.Name,
            Artists:  artists,
            Songs:    songs,
            MediaURL: album.Images[0].URL,
        }
    }
    return albums, nil
}
