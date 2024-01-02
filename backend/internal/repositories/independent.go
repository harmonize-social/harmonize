package repositories

import (
    "backend/internal/models"
    "context"
    "fmt"

    "github.com/google/uuid"
)

func GetSong(platform string, id string) (models.Song, error) {
    var song models.Song
    // Check if there is a song with the give platoform and platform_specific_id
    sqlStatement := "SELECT ps.id AS platform_song_id, ps.platform_specific_id, ps.platform_id, ps.song_id FROM platform_songs ps WHERE ps.platform_id = $1 AND ps.platform_specific_id = $2;"
    var platformSongId uuid.UUID
    err := Pool.QueryRow(context.Background(), sqlStatement, platform, id).Scan(&platformSongId, &id, &platform, &song.ID)

    if err != nil {
        fmt.Println("1:", err)
        return song, err
    }

    sqlStatement = "SELECT name, media_url, preview_url FROM songs WHERE id = $1;"
    err = Pool.QueryRow(context.Background(), sqlStatement, song.ID).Scan(&song.Title, &song.MediaURL, &song.PreviewURL)

    if err != nil {
        fmt.Println("2:", err)
        return song, err
    }

    // Get artists by song id
    sqlStatement = "SELECT a.id, a.name FROM artists a JOIN artists_album aa ON a.id = aa.artist_id JOIN albums al ON aa.album_id = al.id JOIN songs s ON al.id = s.album_id WHERE s.id = $1 GROUP BY a.id;"
    rows, err := Pool.Query(context.Background(), sqlStatement, song.ID)

    if err != nil {
        fmt.Println("3:", err)
        return song, err
    }

    artists := make([]models.Artist, 0)
    for rows.Next() {
        var artist models.Artist
        err = rows.Scan(&artist.ID, &artist.Name)
        if err != nil {
            fmt.Println("4:", err)
            return song, err
        }
        artists = append(artists, artist)
    }

    song.Artists = artists
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

    sqlStatement = "SELECT name, media_url FROM albums WHERE id = $1;"
    err = Pool.QueryRow(context.Background(), sqlStatement, album.ID).Scan(&album.Title, &album.MediaURL)

    if err != nil {
        return album, err
    }

    // Get artists by album id
    sqlStatement = "SELECT a.id, a.name FROM artists a JOIN artists_album aa ON a.id = aa.artist_id JOIN albums al ON aa.album_id = al.id WHERE al.id = $1 GROUP BY a.id;"
    rows, err := Pool.Query(context.Background(), sqlStatement, album.ID)

    if err != nil {
        return album, err
    }

    artists := make([]models.Artist, 0)
    for rows.Next() {
        var artist models.Artist
        err = rows.Scan(&artist.ID, &artist.Name)
        if err != nil {
            return album, err
        }
        artists = append(artists, artist)
    }

    album.Artists = artists

    // Get songs by album id

    //sqlStatement = "SELECT s.id, s.name FROM songs s JOIN albums al ON s.album_id = al.id WHERE al.id = $1;"
    sqlStatement = "SELECT s.id, s.name, s.media_url, s.preview_url FROM songs s JOIN albums al ON s.album_id = al.id WHERE al.id = $1;"

    rows, err = Pool.Query(context.Background(), sqlStatement, album.ID)

    if err != nil {
        return album, err
    }

    songs := make([]models.Song, 0)
    for rows.Next() {
        var song models.Song
        err = rows.Scan(&song.ID, &song.Title, &song.MediaURL, &song.PreviewURL)

        if err != nil {
            return album, err
        }
        songs = append(songs, song)
    }

    album.Songs = songs

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

    sqlStatement = "SELECT name, media_url FROM playlists WHERE id = $1;"
    err = Pool.QueryRow(context.Background(), sqlStatement, playlist.ID).Scan(&playlist.Title, &playlist.MediaURL)

    if err != nil {
        return playlist, err
    }

    // Get songs by playlist id
    // sqlStatement = "SELECT s.id, s.name FROM songs s JOIN playlist_songs ps ON s.id = ps.song_id JOIN playlists p ON ps.playlist_id = p.id WHERE p.id = $1;"
    sqlStatement = "SELECT s.id, s.name, s.media_url, s.preview_url FROM songs s JOIN playlist_songs ps ON s.id = ps.song_id JOIN playlists p ON ps.playlist_id = p.id WHERE p.id = $1;"
    rows, err := Pool.Query(context.Background(), sqlStatement, playlist.ID)

    if err != nil {
        return playlist, err
    }

    songs := make([]models.Song, 0)
    for rows.Next() {
        var song models.Song
        err = rows.Scan(&song.ID, &song.Title, &song.MediaURL, &song.PreviewURL)
        if err != nil {
            return playlist, err
        }
        songs = append(songs, song)
    }

    playlist.Songs = songs
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

    sqlStatement = "SELECT name, media_url FROM artists WHERE id = $1;"
    err = Pool.QueryRow(context.Background(), sqlStatement, artist.ID).Scan(&artist.Name, &artist.MediaURL)

    if err != nil {
        return artist, err
    }

    return artist, nil
}
