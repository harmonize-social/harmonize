package repositories

import (
    "backend/internal/models"
    "context"
    "errors"
    "fmt"

    "github.com/google/uuid"
    "github.com/zmb3/spotify/v2"
)

func SaveSpotifySongs(tracks *spotify.SavedTrackPage) ([]models.PlatformSong, error) {
    independentTracks := make([]models.PlatformSong, len(tracks.Tracks))
    insertAlbumStatment := `SELECT * FROM insert_new_album($1, $2, $3, $4);`
    insertSongStatment := `SELECT * FROM insert_new_song($1, $2, $3, $4, $5, $6);`
    insertArtistStatment := `SELECT * FROM insert_new_artist($1, $2, $3, $4);`
    for i, track := range tracks.Tracks {
        var albumId uuid.UUID
        var albumPlatformId string
        err := Pool.QueryRow(context.Background(), insertAlbumStatment, "spotify", track.Album.ID.String(), track.Album.Name, track.Album.Images[0].URL).Scan(&albumId, &albumPlatformId)
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
