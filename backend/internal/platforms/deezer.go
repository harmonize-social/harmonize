package platforms

import (
	"backend/internal/auth"
	"backend/internal/models"
	"backend/internal/repositories"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/markbates/goth/providers/deezer"
	"github.com/oklookat/deezus"
)

func DeezerAuthProvider() *deezer.Provider {
	id := repositories.DeezerClientId
	secret := repositories.DeezerSecret
	provider := deezer.New(id, secret, repositories.DeezerRedirect, "basic_access", "email", "offline_access", "manage_library", "manage_community", "delete_library", "listening_history")
	return provider
}

func DeezerURL(csrf string) (string, error) {
	provider := DeezerAuthProvider()
	session, err := provider.BeginAuth(csrf)
	if err != nil {
		return "", err
	}
	url, err := session.GetAuthURL()
	if err != nil {
		return "", err
	}
	return url, nil
}

type DeezerAccessToken struct {
	AccessToken string `json:"access_token"`
}

func GetDeezerSession(provider *deezer.Provider, code string) (*deezer.Session, error) {
	var token DeezerAccessToken
	url := "https://connect.deezer.com/oauth/access_token.php?app_id=" + provider.ClientKey + "&secret=" + provider.Secret + "&code=" + code + "&output=json"
	response, err := provider.Client().Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	json.Unmarshal(body, &token)
	expiresAt := time.Now().Add(time.Hour * 24 * 365 * 100)
	session := &deezer.Session{
		AuthURL:	 "",
		AccessToken: token.AccessToken,
		ExpiresAt:	 expiresAt,
	}
	return session, nil
}

func DeezerCallback(w http.ResponseWriter, r *http.Request) {
	state := r.URL.Query().Get("state")
	provider := DeezerAuthProvider()
	user, err := auth.GetUserFromSession(uuid.MustParse(state))
	if err != nil {
		fmt.Println("Error: ", err.Error())
		models.Error(w, http.StatusUnauthorized, "Invalid token")
		return
	}
	code := r.URL.Query().Get("code")
	session, err := GetDeezerSession(provider, code)

	repositories.CreateConnectionAndLibrary(user.ID, "deezer", session.AccessToken, "", time.Now().Add(time.Hour*24*365*100))

	models.Result(w, "Ok")
}

func DeezerClientId(userId uuid.UUID) (*deezus.Client, error) {
	var client *deezus.Client
	accessToken := ""
	err := repositories.Pool.QueryRow(context.Background(), "SELECT access_token FROM connections JOIN libraries ON connections.id = libraries.connection_id WHERE user_id = $1 AND platform_id = $2", userId, "deezer").Scan(&accessToken)
	if err != nil {
		return client, err
	}

	if accessToken == "" {
		return client, errors.New("Access token not found")
	}

	client, err = deezus.New(accessToken)

	if err != nil {
		return client, err
	}
	return client, nil
}

type DeezerProvider struct {
	UserID uuid.UUID
}

func (provider DeezerProvider) GetSongs(limit int, offset int) ([]models.PlatformSong, error) {
	client, err := DeezerClientId(provider.UserID)
	if err != nil {
		return nil, err
	}
	schema, err := client.UserMeTracks(context.Background(), limit, offset)
	if err != nil {
		return nil, err
	}

	if schema.Error != nil {
		return nil, errors.New(schema.Error.Message)
	}

	simpleSongs := schema.Data

	platformSongs := make([]models.PlatformSong, 0)

	for _, simpleSong := range simpleSongs {
		song, err := client.Track(context.Background(), simpleSong.ID)
		if err != nil {
			return nil, err
		}

		if song.Error != nil {
			return nil, errors.New(schema.Error.Message)
		}
		artists := make([]models.PlatformArtist, 1)
		artists[0] = models.PlatformArtist{
			Platform: "deezer",
			ID:		  song.Artist.ID.String(),
			Name:	  song.Artist.Name,
		}

		album := models.PlatformAlbum{
			Platform: "deezer",
			ID:		  song.Album.ID.String(),
			Title:	  song.Album.Title,
			Artists:  artists,
			MediaURL: song.Album.Cover,
		}

		platformSongs = append(platformSongs, models.PlatformSong{
			Platform:	"deezer",
			ID:			song.ID.String(),
			Title:		song.Title,
			Artists:	artists,
			Album:		album,
			MediaURL:	song.Album.Cover,
			PreviewURL: song.Preview,
		})
	}
	return platformSongs, nil
}

func (provider DeezerProvider) GetAlbums(limit int, offset int) ([]models.PlatformAlbum, error) {
	return nil, nil
}

func (provider DeezerProvider) GetPlaylists(limit int, offset int) ([]models.PlatformPlaylist, error) {
	return nil, nil
}

func (provider DeezerProvider) GetArtists(limit int, offset int) ([]models.PlatformArtist, error) {
	return nil, nil
}
