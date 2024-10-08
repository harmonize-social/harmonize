package repositories

import (
    "backend/internal/models"
    "context"
    "errors"

    "github.com/google/uuid"
)

const (
    GetPostByIDQuery          = "SELECT id, user_id, caption, type, created_at, type_specific_id FROM posts WHERE id = $1;"
    GetPostPlatformByIDQuery  = `SELECT
                                    COALESCE(
                                        ps_song.platform_id,
                                        ps_album.platform_id,
                                        ps_playlist.platform_id,
                                        ps_artist.platform_id
                                    ) AS platform_id,
                                    COALESCE(
                                        ps_song.platform_specific_id,
                                        ps_album.platform_specific_id,
                                        ps_playlist.platform_specific_id,
                                        ps_artist.platform_specific_id
                                    ) AS platform_specific_id,
                                    CASE
                                        WHEN p.type = 'song' THEN 'song'
                                        WHEN p.type = 'album' THEN 'album'
                                        WHEN p.type = 'playlist' THEN 'playlist'
                                        WHEN p.type = 'artist' THEN 'artist'
                                        ELSE 'unknown'
                                    END AS type
                                FROM
                                    posts p
                                LEFT JOIN
                                    platform_songs ps_song ON p.type = 'song' AND p.type_specific_id = ps_song.song_id
                                LEFT JOIN
                                    platform_albums ps_album ON p.type = 'album' AND p.type_specific_id = ps_album.album_id
                                LEFT JOIN
                                    platform_playlists ps_playlist ON p.type = 'playlist' AND p.type_specific_id = ps_playlist.playlist_id
                                LEFT JOIN
                                    platform_artists ps_artist ON p.type = 'artist' AND p.type_specific_id = ps_artist.artist_id
                                WHERE
                                    p.id = $1;`
    GetSongByIDQuery          = "SELECT id, name, album_id, media_url, preview_url FROM songs WHERE id = $1;"
    GetAlbumByIDQuery         = "SELECT id, name, media_url FROM albums WHERE id = $1;"
    GetAlbumSongsByIDQuery    = "SELECT id, name, album_id, media_url, preview_url FROM songs WHERE album_id = $1;"
    GetAlbumArtistsByIDQuery  = "SELECT id, name, media_url FROM artists WHERE id IN (SELECT artist_id FROM artists_album WHERE album_id = $1);"
    GetArtistByIDQuery        = "SELECT id, name, media_url FROM artists WHERE id = $1;"
    GetPlaylistByIDQuery      = "SELECT id, name, media_url FROM playlists WHERE id = $1;"
    GetPlaylistSongsByIDQuery = "SELECT id, name, album_id, media_url, preview_url FROM songs WHERE id IN (SELECT song_id FROM playlist_songs WHERE playlist_id = $1);"
    GetCommentsByPostIDQuery  = "SELECT id, post_id, username, reply_to_id, created_at, message FROM comments WHERE post_id = $1;"
)

/*
Create a new post in the database
*/
func CreatePost(userID uuid.UUID, caption string, postType string, typeSpecificID uuid.UUID) (models.Post, error) {
    var post models.Post

    sqlStatement := "INSERT INTO posts (id, user_id, caption, type, created_at, type_specific_id) VALUES (uuid_generate_v4() ,$1, $2, $3, now(), $4) RETURNING id, created_at, caption, type;"
    err := Pool.QueryRow(context.Background(), sqlStatement, userID, caption, postType, typeSpecificID).Scan(&post.ID, &post.CreatedAt, &post.Caption, &post.Type)

    if err != nil {
        return post, err
    }

    return post, nil
}

/*
Generalised platform model for when we don't care about the type's other fields
*/
type PlatformItem struct {
    Platform string
    ID       string
    Type     string
}

/*
Get a post's platform items so that we can save them in a user's library
*/
func GetPostPlatform(postID uuid.UUID) ([]PlatformItem, error) {
    platforms := make([]PlatformItem, 0)
    rows, err := Pool.Query(context.Background(), GetPostPlatformByIDQuery, postID)

    if err != nil {
        return platforms, err
    }

    for rows.Next() {
        var platformItem PlatformItem
        err = rows.Scan(&platformItem.Platform, &platformItem.ID, &platformItem.Type)

        if err != nil {
            return platforms, err
        }

        platforms = append(platforms, platformItem)
    }

    if len(platforms) == 0 {
        return platforms, errors.New("No platform found")
    }

    return platforms, nil
}

/*
Used for populating a post's content
*/
func GetPostContent(typeName string, typeSpecificID uuid.UUID) (interface{}, error) {
    var content interface{}
    var err error

    switch typeName {
    case "song":
        content, err = GetFullPostSong(typeSpecificID)
    case "artist":
        content, err = GetFullPostArtist(typeSpecificID)
    case "album":
        content, err = GetFullPostAlbum(typeSpecificID)
    case "playlist":
        content, err = GetFullPostPlaylist(typeSpecificID)
    default:
        err = errors.New("Invalid type name")
    }

    return content, err
}

/*
Get standalone song, populate album artists and album
*/
func GetFullPostSong(typeSpecificID uuid.UUID) (models.Song, error) {
    var song models.Song

    albumId := uuid.UUID{}
    err := Pool.QueryRow(context.Background(), GetSongByIDQuery, typeSpecificID).Scan(&song.ID, &song.Title, &albumId, &song.MediaURL, &song.PreviewURL)

    if err != nil {
        return song, err
    }

    var album models.Album
    err = Pool.QueryRow(context.Background(), GetAlbumByIDQuery, albumId).Scan(&album.ID, &album.Title, &album.MediaURL)

    if err != nil {
        return song, err
    }

    artists, err := GetAlbumArtists(albumId)

    album.Artists = artists
    song.Album = album
    return song, nil
}

/*
Get album songs, don't populate albums, artists for each song
*/
func GetAlbumSongs(albumID uuid.UUID) ([]models.Song, error) {
    songs := make([]models.Song, 0)
    rows, err := Pool.Query(context.Background(), GetAlbumSongsByIDQuery, albumID)
    defer rows.Close()

    if err != nil {
        return songs, err
    }

    for rows.Next() {
        var song models.Song
        var albumId uuid.UUID
        err = rows.Scan(&song.ID, &song.Title, &albumId, &song.MediaURL, &song.PreviewURL)

        if err != nil {
            return songs, err
        }

        songs = append(songs, song)
    }

    return songs, nil
}

/*
Get song which is a part of larger object, don't populate album and artists
*/
func GetPartialSong(typeSpecificID uuid.UUID) (models.Song, error) {
    var song models.Song
    var albumId uuid.UUID

    err := Pool.QueryRow(context.Background(), GetSongByIDQuery, typeSpecificID).Scan(&song.ID, &song.Title, &albumId, &song.MediaURL, &song.PreviewURL)

    if err != nil {
        return song, err
    }

    return song, nil
}

/*
Get album artists
*/
func GetAlbumArtists(albumID uuid.UUID) ([]models.Artist, error) {
    artists := make([]models.Artist, 0)
    rows, err := Pool.Query(context.Background(), GetAlbumArtistsByIDQuery, albumID)
    defer rows.Close()

    if err != nil {
        return artists, err
    }

    for rows.Next() {
        var artist models.Artist
        err = rows.Scan(&artist.ID, &artist.Name, &artist.MediaURL)

        if err != nil {
            return artists, err
        }

        artists = append(artists, artist)
    }

    return artists, nil
}

/*
Get standalone artist (this type has no songs or albums)
*/
func GetFullPostArtist(typeSpecificID uuid.UUID) (models.Artist, error) {
    var artist models.Artist
    err := Pool.QueryRow(context.Background(), GetArtistByIDQuery, typeSpecificID).Scan(&artist.ID, &artist.Name, &artist.MediaURL)

    if err != nil {
        return artist, err
    }

    return artist, nil
}

/*
Get album with artists
*/
func GetAlbumWithArtists(albumID uuid.UUID) (models.Album, error) {
    var album models.Album
    err := Pool.QueryRow(context.Background(), GetAlbumByIDQuery, albumID).Scan(&album.ID, &album.Title, &album.MediaURL)

    if err != nil {
        return album, err
    }

    artists, err := GetAlbumArtists(albumID)

    if err != nil {
        return album, err
    }

    album.Artists = artists

    return album, nil
}

/*
Get standalone album, populate artists, populate songs
*/
func GetFullPostAlbum(typeSpecificID uuid.UUID) (models.Album, error) {
    album, err := GetAlbumWithArtists(typeSpecificID)

    if err != nil {
        return album, err
    }

    songs, err := GetAlbumSongs(typeSpecificID)

    if err != nil {
        return album, err
    }

    album.Songs = songs

    return album, nil
}

/*
Get playlist songs
*/
func GetPlaylistSongs(playlistID uuid.UUID) ([]models.Song, error) {
    songs := make([]models.Song, 0)
    rows, err := Pool.Query(context.Background(), GetPlaylistSongsByIDQuery, playlistID)
    defer rows.Close()

    if err != nil {
        return songs, err
    }

    for rows.Next() {
        var song models.Song
        var albumId uuid.UUID
        err = rows.Scan(&song.ID, &song.Title, &albumId, &song.MediaURL, &song.PreviewURL)

        if err != nil {
            return songs, err
        }

        album, err := GetAlbumWithArtists(albumId)

        if err != nil {
            return songs, err
        }

        song.Album = album

        songs = append(songs, song)
    }
    return songs, nil
}

/*
Get standalone playlist, populate songs, and their albums and artists
*/
func GetFullPostPlaylist(typeSpecificID uuid.UUID) (models.Playlist, error) {
    var playlist models.Playlist
    err := Pool.QueryRow(context.Background(), GetPlaylistByIDQuery, typeSpecificID).Scan(&playlist.ID, &playlist.Title, &playlist.MediaURL)

    if err != nil {
        return playlist, err
    }

    songs, err := GetPlaylistSongs(typeSpecificID)

    if err != nil {
        return playlist, err
    }

    playlist.Songs = songs

    return playlist, nil
}

/*
Get comments for post id
*/
func GetPostComments(postID uuid.UUID) ([]models.RootComment, error) {
    comments := make([]models.Comment, 0)
    processedComments := make([]models.RootComment, 0)
    rows, err := Pool.Query(context.Background(), GetCommentsByPostIDQuery, postID)
    defer rows.Close()

    if err != nil {
        return processedComments, err
    }

    for rows.Next() {
        var comment models.Comment
        err = rows.Scan(&comment.ID, &comment.PostId, &comment.Username, &comment.ReplyToId, &comment.CreatedAt, &comment.Message)

        if err != nil {
            return processedComments, err
        }

        comments = append(comments, comment)
    }

    for _, comment := range comments {
        if comment.ReplyToId == uuid.Nil {
            processedComments = append(processedComments, models.RootComment{
                ID:        comment.ID,
                Username:  comment.Username,
                Message:   comment.Message,
                CreatedAt: comment.CreatedAt,
                Replies:   []models.Comment{},
            })
        }
    }
    for _, comment := range comments {
        if comment.ReplyToId != uuid.Nil {
            for i, rootComment := range processedComments {
                if rootComment.ID == comment.ReplyToId {
                    processedComments[i].Replies = append(processedComments[i].Replies, comment)
                }
            }
        }
    }

    return processedComments, nil
}
