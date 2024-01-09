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
	"go.uber.org/ratelimit"
	"golang.org/x/oauth2"
)

func GetPlatform(platform string, userId uuid.UUID) models.Platform {
	var provider models.Platform

	switch platform {
	case "spotify":
		provider = SpotifyProvider {
			UserID: userId,
		}
	case "deezer":
		provider = DeezerProvider {
			UserID: userId,
		}
	}
	return provider
}

func SpotifyClientId(userId *uuid.UUID) (*spotify.Client, error) {
	tokens, err := repositories.GetTokens("spotify", *userId)
	if err != nil {
		return nil, err
	}
	token := oauth2.Token{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		Expiry:		  tokens.Expiry,
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

type SpotifyProvider struct {
	UserID uuid.UUID
}

func (provider SpotifyProvider) GetSongs(limit int, offset int) ([]models.PlatformSong, error) {
	client, err := SpotifyClientId(&provider.UserID)
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
				ID:		  artist.ID.String(),
				Name:	  artist.Name,
				MediaURL: "",
			}
		}
		albumArtists := make([]models.PlatformArtist, len(track.Album.Artists))
		for j, artist := range track.Album.Artists {
			albumArtists[j] = models.PlatformArtist{
				Platform: "spotify",
				ID:		  artist.ID.String(),
				Name:	  artist.Name,
				MediaURL: "",
			}
		}
		songs[i] = models.PlatformSong{
			Platform: "spotify",
			ID:		  track.ID.String(),
			Title:	  track.Name,
			Album: models.PlatformAlbum{
				Platform: "spotify",
				ID:		  track.Album.ID.String(),
				Title:	  track.Album.Name,
				Artists:  albumArtists,
				MediaURL: track.Album.Images[0].URL,
			},
			Artists:	artists,
			MediaURL:	track.Album.Images[0].URL,
			PreviewURL: track.PreviewURL,
		}
	}
	return songs, nil
}

func (provider SpotifyProvider) GetArtists(limit int, offset int) ([]models.PlatformArtist, error) {
	client, err := SpotifyClientId(&provider.UserID)
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
			ID:		  fullArtist.ID.String(),
			Name:	  fullArtist.Name,
			MediaURL: fullArtist.Images[0].URL,
		}

		artists = append(artists, platformArtist)
	}
	return artists, nil
}

func (provider SpotifyProvider) GetAlbums(limit int, offset int) ([]models.PlatformAlbum, error) {
	client, err := SpotifyClientId(&provider.UserID)
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
				ID:		  artist.ID.String(),
				Name:	  artist.Name,
				MediaURL: "",
			}
		}
		songs := make([]models.PlatformSong, len(fullAlbum.Tracks.Tracks))
		for j, track := range fullAlbum.Tracks.Tracks {
			songs[j] = models.PlatformSong{
				Platform:	"spotify",
				ID:			track.ID.String(),
				Title:		track.Name,
				PreviewURL: track.PreviewURL,
			}
		}

		albums[i] = models.PlatformAlbum{
			Platform: "spotify",
			ID:		  album.ID.String(),
			Title:	  album.Name,
			Artists:  artists,
			Songs:	  songs,
			MediaURL: album.Images[0].URL,
		}
	}
	return albums, nil
}

func (provider SpotifyProvider) GetPlaylists(limit int, offset int) ([]models.PlatformPlaylist, error) {
	rl := ratelimit.New(2)
	rl.Take()
	client, err := SpotifyClientId(&provider.UserID)
	if err != nil {
		return nil, err
	}
	rl.Take()
	playlistsPage, err := client.CurrentUsersPlaylists(context.Background(), spotify.Limit(limit), spotify.Offset(offset))
	if err != nil {
		return nil, err
	}

	playlists := make([]models.PlatformPlaylist, len(playlistsPage.Playlists))
	for i, playlist := range playlistsPage.Playlists {
		rl.Take()
		fullPlaylist, err := client.GetPlaylistItems(context.Background(), playlist.ID)
		if err != nil {
			return nil, err
		}
		songs := make([]models.PlatformSong, 0)
		for _, track := range fullPlaylist.Items {
			if track.Track.Track == nil {
				continue
			}
			track := track.Track.Track
			artists := make([]models.PlatformArtist, len(track.Artists))
			for k, artist := range track.Artists {
				artists[k] = models.PlatformArtist{
					Platform: "spotify",
					ID:		  artist.ID.String(),
					Name:	  artist.Name,
					MediaURL: "",
				}
			}
			song := models.PlatformSong{
				Platform: "spotify",
				ID:		  track.ID.String(),
				Album: models.PlatformAlbum{
					Platform: "spotify",
					ID:		  track.Album.ID.String(),
					Title:	  track.Album.Name,
					MediaURL: track.Album.Images[0].URL,
				},
				Artists:	artists,
				Title:		track.Name,
				PreviewURL: track.PreviewURL,
			}
			songs = append(songs, song)
		}

		playlists[i] = models.PlatformPlaylist{
			Platform: "spotify",
			ID:		  playlist.ID.String(),
			Title:	  playlist.Name,
			Songs:	  songs,
			MediaURL: playlist.Images[0].URL,
		}
	}
	return playlists, nil
}

func (provider SpotifyProvider) Save(typeId string, id string) (bool, error) {
	client, err := SpotifyClientId(&provider.UserID)
	if err != nil {
		return false, err
	}
	spotifyId := spotify.ID(id)
	switch typeId {
	case "artist":
		err = client.FollowArtist(context.Background(), spotifyId)
	case "song":
		err = client.AddTracksToLibrary(context.Background(), spotifyId)
	case "album":
		err = client.AddAlbumsToLibrary(context.Background(), spotifyId)
	case "playlist":
		err = client.FollowPlaylist(context.Background(), spotifyId, false)
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
