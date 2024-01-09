package repositories

import (
	"backend/internal/models"
	"context"
	"fmt"

	"github.com/google/uuid"
)

const (
	GetSongByPlatformSpecificId = "SELECT ps.id, ps.platform_specific_id, ps.platform_id, ps.song_id FROM platform_songs ps WHERE ps.platform_id = $1 AND ps.platform_specific_id = $2;"
)

func GetSong(platform string, id string) (models.Song, error) {
	var song models.Song
	// Check if there is a song with the give platoform and platform_specific_id
	var platformSongId uuid.UUID
	err := Pool.QueryRow(context.Background(), GetSongByPlatformSpecificId, platform, id).Scan(&platformSongId, &id, &platform, &song.ID)

	if err != nil {
		fmt.Println("1:", err)
		return song, err
	}

	song, err = GetFullPostSong(song.ID)

	if err != nil {
		fmt.Println("2:", err)
		return song, err
	}

	return song, nil
}

func GetAlbum(platform string, id string) (models.Album, error) {
	var album models.Album
	// Check if there is a song with the give platoform and platform_specific_id
	sqlStatement := "SELECT pa.id AS platform_album_id, pa.platform_specific_id, pa.platform_id, pa.album_id FROM platform_albums pa WHERE pa.platform_id = $1 AND pa.platform_specific_id = $2;"
	var platformAlbumId uuid.UUID
	err := Pool.QueryRow(context.Background(), sqlStatement, platform, id).Scan(&platformAlbumId, &id, &platform, &album.ID)

	if err != nil {
		return album, err
	}

	album, err = GetFullPostAlbum(album.ID)

	if err != nil {
		return album, err
	}

	return album, nil
}

func GetPlaylist(platform string, id string) (models.Playlist, error) {
	var playlist models.Playlist
	// Check if there is a song with the give platoform and platform_specific_id
	sqlStatement := "SELECT pp.id AS platform_playlist_id, pp.platform_specific_id, pp.platform_id, pp.playlist_id FROM platform_playlists pp WHERE pp.platform_id = $1 AND pp.platform_specific_id = $2;"
	var platformPlaylistId uuid.UUID
	err := Pool.QueryRow(context.Background(), sqlStatement, platform, id).Scan(&platformPlaylistId, &id, &platform, &playlist.ID)

	if err != nil {
		return playlist, err
	}

	playlist, err = GetFullPostPlaylist(playlist.ID)

	if err != nil {
		return playlist, err
	}

	return playlist, nil
}

func GetArtist(platform string, id string) (models.Artist, error) {
	var artist models.Artist
	// Check if there is a song with the give platoform and platform_specific_id
	sqlStatement := "SELECT pa.id AS platform_artist_id, pa.platform_specific_id, pa.platform_id, pa.artist_id FROM platform_artists pa WHERE pa.platform_id = $1 AND pa.platform_specific_id = $2;"
	var platformArtistId uuid.UUID
	err := Pool.QueryRow(context.Background(), sqlStatement, platform, id).Scan(&platformArtistId, &id, &platform, &artist.ID)

	if err != nil {
		return artist, err
	}

	artist, err = GetFullPostArtist(artist.ID)

	if err != nil {
		return artist, err
	}

	return artist, nil
}
