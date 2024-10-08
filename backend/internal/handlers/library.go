package handlers

import (
	"backend/internal/auth"
	"backend/internal/models"
	"backend/internal/platforms"
	"backend/internal/repositories"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

/*
Get all songs, artists, albums, or playlists from a platform

GET /library/{service}/{type}?limit={limit}&offset={offset}
*/
func LibraryHandler(w http.ResponseWriter, r *http.Request) {
	limit, offset, user, err := GetLimitOffsetSession(r)
	if err != nil {
		models.Error(w, http.StatusUnauthorized, "Malformed session")
		return
	}
	params := mux.Vars(r)
	platform := params["service"]

	if platform == "" {
		models.Error(w, http.StatusBadRequest, "Missing platform")
		return
	}

	var provider models.Platform

	switch platform {
	case "spotify":
		provider = platforms.SpotifyProvider{
			UserID: user.ID,
		}
	case "deezer":
		provider = platforms.DeezerProvider{
			UserID: user.ID,
		}
	default:
		models.Error(w, http.StatusBadRequest, "Invalid platform")
		return
	}

	itemType := params["type"]
	switch itemType {
	case "songs":
		platformSongs, err := provider.GetSongs(limit, offset)
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
		platformArtists, err := provider.GetArtists(limit, offset)
		if err != nil {
			fmt.Println("artists 1: ", err)
			models.Error(w, http.StatusInternalServerError, "Internal server error")
			return
		}
		_, err = repositories.SaveArtists(platformArtists)
		if err != nil {
			fmt.Println("artists 2: ", err)
			models.Error(w, http.StatusInternalServerError, "Internal server error")
			return
		}
		models.Result(w, platformArtists)
	case "albums":
		platformAlbums, err := provider.GetAlbums(limit, offset)
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
		platformPlaylists, err := provider.GetPlaylists(limit, offset)
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

/*
Delete the user's library for a platform

DELETE /library/disconnect?platform={platform}
*/
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

/*
Get the platform names that the user is connected to:

e.g. ["spotify", "deezer"]

GET /library/connected
*/
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

/*
Return the platforms that the user is not connected to:

e.g.
{
	[platform]: [oauth url],
}


GET /library/unconnected
*/
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
