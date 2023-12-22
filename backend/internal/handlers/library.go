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
        return
    }
    tracks, err := client.CurrentUsersTracks(context.Background(), spotify.Limit(limit), spotify.Offset(offset))

    independentTracks := make([]models.PlatformSong, len(tracks.Tracks))
    for i, track := range tracks.Tracks {
        artists := make([]models.PlatformArtist, len(track.Artists))
        for j, artist := range track.Artists {
            artists[j] = models.PlatformArtist{
                ID:   artist.ID.String(),
                Name: artist.Name,
            }
        }
        independentTracks[i] = models.PlatformSong{
            ID:       track.ID.String(),
            Title:    track.Name,
            Artists:  artists,
            MediaURL: track.Album.Images[0].URL,
            PreviewURL: track.PreviewURL,
        }
    }

    models.Result(w, independentTracks)
}

func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
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
        return
    }
    artists, err := client.CurrentUsersTopArtists(context.Background(), spotify.Limit(limit), spotify.Offset(offset))
    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Try logging into service again")
        return
    }

    independentArtists := make([]models.PlatformArtist, len(artists.Artists))
    for i, artist := range artists.Artists {
        independentArtists[i] = models.PlatformArtist{
            ID:       artist.ID.String(),
            Name:     artist.Name,
            MediaURL: artist.Images[0].URL,
        }
    }

    models.Result(w, independentArtists)
}

func AlbumHandler(w http.ResponseWriter, r *http.Request) {
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
        return
    }

    rl := ratelimit.New(2)
    rl.Take()

    albums, err := client.CurrentUsersAlbums(context.Background(), spotify.Limit(limit), spotify.Offset(offset))
    fmt.Println(len(albums.Albums))
    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Try logging into service again")
        return
    }

    var albumIds []spotify.ID
    for _, album := range albums.Albums {
        albumIds = append(albumIds, album.ID)
    }

    albumTracks, err := client.GetAlbums(context.Background(), albumIds, spotify.Limit(limit))

    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Try logging into service again")
        return
    }

    independentAlbums := make([]models.PlatformAlbum, len(albums.Albums))
    for i, album := range albumTracks {
        songs := make([]models.PlatformSong, 0)
        for _, track := range album.Tracks.Tracks {
            artists := make([]models.PlatformArtist, len(track.Artists))
            for j, artist := range track.Artists {
                artists[j] = models.PlatformArtist{
                    ID:   artist.ID.String(),
                    Name: artist.Name,
                }
            }
            songs = append(songs, models.PlatformSong{
                ID:         track.ID.String(),
                Title:      track.Name,
                Artists:    artists,
                MediaURL:   track.Album.Images[0].URL,
                PreviewURL: track.PreviewURL,
            })
        }
        independentAlbums[i] = models.PlatformAlbum{
            ID:       album.ID.String(),
            Title:    album.Name,
            MediaURL: album.Images[0].URL,
            Songs:    songs,
        }
    }
    models.Result(w, independentAlbums)
}

func PlaylistHandler(w http.ResponseWriter, r *http.Request) {
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
        return
    }

    rl := ratelimit.New(2)
    rl.Take()
    playlists, err := client.CurrentUsersPlaylists(context.Background(), spotify.Limit(limit), spotify.Offset(offset))
    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Try logging into service again")
        return
    }

    independentPlaylists := make([]models.PlatformPlaylist, len(playlists.Playlists))
    for i, playlist := range playlists.Playlists {
        rl.Take()
        playlistTracks, err := client.GetPlaylistItems(context.Background(), playlist.ID, spotify.Limit(limit))
        if err != nil {
            models.Error(w, http.StatusInternalServerError, "Try logging into service again")
            return
        }
        songs := make([]models.PlatformSong, 0)
        for _, track := range playlistTracks.Items {
            if track.Track.Track == nil {
                continue
            }
            artists := make([]models.PlatformArtist, len(track.Track.Track.Artists))
            for j, artist := range track.Track.Track.Artists {
                artists[j] = models.PlatformArtist{
                    ID:   artist.ID.String(),
                    Name: artist.Name,
                }
            }
            songs = append(songs, models.PlatformSong {
                ID:       track.Track.Track.ID.String(),
                Title:    track.Track.Track.Name,
                Artists:  artists,
                MediaURL: track.Track.Track.Album.Images[0].URL,
                PreviewURL: track.Track.Track.PreviewURL,
            })
        }
        independentPlaylists[i] = models.PlatformPlaylist{
            ID:       playlist.ID.String(),
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
    client := spotify.New(httpClient) // spotify.WithRetry(true),

    return client, nil
}
