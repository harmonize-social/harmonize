package handlers

import (
    "backend/internal/auth"
    "backend/internal/models"
    "backend/internal/platforms"
    "backend/internal/repositories"
    "fmt"
    "net/http"
    "strconv"

    "github.com/google/uuid"
    "github.com/gorilla/mux"
)

func GetLimitOffsetSession(r *http.Request) (int, int, models.User, error) {
    limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
    if err != nil {
        limit = 10
    }
    offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
    if err != nil {
        offset = 0
    }
    id := uuid.MustParse(r.Header.Get("id"))
    _, err = auth.GetUserFromSession(id)
    var user models.User
    if err != nil {
        return 0, 0, user, err
    }
    user, err = auth.GetUserFromSession(id)
    if err != nil {
        return 0, 0, user, err
    }
    return limit, offset, user, nil
}

func LibraryHandler(w http.ResponseWriter, r *http.Request) {
    limit, offset, user, err := GetLimitOffsetSession(r)
    if err != nil {
        models.Error(w, http.StatusUnauthorized, "Malformed session")
        return
    }
    params := mux.Vars(r)
    _ = params["platform"] // TODO: use this
    itemType := params["type"]
    switch itemType {
    case "songs":
        platformSongs, err := platforms.GetSpotifySongs(&user.ID, limit, offset)
        if err != nil {
            models.Error(w, http.StatusInternalServerError, "Internal server error")
            return
        }
        err = repositories.SaveFullSongs(platformSongs)
        if err != nil {
            models.Error(w, http.StatusInternalServerError, "Internal server error")
            return
        }
        models.Result(w, platformSongs)
    case "artists":
        platformArtists, err := platforms.GetSpotifyArtists(&user.ID, limit, offset)
        if err != nil {
            models.Error(w, http.StatusInternalServerError, "Internal server error")
            return
        }
        _, err = repositories.SaveArtists(platformArtists)
        if err != nil {
            models.Error(w, http.StatusInternalServerError, "Internal server error")
            return
        }
        models.Result(w, platformArtists)
    case "albums":
        platformAlbums, err := platforms.GetSpotifyAlbums(&user.ID, limit, offset)
        if err != nil {
            models.Error(w, http.StatusInternalServerError, "Internal server error")
            return
        }
        _, err = repositories.SaveFullAlbums(platformAlbums)
        if err != nil {
            models.Error(w, http.StatusInternalServerError, "Internal server error")
            return
        }
        models.Result(w, platformAlbums)
    case "playlists":
        platformPlaylists, err := platforms.GetSpotifyPlaylists(&user.ID, limit, offset)
        if err != nil {
            models.Error(w, http.StatusInternalServerError, "Internal server error")
            return
        }
        _, err = repositories.SaveFullPlaylists(platformPlaylists)
        if err != nil {
            models.Error(w, http.StatusInternalServerError, "Internal server error")
            return
        }
        models.Result(w, platformPlaylists)
    default:
        models.Error(w, http.StatusBadRequest, "Invalid type")
    }

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
    rows, err := repositories.Pool.Query(r.Context(), "SELECT id FROM platforms WHERE id NOT IN (SELECT platform_id FROM libraries JOIN connections ON libraries.connection_id = connections.id WHERE user_id = $1)", user.ID)
    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Internal server error")
        fmt.Println(err)
        return
    }
    defer rows.Close()
    for rows.Next() {
        var platform string
        err := rows.Scan(&platform)
        if err != nil {
            fmt.Println(err)
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
            fmt.Println(err)
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
