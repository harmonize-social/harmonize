package repositories

import (
    "backend/internal/models"
    "context"

    "github.com/google/uuid"
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
