package handlers

import (
    "backend/internal/auth"
    "backend/internal/models"
    "backend/internal/platforms"
    "backend/internal/repositories"
    "context"
    "fmt"
    "net/http"
    "strconv"

    "github.com/google/uuid"
    "github.com/zmb3/spotify/v2"
    "go.uber.org/ratelimit"
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
    user, err := auth.GetUserFromSession(id)
    if err != nil {
        models.Error(w, http.StatusUnauthorized, "Malformed session")
        return
    }
    client, err := platforms.SpotifyClientId(&user.ID)
    if err != nil {
        fmt.Println(err)
        models.Error(w, http.StatusInternalServerError, "Try logging into service again")
        return
    }
    tracks, err := client.CurrentUsersTracks(context.Background(), spotify.Limit(limit), spotify.Offset(offset))

    independentTracks, err := repositories.SaveSpotifySongs(tracks)

    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Try logging into service again")
        fmt.Println(err)
        return
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
    user, err := auth.GetUserFromSession(id)
    if err != nil {
        models.Error(w, http.StatusUnauthorized, "Malformed session")
        return
    }
    client, err := platforms.SpotifyClientId(&user.ID)
    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Try logging into service again")
        return
    }
    artistsPage, err := client.CurrentUsersFollowedArtists(context.Background(), spotify.Limit(limit), spotify.Offset(offset))
    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Try logging into service again")
        return
    }

    artists := make([]spotify.FullArtist, 0)
    for _, artist := range artistsPage.Artists {
        fullArtist, err := client.GetArtist(context.Background(), artist.ID)

        if err != nil {
            models.Error(w, http.StatusInternalServerError, "Try logging into service again")
            return
        }

        artists = append(artists, *fullArtist)
    }

    independentArtists, err := repositories.SaveSpotifyArtists(artists)
    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Try logging into service again")
        return
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
    user, err := auth.GetUserFromSession(id)
    if err != nil {
        models.Error(w, http.StatusUnauthorized, "Malformed session")
        return
    }
    client, err := platforms.SpotifyClientId(&user.ID)
    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Try logging into service again")
        fmt.Println(err)
        return
    }

    rl := ratelimit.New(2)
    rl.Take()

    albums, err := client.CurrentUsersAlbums(context.Background(), spotify.Limit(limit), spotify.Offset(offset))
    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Try logging into service again")
        fmt.Println(err)
        return
    }

    independentAlbums, err := repositories.SaveSpotifyAlbums(albums)

    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Try logging into service again")
        fmt.Println(err)
        return
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
    user, err := auth.GetUserFromSession(id)
    if err != nil {
        models.Error(w, http.StatusUnauthorized, "Malformed session")
        return
    }
    client, err := platforms.SpotifyClientId(&user.ID)
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

    playlistTracks := make(map[string][]spotify.FullTrack, 0)
    for _, playlist := range playlists.Playlists {
        rl.Take()
        tracks, err := client.GetPlaylistItems(context.Background(), spotify.ID(playlist.ID), spotify.Limit(100))
        if err != nil {
            models.Error(w, http.StatusInternalServerError, "Try logging into service again")
            fmt.Println(err)
            return
        }
        playlistTracks[playlist.ID.String()] = make([]spotify.FullTrack, 0)
        for _, track := range tracks.Items {
            if track.Track.Track == nil {
                continue
            }
            playlistTracks[playlist.ID.String()] = append(playlistTracks[playlist.ID.String()], *track.Track.Track)
        }
    }

    independentPlaylists, err := repositories.SaveSpotifyPlaylists(playlists, playlistTracks)

    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Try logging into service again")
        fmt.Println(err)
        return
    }

    models.Result(w, independentPlaylists)
}

func DeleteConnectedPlatformHandler(w http.ResponseWriter, r *http.Request) {
    platform := r.URL.Query().Get("platform")
    id := uuid.MustParse(r.Header.Get("id"))
    user, err := auth.GetUserFromSession(id)
    if err != nil {
        models.Error(w, http.StatusUnauthorized, "Malformed session")
        return
    }

    if platform == "" {
        models.Error(w, http.StatusBadRequest, "Missing platform")
        return
    }

    // Library has platform_id and connection_id, connection has user_id, so you have to join to get connection_id
    var connectionId uuid.UUID
    err = repositories.Pool.QueryRow(r.Context(), "SELECT connection_id FROM libraries JOIN connections ON libraries.connection_id = connections.id WHERE user_id = $1 AND platform_id = $2", user.ID, platform).Scan(&connectionId)
    if err != nil {
        fmt.Println("1:", err)
        models.Error(w, http.StatusInternalServerError, "Internal server error")
        return
    }

    _, err = repositories.Pool.Exec(r.Context(), "DELETE FROM libraries WHERE connection_id = $1 AND platform_id = $2", connectionId, platform)
    if err != nil {
        fmt.Println("2:", err)
        models.Error(w, http.StatusInternalServerError, "Internal server error")
        return
    }

    _, err = repositories.Pool.Exec(r.Context(), "DELETE FROM connections WHERE id = $1", connectionId)
    if err != nil {
        fmt.Println("2:", err)
        models.Error(w, http.StatusInternalServerError, "Internal server error")
        return
    }

    models.Result(w, "Success")
}

func ConnectedPlatforumsHandler(w http.ResponseWriter, r *http.Request) {
    id := uuid.MustParse(r.Header.Get("id"))
    user, err := auth.GetUserFromSession(id)
    if err != nil {
        models.Error(w, http.StatusUnauthorized, "Malformed session")
        return
    }
    platforms := make([]string, 0)
    rows, err := repositories.Pool.Query(r.Context(), "SELECT platform_id FROM libraries JOIN connections ON libraries.connection_id = connections.id WHERE user_id = $1", user.ID)
    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Internal server error")
        return
    }
    defer rows.Close()
    for rows.Next() {
        var platform string
        err := rows.Scan(&platform)
        if err != nil {
            models.Error(w, http.StatusInternalServerError, "Internal server error")
            return
        }
        platforms = append(platforms, platform)
    }
    models.Result(w, platforms)
}

func UnconnectedPlatformsHandler(w http.ResponseWriter, r *http.Request) {
    id := uuid.MustParse(r.Header.Get("id"))
    user, err := auth.GetUserFromSession(id)
    if err != nil {
        models.Error(w, http.StatusUnauthorized, "Malformed session")
        return
    }
    platformNames := make([]string, 0)
    rows, err := repositories.Pool.Query(r.Context(), "SELECT id FROM platformNames WHERE id NOT IN (SELECT platform_id FROM libraries JOIN connections ON libraries.connection_id = connections.id WHERE user_id = $1)", user.ID)
    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Internal server error")
        return
    }
    defer rows.Close()
    for rows.Next() {
        var platform string
        err := rows.Scan(&platform)
        if err != nil {
            models.Error(w, http.StatusInternalServerError, "Internal server error")
            return
        }
        platformNames = append(platformNames, platform)
    }
    platformOauths := make(map[string]string, 0)
    for _, platform := range platformNames {
        url := ""
        if platform == "spotify" {
            url, err = platforms.SpotifyURL(id.String())
        } else if platform == "deezer" {
            url, err = platforms.DeezerURL(id.String())
        }
        if err != nil {
            models.Error(w, http.StatusInternalServerError, "Internal server error")
            return
        }
        if url == "" {
            continue
        }
        platformOauths[platform] = url
    }
    models.Result(w, platformOauths)
}
