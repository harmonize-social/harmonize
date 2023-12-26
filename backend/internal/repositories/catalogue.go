package repositories

import (
    "backend/internal/models"
    "context"
    "errors"
    "fmt"

    "github.com/google/uuid"
    "github.com/zmb3/spotify/v2"
)

func CreateSongPost(song models.Song, userID uuid.UUID, caption string) (models.Post, error) {
    var post models.Post

    sqlStatement := "INSERT INTO posts (id, user_id, caption, type, type_specific_id, created_at) VALUES (uuid_generate_v4() ,$1, $2, $3, $4, now()) RETURNING id, created_at, caption, type;"
    err := Pool.QueryRow(context.Background(), sqlStatement, userID, caption, "song", song.ID).Scan(&post.ID, &post.CreatedAt, &post.Caption, &post.Type)

    if err != nil {
        return post, err
    }

    post.Content = song

    return post, nil
}

func CreateAlbumPost(album models.Album, userID uuid.UUID, caption string) (models.Post, error) {
    var post models.Post

    sqlStatement := "INSERT INTO posts (id, user_id, caption, type, type_specific_id, created_at) VALUES (uuid_generate_v4() ,$1, $2, $3, $4, now()) RETURNING id, created_at, caption, type;"
    err := Pool.QueryRow(context.Background(), sqlStatement, userID, caption, "album", album.ID).Scan(&post.ID, &post.CreatedAt, &post.Caption, &post.Type)

    if err != nil {
        return post, err
    }

    post.Content = album

    return post, nil
}

func CreatePlaylistPost(playlist models.Playlist, userID uuid.UUID, caption string) (models.Post, error) {
    var post models.Post

    sqlStatement := "INSERT INTO posts (id, user_id, caption, type, type_specific_id, created_at) VALUES (uuid_generate_v4() ,$1, $2, $3, $4, now()) RETURNING id, created_at, caption, type;"
    err := Pool.QueryRow(context.Background(), sqlStatement, userID, caption, "playlist", playlist.ID).Scan(&post.ID, &post.CreatedAt, &post.Caption, &post.Type)

    if err != nil {
        return post, err
    }

    post.Content = playlist

    return post, nil
}

func CreateArtistPost(artist models.Artist, userID uuid.UUID, caption string) (models.Post, error) {
    var post models.Post

    sqlStatement := "INSERT INTO posts (id, user_id, caption, type, type_specific_id, created_at) VALUES (uuid_generate_v4() ,$1, $2, $3, $4, now()) RETURNING id, created_at, caption, type;"
    err := Pool.QueryRow(context.Background(), sqlStatement, userID, caption, "artist", artist.ID).Scan(&post.ID, &post.CreatedAt, &post.Caption, &post.Type)

    if err != nil {
        return post, err
    }

    post.Content = artist

    return post, nil
}

func SaveSpotifySongs(tracks *spotify.SavedTrackPage) ([]models.PlatformSong, error) {
    independentTracks := make([]models.PlatformSong, len(tracks.Tracks))
    insertAlbumStatment := `SELECT * FROM insert_new_album($1, $2, $3, $4);`
    insertSongStatment := `SELECT * FROM insert_new_song($1, $2, $3, $4, $5, $6);`
    insertArtistStatment := `SELECT * FROM insert_new_artist($1, $2, $3, $4);`
    for i, track := range tracks.Tracks {
        var albumId uuid.UUID
        var albumPlatformId string
        err := Pool.QueryRow(context.Background(), insertAlbumStatment, "spotify", track.Album.ID.String(), track.Album.Name, track.Album.Images[0]).Scan(&albumId, &albumPlatformId)
        if err != nil {
            return independentTracks, err
        }
        artists := make([]models.PlatformArtist, len(track.Artists))
        for j, artist := range track.Artists {
            artists[j] = models.PlatformArtist{
                ID:   artist.ID.String(),
                Name: artist.Name,
            }
            var artistID uuid.UUID
            err := Pool.QueryRow(context.Background(), insertArtistStatment, "spotify", artist.ID.String(), artist.Name, "").Scan(&artistID, &artist.ID)

            if err != nil {
                return independentTracks, err
            }

            tag, err := Pool.Exec(context.Background(), "INSERT INTO artists_album (id, artist_id, album_id) VALUES (uuid_generate_v4(), $1, $2);", artistID, albumId)

            if err != nil {
                return independentTracks, err
            }

            if tag.RowsAffected() == 0 {
                return independentTracks, errors.New("Error saving artist album")
            }
        }
        independentTracks[i] = models.PlatformSong{
            ID:         track.ID.String(),
            Title:      track.Name,
            Artists:    artists,
            MediaURL:   track.Album.Images[0].URL,
            PreviewURL: track.PreviewURL,
        }

        tag, err := Pool.Exec(context.Background(), insertSongStatment, "spotify", albumId, track.ID.String(), track.Name, track.Album.Images[0].URL, track.PreviewURL)
        if err != nil {
            return independentTracks, err
        }
        if tag.RowsAffected() == 0 {
            return independentTracks, errors.New("Error saving song")
        }
    }
    return independentTracks, nil
}

func SaveSpotifyAlbums(albums *spotify.SavedAlbumPage) ([]models.PlatformAlbum, error) {
    independentAlbums := make([]models.PlatformAlbum, len(albums.Albums))
    insertAlbumStatment := `SELECT * FROM insert_new_album($1, $2, $3, $4);`
    insertArtistStatment := `SELECT * FROM insert_new_artist($1, $2, $3, $4);`
    insertSongStatment := `SELECT * FROM insert_new_song($1, $2, $3, $4, $5, $6);`
    for i, album := range albums.Albums {
        var albumId uuid.UUID
        var albumPlatformId uuid.UUID
        err := Pool.QueryRow(context.Background(), insertAlbumStatment, "spotify", album.ID.String(), album.Name, album.Images[0].URL).Scan(&albumId, &albumPlatformId)
        fmt.Println(albumId, albumPlatformId)
        if err != nil {
            return independentAlbums, err
        }
        artists := make([]models.PlatformArtist, len(album.Artists))
        for j, artist := range album.Artists {
            artists[j] = models.PlatformArtist{
                ID:   artist.ID.String(),
                Name: artist.Name,
            }
            var artistID uuid.UUID
            var artistPlatformId uuid.UUID
            err := Pool.QueryRow(context.Background(), insertArtistStatment, "spotify", artist.ID.String(), artist.Name, "").Scan(&artistID, &artistPlatformId)

            if err != nil {
                return independentAlbums, err
            }

            tag, err := Pool.Exec(context.Background(), "INSERT INTO artists_album (id, artist_id, album_id) VALUES (uuid_generate_v4(), $1, $2);", artistID, albumId)

            if err != nil {
                return independentAlbums, err
            }

            if tag.RowsAffected() == 0 {
                return independentAlbums, errors.New("Error saving artist album")
            }
        }
        songs := make([]models.PlatformSong, 0)
        for _, track := range album.Tracks.Tracks {
            var songId uuid.UUID
            var songPlatformId uuid.UUID
            err := Pool.QueryRow(context.Background(), insertSongStatment, "spotify", albumId, track.ID.String(), track.Name, track.Album.Images[0].URL, track.PreviewURL).Scan(&songId, &songPlatformId)
            if err != nil {
                fmt.Println(err)
                return independentAlbums, err
            }
            songs = append(songs, models.PlatformSong{
                ID:         track.ID.String(),
                Title:      track.Name,
                Artists:    artists,
                MediaURL:   album.Images[0].URL,
                PreviewURL: track.PreviewURL,
            })
        }
        independentAlbums[i] = models.PlatformAlbum{
            ID:       album.ID.String(),
            Title:    album.Name,
            Artists:  artists,
            Songs:    songs,
            MediaURL: album.Images[0].URL,
        }
    }
    return independentAlbums, nil
}

func SaveSpotifyPlaylists(playlists *spotify.SimplePlaylistPage, playlistTracks map[string][]spotify.FullTrack) ([]models.PlatformPlaylist, error) {
    independentPlaylists := make([]models.PlatformPlaylist, len(playlists.Playlists))
    insertPlaylistStatment := `SELECT * FROM insert_new_playlist($1, $2, $3, $4);`
    insertSongIntoPlaylistStatment := `INSERT INTO playlist_songs (id, playlist_id, song_id) VALUES (uuid_generate_v4(), $1, $2);`
    insertAlbumStatment := `SELECT * FROM insert_new_album($1, $2, $3, $4);`
    insertSongStatment := `SELECT * FROM insert_new_song($1, $2, $3, $4, $5, $6);`
    insertArtistStatment := `SELECT * FROM insert_new_artist($1, $2, $3, $4);`
    for i, playlist := range playlists.Playlists {
        tracks := playlistTracks[playlist.ID.String()]
        var playlistId uuid.UUID
        var playlistPlatformId uuid.UUID
        err := Pool.QueryRow(context.Background(), insertPlaylistStatment, "spotify", playlist.ID.String(), playlist.Name, playlist.Images[0].URL).Scan(&playlistId, &playlistPlatformId)
        if err != nil {
            return independentPlaylists, err
        }
        fmt.Println(len(playlist.Images))
        independentPlaylists[i] = models.PlatformPlaylist{
            ID:       playlist.ID.String(),
            Title:    playlist.Name,
            MediaURL: playlist.Images[0].URL,
        }
        var songs []models.PlatformSong
        for _, track := range tracks {
            var albumId uuid.UUID
            var albumPlatformId uuid.UUID
            err := Pool.QueryRow(context.Background(), insertAlbumStatment, "spotify", track.Album.ID.String(), track.Album.Name, track.Album.Images[0].URL, track.PreviewURL).Scan(&albumId, &albumPlatformId)
            if err != nil {
                return independentPlaylists, err
            }
            artists := make([]models.PlatformArtist, len(track.Artists))
            for j, artist := range track.Artists {
                artists[j] = models.PlatformArtist{
                    ID:   artist.ID.String(),
                    Name: artist.Name,
                }
                var artistID uuid.UUID
                var artistPlatformId uuid.UUID
                err := Pool.QueryRow(context.Background(), insertArtistStatment, "spotify", artist.ID.String(), artist.Name, "").Scan(&artistID, &artistPlatformId)

                if err != nil {
                    return independentPlaylists, err
                }

                tag, err := Pool.Exec(context.Background(), "INSERT INTO artists_album (id, artist_id, album_id) VALUES (uuid_generate_v4(), $1, $2);", artistID, albumId)

                if err != nil {
                    return independentPlaylists, err
                }

                if tag.RowsAffected() == 0 {
                    return independentPlaylists, errors.New("Error saving artist album")
                }

            }
            var songId uuid.UUID
            var songPlatformId uuid.UUID
            err = Pool.QueryRow(context.Background(), insertSongStatment, "spotify", albumId, track.ID.String(), track.Name, track.Album.Images[0].URL, track.PreviewURL).Scan(&songId, &songPlatformId)
            if err != nil {
                return independentPlaylists, err
            }

            tag, err := Pool.Exec(context.Background(), insertSongIntoPlaylistStatment, playlistId, songId)

            if err != nil {
                return independentPlaylists, err
            }

            if tag.RowsAffected() == 0 {
                return independentPlaylists, errors.New("Error saving song")
            }

            songs = append(songs, models.PlatformSong{
                ID:         track.ID.String(),
                Title:      track.Name,
                Artists:    artists,
                MediaURL:   track.Album.Images[0].URL,
                PreviewURL: track.PreviewURL,
            })
        }
        independentPlaylists[i].Songs = songs
    }
    return independentPlaylists, nil
}

func SaveSpotifyArtists(artists []spotify.FullArtist) ([]models.PlatformArtist, error) {
    independentArtists := make([]models.PlatformArtist, len(artists))
    insertArtistStatment := `SELECT * FROM insert_new_artist($1, $2, $3, $4);`
    for i, artist := range artists {
        var artistID uuid.UUID
        var artistPlatformId uuid.UUID
        err := Pool.QueryRow(context.Background(), insertArtistStatment, "spotify", artist.ID.String(), artist.Name, artist.Images[0].URL).Scan(&artistID, &artistPlatformId)
        if err != nil {
            return independentArtists, err
        }
        independentArtists[i] = models.PlatformArtist{
            ID:       artist.ID.String(),
            Name:     artist.Name,
            MediaURL: artist.Images[0].URL,
        }
    }
    return independentArtists, nil
}

func GetSong(platform string, id string) (models.Song, error) {
    var song models.Song
    // Check if there is a song with the give platoform and platform_specific_id
    sqlStatement := "SELECT ps.id AS platform_song_id, ps.platform_specific_id, ps.platform_id, ps.song_id FROM platform_songs ps WHERE ps.platform_id = $1 AND ps.platform_specific_id = $2;"
    var platformSongId uuid.UUID
    err := Pool.QueryRow(context.Background(), sqlStatement, platform, id).Scan(&platformSongId, &id, &platform, &song.ID)

    if err != nil {
        return song, err
    }

    sqlStatement = "SELECT name, media_url, preview_url FROM songs WHERE id = $1;"
    err = Pool.QueryRow(context.Background(), sqlStatement, song.ID).Scan(&song.Title, &song.MediaURL, &song.PreviewURL)

    if err != nil {
        return song, err
    }

    // Get artists by song id
    sqlStatement = "SELECT a.id, a.name FROM artists a JOIN artists_album aa ON a.id = aa.artist_id JOIN albums al ON aa.album_id = al.id JOIN songs s ON al.id = s.album_id WHERE s.id = $1;"
    rows, err := Pool.Query(context.Background(), sqlStatement, song.ID)

    if err != nil {
        return song, err
    }

    artists := make([]models.Artist, 0)
    for rows.Next() {
        var artist models.Artist
        err = rows.Scan(&artist.ID, &artist.Name)
        if err != nil {
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

/*
DROP FUNCTION IF EXISTS insert_new_artist;
CREATE OR REPLACE FUNCTION insert_new_artist(new_platform_id VARCHAR(1024), platform_specific_id_input VARCHAR(1024), new_name VARCHAR(1024), new_media_url VARCHAR(1024))
RETURNS TABLE (value1 UUID, value2 UUID) AS $$
DECLARE
    artists_artist_id UUID;
    platform_artists_artist_id UUID;
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM platform_artists WHERE platform_specific_id = platform_specific_id_input
    ) THEN
        INSERT INTO artists (id, name, media_url)
        VALUES (uuid_generate_v4(), new_name, new_media_url)
        RETURNING id INTO artists_artist_id;

        INSERT INTO platform_artists (id, platform_id, platform_specific_id, artist_id)
        VALUES (uuid_generate_v4(), new_platform_id, platform_specific_id_input, artists_artist_id)
        RETURNING id INTO platform_artists_artist_id;

        RETURN QUERY SELECT artists_artist_id, platform_artists_artist_id;
    ELSE
        SELECT artists.id FROM platform_artists
        JOIN artists ON platform_artists.artist_id = artists.id
        WHERE platform_specific_id = platform_specific_id_input INTO artists_artist_id;

        SELECT platform_artists.id FROM platform_artists
        WHERE platform_specific_id = platform_specific_id_input INTO platform_artists_artist_id;

        RETURN QUERY SELECT artists_artist_id, platform_artists_artist_id;
    END IF;
END;
$$ LANGUAGE plpgsql;


DROP FUNCTION IF EXISTS insert_new_album;
CREATE OR REPLACE FUNCTION insert_new_album(new_platform_id VARCHAR(1024), platform_specific_album_id VARCHAR(1024), new_name VARCHAR(1024), new_media_url VARCHAR(1024))
RETURNS TABLE (value1 UUID, value2 UUID) AS $$
DECLARE
    albums_album_id UUID;
    platform_albums_album_id UUID;
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM platform_albums WHERE platform_specific_id = platform_specific_album_id
    ) THEN
        INSERT INTO albums (id, name, media_url)
        VALUES (uuid_generate_v4(), new_name, new_media_url)
        RETURNING id INTO albums_album_id;

        INSERT INTO platform_albums (id, platform_id, platform_specific_id, album_id)
        VALUES (uuid_generate_v4(), new_platform_id, platform_specific_album_id, albums_album_id)
        RETURNING id INTO platform_albums_album_id;

        RETURN QUERY SELECT albums_album_id, platform_albums_album_id;
    ELSE
        SELECT albums.id FROM platform_albums
        JOIN albums ON platform_albums.album_id = albums.id
        WHERE platform_specific_id = platform_specific_album_id INTO albums_album_id;

        SELECT platform_albums.id FROM platform_albums
        WHERE platform_specific_id = platform_specific_album_id INTO platform_albums_album_id;

        RETURN QUERY SELECT albums_album_id, platform_albums_album_id;
    END IF;
END;
$$ LANGUAGE plpgsql;

DROP FUNCTION IF EXISTS insert_new_playlist;
CREATE OR REPLACE FUNCTION insert_new_playlist(new_platform_id VARCHAR(1024), platform_specific_id_input VARCHAR(1024), new_name VARCHAR(1024), new_media_url VARCHAR(1024))
RETURNS TABLE (value1 UUID, value2 UUID) AS $$
DECLARE
    playlists_playlist_id UUID;
    platform_playlists_playlist_id UUID;
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM platform_playlists WHERE platform_specific_id = platform_specific_id_input
    ) THEN
        INSERT INTO playlists (id, name, media_url)
        VALUES (uuid_generate_v4(), new_name, new_media_url)
        RETURNING id INTO playlists_playlist_id;

        INSERT INTO platform_playlists (id, platform_id, platform_specific_id, playlist_id)
        VALUES (uuid_generate_v4(), new_platform_id, platform_specific_id_input, playlists_playlist_id)
        RETURNING id INTO platform_playlists_playlist_id;

        RETURN QUERY SELECT playlists_playlist_id, platform_playlists_playlist_id;
    ELSE
        SELECT playlists.id FROM platform_playlists
        JOIN playlists ON platform_playlists.playlist_id = playlists.id
        WHERE platform_specific_id = platform_specific_id_input INTO playlists_playlist_id;

        SELECT platform_playlists.id FROM platform_playlists
        WHERE platform_specific_id = platform_specific_id_input INTO platform_playlists_playlist_id;

        RETURN QUERY SELECT playlists_playlist_id, platform_playlists_playlist_id;
    END IF;
END;
$$ LANGUAGE plpgsql;

DROP FUNCTION IF EXISTS insert_new_song;
CREATE OR REPLACE FUNCTION insert_new_song(new_platform_id VARCHAR(1024), album_id UUID, platform_specific_song_id VARCHAR(1024), new_name VARCHAR(1024), new_media_url VARCHAR(1024), new_preview_url VARCHAR(1024))
RETURNS TABLE (value1 UUID, value2 UUID) AS $$
DECLARE
    songs_song_id UUID;
    platform_songs_song_id UUID;
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM platform_songs WHERE platform_specific_id = platform_specific_song_id
    ) THEN
        INSERT INTO songs (id, name, album_id, media_url, preview_url)
        VALUES (uuid_generate_v4(), new_name, album_id, new_media_url, new_preview_url)
        RETURNING id INTO songs_song_id;

        INSERT INTO platform_songs (id, platform_id, platform_specific_id, song_id)
        VALUES (uuid_generate_v4(), new_platform_id, platform_specific_song_id, songs_song_id)
        RETURNING id INTO platform_songs_song_id;

        RETURN QUERY SELECT songs_song_id, platform_songs_song_id;
    ELSE
        SELECT songs.id FROM platform_songs
        JOIN songs ON platform_songs.song_id = songs.id
        WHERE platform_specific_id = platform_specific_song_id INTO songs_song_id;

        SELECT platform_songs.id FROM platform_songs
        WHERE platform_specific_id = platform_specific_song_id INTO platform_songs_song_id;

        RETURN QUERY SELECT songs_song_id, platform_songs_song_id;
    END IF;
END;
$$ LANGUAGE plpgsql;
*/
