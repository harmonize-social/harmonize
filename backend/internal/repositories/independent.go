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

/*

CREATE TABLE IF NOT EXISTS platforms(
    id VARCHAR(1024) PRIMARY KEY,
    name VARCHAR(1024) NOT NULL,
    icon_id UUID REFERENCES images (id) NOT NULL
);


CREATE TABLE IF NOT EXISTS artists(
    id UUID PRIMARY KEY,
    name VARCHAR(1024) NOT NULL,
    media_url VARCHAR(1024) NOT NULL
);

CREATE TABLE IF NOT EXISTS albums(
    id UUID PRIMARY KEY,
    name VARCHAR(1024) NOT NULL,
    media_url VARCHAR(1024) NOT NULL
);

CREATE TABLE IF NOT EXISTS artists_album(
    id UUID PRIMARY KEY,
    artist_id UUID REFERENCES artists (id) NOT NULL,
    album_id UUID REFERENCES albums (id) NOT NULL
);


CREATE TABLE IF NOT EXISTS songs(
    id UUID PRIMARY KEY,
    name VARCHAR(1024) NOT NULL,
    album_id UUID REFERENCES albums (id) NOT NULL,
    media_url VARCHAR(1024) NOT NULL,
    preview_url VARCHAR(1024) NOT NULL
);

CREATE TABLE IF NOT EXISTS playlists(
    id UUID PRIMARY KEY,
    name VARCHAR(1024) NOT NULL,
    media_url VARCHAR(1024) NOT NULL
);


CREATE TABLE IF NOT EXISTS playlist_songs(
    id UUID PRIMARY KEY,
    playlist_id UUID REFERENCES playlists (id) NOT NULL,
    song_id UUID REFERENCES songs (id) NOT NULL
);


CREATE TABLE IF NOT EXISTS platform_artists(
    id UUID PRIMARY KEY,
    platform_specific_id VARCHAR(1024) NOT NULL,
    platform_id VARCHAR(1024) REFERENCES platforms(id) NOT NULL,
    artist_id UUID REFERENCES artists (id) NOT NULL
);


CREATE TABLE IF NOT EXISTS platform_albums(
    id UUID PRIMARY KEY,
    platform_specific_id VARCHAR(1024) NOT NULL,
    platform_id VARCHAR(1024) REFERENCES platforms(id) NOT NULL,
    album_id UUID REFERENCES albums (id) NOT NULL
);


CREATE TABLE IF NOT EXISTS platform_songs(
    id UUID PRIMARY KEY,
    platform_specific_id VARCHAR(1024) NOT NULL,
    platform_id VARCHAR(1024) REFERENCES platforms(id) NOT NULL,
    song_id UUID REFERENCES songs (id) NOT NULL
);


CREATE TABLE IF NOT EXISTS platform_playlists(
    id UUID PRIMARY KEY,
    platform_specific_id VARCHAR(1024) NOT NULL,
    platform_id VARCHAR(1024) REFERENCES platforms(id) NOT NULL,
    playlist_id UUID REFERENCES playlists (id) NOT NULL
);


CREATE TABLE IF NOT EXISTS posts(
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users (id) NOT NULL,
    created_at timestamptz NOT NULL,
    caption VARCHAR(1024) NOT NULL,
    type VARCHAR(1024) NOT NULL,
    type_specific_id UUID
);
*/
