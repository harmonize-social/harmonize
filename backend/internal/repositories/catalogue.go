package repositories

import (
    "backend/internal/models"
    "context"
    "errors"
    "fmt"

    "github.com/google/uuid"
    "github.com/zmb3/spotify/v2"
)

const (
    insertNewAlbumStatment         = `SELECT * FROM insert_new_album($1, $2, $3, $4);`
    insertNewSongStatment          = `SELECT * FROM insert_new_song($1, $2, $3, $4, $5, $6);`
    insertNewArtistStatment        = `SELECT * FROM insert_new_artist($1, $2, $3, $4);`
    insertNewPlaylistStatment      = `SELECT * FROM insert_new_playlist($1, $2, $3, $4);`
    insertSongIntoPlaylistStatment = `INSERT INTO playlist_songs (id, playlist_id, song_id) VALUES (uuid_generate_v4(), $1, $2);`
    insertArtistAlbumStatment      = `INSERT INTO artists_album (id, artist_id, album_id) VALUES (uuid_generate_v4(), $1, $2);`
)

/*
Saves artist without albums or songs
*/
func SaveArtist(platformArtist models.PlatformArtist) (uuid.UUID, models.Artist, error) {
    var artist models.Artist
    var artistPlatformId uuid.UUID
    err := Pool.QueryRow(context.Background(), insertNewArtistStatment, platformArtist.Platform, platformArtist.ID, platformArtist.Name, platformArtist.MediaURL).Scan(&artist.ID, &artistPlatformId)
    if err != nil {
        return uuid.Nil, artist, err
    }
    return artistPlatformId, artist, nil
}

/*
Saves multiple artists
*/
func SaveArtists(artists []models.PlatformArtist) ([]models.Artist, error) {
    independentArtists := make([]models.Artist, len(artists))
    for i, artist := range artists {
        _, independentArtist, err := SaveArtist(artist)
        if err != nil {
            return independentArtists, err
        }
        independentArtists[i] = independentArtist
    }
    return independentArtists, nil
}

/*
Saves album without songs
*/
func SaveAlbum(platformAlbum models.PlatformAlbum) (uuid.UUID, models.Album, error) {
    var album models.Album
    var albumPlatformId uuid.UUID
    err := Pool.QueryRow(context.Background(), insertNewAlbumStatment, platformAlbum.Platform, platformAlbum.ID, platformAlbum.Title, platformAlbum.MediaURL).Scan(&album.ID, &albumPlatformId)
    if err != nil {
        return uuid.Nil, album, err
    }
    return albumPlatformId, album, nil
}

/*
Saves playlist without songs
*/
func SavePlaylist(platformPlaylist models.PlatformPlaylist) (uuid.UUID, models.Playlist, error) {
    var playlist models.Playlist
    var playlistPlatformId uuid.UUID
    err := Pool.QueryRow(context.Background(), insertNewPlaylistStatment, platformPlaylist.Platform, platformPlaylist.ID, platformPlaylist.Title, platformPlaylist.MediaURL).Scan(&playlist.ID, &playlistPlatformId)
    if err != nil {
        return uuid.Nil, playlist, err
    }
    return playlistPlatformId, playlist, nil
}

/*
Saves artist-album relationship
*/
func SaveArtistAlbum(artistId uuid.UUID, albumId uuid.UUID) error {
    tag, err := Pool.Exec(context.Background(), insertArtistAlbumStatment, artistId, albumId)
    if err != nil {
        return err
    }
    if tag.RowsAffected() == 0 {
        return errors.New("Error saving artist album")
    }
    return nil
}

/*
Saves song-playlist relationship
*/
func SavePlaylistSong(playlistId uuid.UUID, songId uuid.UUID) error {
    tag, err := Pool.Exec(context.Background(), insertSongIntoPlaylistStatment, playlistId, songId)
    if err != nil {
        return err
    }
    if tag.RowsAffected() == 0 {
        return errors.New("Error saving song")
    }
    return nil
}

/*
Saves song belonging to album
*/
func SaveSong(albumID uuid.UUID, platformSong models.PlatformSong) (models.Song, error) {
    var song models.Song
    var songPlatformId uuid.UUID
    err := Pool.QueryRow(context.Background(), insertNewSongStatment, platformSong.Platform, albumID, platformSong.ID, platformSong.Title, platformSong.MediaURL, platformSong.PreviewURL).Scan(&song.ID, &songPlatformId)
    if err != nil {
        return song, err
    }
    return song, nil
}

/*
Saves playlist, saves songs, and then saves the playlist songs
*/
func SaveFullPlaylistAndSongs(platformPlaylist models.PlatformPlaylist) (models.Playlist, error) {
    _, independentPlaylist, err := SavePlaylist(platformPlaylist)
    if err != nil {
        return independentPlaylist, err
    }
    for _, song := range platformPlaylist.Songs {
        independentSong, err := SaveFullSong(song)
        if err != nil {
            return independentPlaylist, err
        }
        err = SavePlaylistSong(independentPlaylist.ID, independentSong.ID)
        if err != nil {
            return independentPlaylist, err
        }
    }
    return independentPlaylist, nil
}

/*
Saves multple full playlists
*/
func SaveFullPlaylists(playlists []models.PlatformPlaylist) ([]models.Playlist, error) {
    independentPlaylists := make([]models.Playlist, len(playlists))
    for i, playlist := range playlists {
        independentPlaylist, err := SaveFullPlaylistAndSongs(playlist)
        if err != nil {
            return independentPlaylists, err
        }
        independentPlaylists[i] = independentPlaylist
    }
    return independentPlaylists, nil
}

/*
Saves album, saves artists, saves artists album, and then saves all the songs
*/
func SaveFullAlbum(platformAlbum models.PlatformAlbum) (models.Album, error) {
    _, independentAlbum, err := SaveAlbum(platformAlbum)
    if err != nil {
        return independentAlbum, err
    }
    for _, artist := range platformAlbum.Artists {
        _, independentArtist, err := SaveArtist(artist)
        if err != nil {
            return independentAlbum, err
        }
        err = SaveArtistAlbum(independentArtist.ID, independentAlbum.ID)
        if err != nil {
            return independentAlbum, err
        }
    }
    for _, song := range platformAlbum.Songs {
        _, err := SaveSong(independentAlbum.ID, song)
        if err != nil {
            return independentAlbum, err
        }
    }
    return independentAlbum, nil
}

/*
Saves multiple full albums
*/
func SaveFullAlbums(albums []models.PlatformAlbum) ([]models.Album, error) {
    independentAlbums := make([]models.Album, len(albums))
    for i, album := range albums {
        independentAlbum, err := SaveFullAlbum(album)
        if err != nil {
            return independentAlbums, err
        }
        independentAlbums[i] = independentAlbum
    }
    return independentAlbums, nil
}

/*
Saves album, saves artist, saves artists album, and then the song
*/
func SaveFullSong(platformSong models.PlatformSong) (models.Song, error) {
    _, independentAlbum, err := SaveAlbum(platformSong.Album)
    if err != nil {
        return models.Song{}, err
    }
    for _, artist := range platformSong.Album.Artists {
        _, independentArtist, err := SaveArtist(artist)
        if err != nil {
            return models.Song{}, err
        }
        err = SaveArtistAlbum(independentArtist.ID, independentAlbum.ID)
        if err != nil {
            return models.Song{}, err
        }
    }
    independentSong, err := SaveSong(independentAlbum.ID, platformSong)
    if err != nil {
        return models.Song{}, err
    }
    return independentSong, nil
}

/*
Saves multiple full songs
*/
func SaveFullSongs(songs []models.PlatformSong) error {
    for _, song := range songs {
        _, err := SaveFullSong(song)
        if err != nil {
            return err
        }
    }
    return nil
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
            err := Pool.QueryRow(context.Background(), insertArtistStatment, "spotify", artist.ID.String(), artist.Name, album.Images[0].URL).Scan(&artistID, &artistPlatformId)

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
            err := Pool.QueryRow(context.Background(), insertSongStatment, "spotify", albumId, track.ID.String(), track.Name, album.Images[0].URL, track.PreviewURL).Scan(&songId, &songPlatformId)
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
            err := Pool.QueryRow(context.Background(), insertAlbumStatment, "spotify", track.Album.ID.String(), track.Album.Name, track.Album.Images[0].URL).Scan(&albumId, &albumPlatformId)
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
            err = Pool.QueryRow(context.Background(), insertSongStatment, "spotify", albumId, track.ID.String(), track.Name, "", track.PreviewURL).Scan(&songId, &songPlatformId)
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
