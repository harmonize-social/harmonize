package handlers

import (
    "backend/internal/models"
    "backend/internal/repositories"
    "context"
    "fmt"
    "net/http"
    "strconv"

    "github.com/google/uuid"
    "github.com/gorilla/mux"
    "github.com/zmb3/spotify/v2"
    spotifyauth "github.com/zmb3/spotify/v2/auth"
    "go.uber.org/ratelimit"
    "golang.org/x/oauth2"
)

func SongsHandler(w http.ResponseWriter, r *http.Request) {

    setCommonHeaders(w)
    setAdditionalHeaders(w, "GET")

    limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
    if err != nil {
        limit = 10
    }
    offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
    if err != nil {
        offset = 0
    }
    id := uuid.MustParse(r.Header.Get("id"))
    user, err := getUserFromSession(id)
    if err != nil {
        models.Error(w, http.StatusUnauthorized, "Malformed session")
        return
    }
    client, err := SpotifyClientFromRequest(r, &user.ID)
    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Try logging into service again")
    }
    tracks, err := client.CurrentUsersTracks(context.Background(), spotify.Limit(limit), spotify.Offset(offset))

    independentTracks := make([]models.Song, len(tracks.Tracks))
    for i, track := range tracks.Tracks {
        artists := make([]models.Artist, len(track.Artists))
        for j, artist := range track.Artists {
            artists[j] = models.Artist{
                ID:   uuid.New(),
                Name: artist.Name,
            }
        }
        independentTracks[i] = models.Song{
            ID:       uuid.New(),
            Title:    track.Name,
            Artists:  artists,
            MediaURL: track.Album.Images[0].URL,
        }
    }

    models.Result(w, independentTracks)

    /*
       CREATE TABLE IF NOT EXISTS connections(
           id UUID PRIMARY KEY,
           user_id UUID REFERENCES users (id) NOT NULL,
           access_token VARCHAR(1024) NOT NULL,
           refresh_token VARCHAR(1024) NOT NULL,
           expiry timestamptz NOT NULL
       );


       CREATE TABLE IF NOT EXISTS libraries(
           id UUID PRIMARY KEY,
           platform_id VARCHAR(1024) NOT NULL,
           connection_id UUID REFERENCES connections (id) NOT NULL
       );

    */
    // models.Result(w, param)
}

func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
    setCommonHeaders(w)
    setAdditionalHeaders(w, "GET")

    limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
    if err != nil {
        limit = 10
    }
    offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
    if err != nil {
        offset = 0
    }
    id := uuid.MustParse(r.Header.Get("id"))
    user, err := getUserFromSession(id)
    if err != nil {
        models.Error(w, http.StatusUnauthorized, "Malformed session")
        return
    }
    client, err := SpotifyClientFromRequest(r, &user.ID)
    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Try logging into service again")
    }
    artists, err := client.CurrentUsersTopArtists(context.Background(), spotify.Limit(limit), spotify.Offset(offset))
    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Try logging into service again")
    }

    independentArtists := make([]models.Artist, len(artists.Artists))
    for i, artist := range artists.Artists {
        independentArtists[i] = models.Artist{
            ID:       uuid.New(),
            Name:     artist.Name,
            MediaURL: artist.Images[0].URL,
        }
    }

    models.Result(w, independentArtists)
}

func AlbumHandler(w http.ResponseWriter, r *http.Request) {
    setCommonHeaders(w)
    setAdditionalHeaders(w, "GET")

    limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
    if err != nil {
        limit = 10
    }
    offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
    if err != nil {
        offset = 0
    }
    id := uuid.MustParse(r.Header.Get("id"))
    user, err := getUserFromSession(id)
    if err != nil {
        models.Error(w, http.StatusUnauthorized, "Malformed session")
        return
    }
    client, err := SpotifyClientFromRequest(r, &user.ID)
    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Try logging into service again")
    }

    rl := ratelimit.New(2)
    rl.Take()

    albums, err := client.CurrentUsersAlbums(context.Background(), spotify.Limit(limit), spotify.Offset(offset))
    fmt.Println(len(albums.Albums))
    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Try logging into service again")
    }

    var albumIds []spotify.ID
    for _, album := range albums.Albums {
        albumIds = append(albumIds, album.ID)
    }

    albumTracks, err := client.GetAlbums(context.Background(), albumIds, spotify.Limit(limit))

    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Try logging into service again")
    }

    independentAlbums := make([]models.Album, len(albums.Albums))
    for i, album := range albumTracks {
        songs := make([]models.Song, 0)
        for _, track := range album.Tracks.Tracks {
            artists := make([]models.Artist, len(track.Artists))
            for j, artist := range track.Artists {
                artists[j] = models.Artist{
                    ID:   uuid.New(),
                    Name: artist.Name,
                }
            }
            songs = append(songs, models.Song{
                ID:       uuid.New(),
                Title:    track.Name,
                Artists:  artists,
                MediaURL: track.Album.Images[0].URL,
            })
        }
        independentAlbums[i] = models.Album{
            ID:       uuid.New(),
            Title:    album.Name,
            MediaURL: album.Images[0].URL,
            Songs:    songs,
        }
    }

}

func PlaylistHandler(w http.ResponseWriter, r *http.Request) {
    setCommonHeaders(w)
    setAdditionalHeaders(w, "GET")

    limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
    if err != nil {
        limit = 10
    }
    offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
    if err != nil {
        offset = 0
    }
    id := uuid.MustParse(r.Header.Get("id"))
    user, err := getUserFromSession(id)
    if err != nil {
        models.Error(w, http.StatusUnauthorized, "Malformed session")
        return
    }
    client, err := SpotifyClientFromRequest(r, &user.ID)
    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Try logging into service again")
    }

    rl := ratelimit.New(2)
    rl.Take()
    playlists, err := client.CurrentUsersPlaylists(context.Background(), spotify.Limit(limit), spotify.Offset(offset))
    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Try logging into service again")
    }

    independentPlaylists := make([]models.Playlist, len(playlists.Playlists))
    for i, playlist := range playlists.Playlists {
        rl.Take()
        playlistTracks, err := client.GetPlaylistItems(context.Background(), playlist.ID, spotify.Limit(limit))
        if err != nil {
            models.Error(w, http.StatusInternalServerError, "Try logging into service again")
        }
        songs := make([]models.Song, 0)
        for _, track := range playlistTracks.Items {
            if track.Track.Track == nil {
                continue
            }
            artists := make([]models.Artist, len(track.Track.Track.Artists))
            for j, artist := range track.Track.Track.Artists {
                artists[j] = models.Artist{
                    ID:   uuid.New(),
                    Name: artist.Name,
                }
            }
            songs = append(songs, models.Song{
                ID:       uuid.New(),
                Title:    track.Track.Track.Name,
                Artists:  artists,
                MediaURL: track.Track.Track.Album.Images[0].URL,
            })
        }
        independentPlaylists[i] = models.Playlist{
            ID:       uuid.New(),
            Title:    playlist.Name,
            MediaURL: playlist.Images[0].URL,
            Songs:    songs,
        }
    }
    models.Result(w, independentPlaylists)
}

func SpotifyClientFromRequest(r *http.Request, userId *uuid.UUID) (*spotify.Client, error) {
    params := mux.Vars(r)
    param := params["service"]
    var token oauth2.Token
    err := repositories.Pool.QueryRow(r.Context(), "SELECT access_token, refresh_token, expiry FROM connections JOIN libraries ON connections.id = libraries.connection_id WHERE user_id = $1 AND platform_id = $2", userId, param).Scan(&token.AccessToken, &token.RefreshToken, &token.Expiry)
    if err != nil {
        return nil, err
    }
    auth := spotifyauth.New(
        spotifyauth.WithClientID("8c3d77ea95764c898afa8ed598c1db01"),
        spotifyauth.WithClientSecret("3773a05af958442cb77482f1e4601299"),
    )
    newToken, err := auth.RefreshToken(context.Background(), &token)
    if err != nil {
        return nil, err
    }
    newAuth := spotifyauth.New()
    httpClient := newAuth.Client(context.Background(), newToken)
    client := spotify.New(httpClient,
        // spotify.WithRetry(true),
    )
    return client, nil
}
