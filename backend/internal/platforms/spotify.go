package platforms

import (
    "backend/internal/models"
    "backend/internal/repositories"
    "context"

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

func GetSpotifySong(userId *uuid.UUID, songId string) (models.Song, error) {
    song, err := repositories.GetSong("spotify", songId)
    if err != nil {
        return song, err
    }

    spotifyId := spotify.ID(songId)
    client, err := SpotifyClientId(userId)

    if err != nil {
        return song, err
    }

    apiSong, err := client.GetTrack(context.Background(), spotifyId)

    if err != nil {
        return song, err
    }


    artists := make([]models.Artist, 0)
    for _, artist := range apiSong.Artists {
        artists = append(artists, models.Artist{
            Name: artist.Name,
            ID: uuid.New(),
        })
    }

    song = models.Song{
        Title: apiSong.Name,
        ID: uuid.New(),
        Artists: artists,
        MediaURL: apiSong.Album.Images[0].URL,
        PreviewURL: apiSong.PreviewURL,
    }

    return song, nil
}
