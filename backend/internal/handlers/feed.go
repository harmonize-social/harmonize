package handlers

import (
    "backend/internal/models"
    "backend/internal/repositories"
    "context"
    "net/http"
    "strconv"

    "github.com/google/uuid"
)

func GetFeed(w http.ResponseWriter, r *http.Request) {
    limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
    if err != nil {
        limit = 10
    }
    offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
    if err != nil {
        offset = 0
    }
    sessionId := r.Header.Get("id")
    user, err := getUserFromSession(uuid.MustParse(sessionId))
    if err != nil {
        models.Error(w, http.StatusUnauthorized, "Invalid session")
        return
    }
    sqlStatement := `SELECT posts.*, users.username FROM follows JOIN posts ON follows.followed_id = posts.user_id JOIN users ON posts.user_id = users.id WHERE follows.follower_id = $1 ORDER BY posts.created_at DESC LIMIT $2 OFFSET $3;`
    rows, err := repositories.Pool.Query(context.Background(), sqlStatement, user.ID, limit, offset)
    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Error getting posts")
        return
    }
    var posts []models.Post
    for rows.Next() {
        var post models.Post
        var typeSpecificId uuid.UUID
        var userId uuid.UUID
        err = rows.Scan(&post.ID, &userId, &post.CreatedAt, &post.Caption, &post.Type, &typeSpecificId, &post.Username)
        if err != nil {
            models.Error(w, http.StatusInternalServerError, "Error getting posts")
            return
        }
        var content interface{}
        if post.Type == "playlist" {
            content, err = getPlaylist(typeSpecificId)
            post.Content = content
        } else if post.Type == "song" {
            content, err = getSong(typeSpecificId)
            post.Content = content
        } else if post.Type == "album" {
            content, err = getAlbum(typeSpecificId)
            post.Content = content
        } else if post.Type == "artist" {
            content, err = getArtist(typeSpecificId)
            post.Content = content
        }
        if err != nil {
            println(err.Error())
            models.Error(w, http.StatusInternalServerError, "Error getting posts")
            return
        }
        posts = append(posts, post)
    }
    models.Result(w, posts)
}

func getArtist(Id uuid.UUID) (models.Artist, error) {
    var artist models.Artist
    err := repositories.Pool.QueryRow(context.Background(), `SELECT * FROM artists WHERE id = $1;`, Id).Scan(&artist.ID, &artist.Name)
    return artist, err
}

func getAlbum(Id uuid.UUID) (models.Album, error) {
    var album models.Album
    err := repositories.Pool.QueryRow(context.Background(), `SELECT id, name FROM albums WHERE id = $1;`, Id).Scan(&album.ID, &album.Title)
    if err != nil {
        return album, err
    }
    var artists []models.Artist
    rows, err := repositories.Pool.Query(context.Background(), `SELECT artists.id, artists.name FROM artists JOIN artists_album ON artists.id = artists_album.artist_id WHERE artists_album.album_id = $1;`, Id)
    if err != nil {
        return album, err
    }
    for rows.Next() {
        var artist models.Artist
        err = rows.Scan(&artist.ID, &artist.Name)
        if err != nil {
            return album, err
        }
        artists = append(artists, artist)
    }
    album.Artists = artists
    rows, err = repositories.Pool.Query(context.Background(), `SELECT id, name FROM songs WHERE album_id = $1;`, Id)
    if err != nil {
        return album, err
    }
    var songs []models.Song
    for rows.Next() {
        var song models.Song
        err = rows.Scan(&song.ID, &song.Title)
        if err != nil {
            return album, err
        }
        songs = append(songs, song)
    }
    album.Songs = songs
    return album, nil
}

func getPlaylist(Id uuid.UUID) (models.Playlist, error) {
    playlist := models.Playlist{}
    err := repositories.Pool.QueryRow(context.Background(), `SELECT * FROM playlists WHERE id = $1;`, Id).Scan(&playlist.ID, &playlist.Title)
    if err != nil {
        return playlist, err
    }
    rows, err := repositories.Pool.Query(context.Background(), `SELECT songs.id, songs.album_id, songs.name FROM songs JOIN playlist_songs ON songs.id = playlist_songs.song_id WHERE playlist_songs.playlist_id = $1;`, Id)
    if err != nil {
        return playlist, err
    }
    songs := []models.Song{}
    for rows.Next() {
        var albumId uuid.UUID
        var song models.Song
        err = rows.Scan(&song.ID, &albumId, &song.Title)
        rows2, err := repositories.Pool.Query(context.Background(), `SELECT artists.id, artists.name FROM artists JOIN artists_album ON artists.id = artists_album.artist_id WHERE artists_album.album_id = $1;`, Id)
        for rows2.Next() {
            var artist models.Artist
            err = rows2.Scan(&artist.ID, &artist.Name)
            if err != nil {
                return playlist, err
            }
            song.Artists = append(song.Artists, artist)
        }
        songs = append(songs, song)
    }
    playlist.Songs = songs
    return playlist, nil
}

func getSong(Id uuid.UUID) (models.Song, error) {
    song := models.Song{}
    var albumId uuid.UUID
    err := repositories.Pool.QueryRow(context.Background(), `SELECT id, album_id, name FROM songs WHERE id = $1;`, Id).Scan(&song.ID, &albumId, &song.Title)
    if err != nil {
        return song, err
    }
    rows, err := repositories.Pool.Query(context.Background(), `SELECT artists.id, artists.name FROM artists JOIN artists_album ON artists.id = artists_album.artist_id WHERE artists_album.album_id = $1;`, Id)
    artists := []models.Artist{}
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

/*
CREATE TABLE IF NOT EXISTS artists(
    id UUID PRIMARY KEY,
    name VARCHAR(128) NOT NULL
);

CREATE TABLE IF NOT EXISTS artists_album(
    id UUID PRIMARY KEY,
    artist_id UUID REFERENCES artists (id) NOT NULL,
    album_id UUID REFERENCES albums (id) NOT NULL
);

CREATE TABLE IF NOT EXISTS albums(
    id UUID PRIMARY KEY,
    artist_id UUID NOT NULL REFERENCES artists (id),
    title VARCHAR(128) NOT NULL
);


CREATE TABLE IF NOT EXISTS playlists(
    id UUID PRIMARY KEY,
    title VARCHAR(128) NOT NULL
);


CREATE TABLE IF NOT EXISTS songs(
    id UUID PRIMARY KEY,
    title VARCHAR(128) NOT NULL,
    album_id UUID REFERENCES albums (id) NOT NULL
);


CREATE TABLE IF NOT EXISTS playlist_songs(
    id UUID PRIMARY KEY,
    playlist_id UUID REFERENCES playlists (id) NOT NULL,
    song_id UUID REFERENCES songs (id) NOT NULL
);

*/
