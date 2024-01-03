package handlers

import (
    "backend/internal/auth"
    "backend/internal/models"
    "backend/internal/repositories"
    "context"
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/google/uuid"
)

func GetFeed(w http.ResponseWriter, r *http.Request) {
    limit, offset, user, err := GetLimitOffsetSession(r)
    if err != nil {
        models.Error(w, http.StatusUnauthorized, "Invalid session")
        return
    }
    sqlStatement := `SELECT posts.id, posts.created_at, posts.caption, posts.type, posts.type_specific_id, users.username,
                        COUNT(liked_posts.post_id) AS like_count,
                        COUNT(CASE WHEN liked_posts.user_id = $1 THEN 1 END) > 0 AS user_has_liked,
                        COUNT(CASE WHEN saved_posts.user_id = $1 THEN 1 END) > 0 AS user_has_saved
                     FROM follows
                     JOIN posts ON follows.followed_id = posts.user_id
                     JOIN users ON posts.user_id = users.id
                     LEFT JOIN liked_posts ON liked_posts.post_id = posts.id
                     LEFT JOIN saved_posts ON saved_posts.post_id = posts.id
                     WHERE follows.follower_id = $1
                     GROUP BY posts.id, users.id
                     ORDER BY posts.created_at DESC
                     LIMIT $2
                     OFFSET $3`
    rows, err := repositories.Pool.Query(context.Background(), sqlStatement, user.ID, limit, offset)
    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Error getting posts")
        return
    }
    posts := make([]models.Post, 0)
    for rows.Next() {
        var post models.Post
        var typeSpecificId uuid.UUID
        err = rows.Scan(&post.ID, &post.CreatedAt, &post.Caption, &post.Type, &typeSpecificId, &post.Username, &post.LikeCount, &post.HasLiked, &post.HasSaved)
        if err != nil {
            println(err.Error())
            models.Error(w, http.StatusInternalServerError, "Error getting posts")
            return
        }

        content, err := repositories.GetPostContent(post.Type, typeSpecificId)

        if err != nil {
            println(err.Error())
            models.Error(w, http.StatusInternalServerError, "Error getting posts")
            return
        }

        post.Content = content

        processedComments, err := repositories.GetPostComments(post.ID)

        if err != nil {
            models.Error(w, http.StatusInternalServerError, "Error getting posts")
            return
        }

        post.Comments = processedComments

        posts = append(posts, post)
    }
    models.Result(w, posts)
}

func GetMePosts(w http.ResponseWriter, r *http.Request) {
    limit, offset, user, err := GetLimitOffsetSession(r)

    sqlStatement := `SELECT posts.id, posts.created_at, posts.caption, posts.type, posts.type_specific_id, users.username,
                         COUNT(liked_posts.post_id) AS like_count,
                         COUNT(CASE WHEN liked_posts.user_id = $1 THEN 1 END) > 0 AS user_has_liked,
                         COUNT(CASE WHEN saved_posts.user_id = $1 THEN 1 END) > 0 AS user_has_saved
                     FROM posts
                     JOIN users ON posts.user_id = users.id
                     LEFT JOIN liked_posts ON liked_posts.post_id = posts.id
                     LEFT JOIN saved_posts ON saved_posts.post_id = posts.id
                     WHERE posts.user_id = $1
                     GROUP BY posts.id, users.id
                     ORDER BY posts.created_at DESC
                     LIMIT $2
                     OFFSET $3`
    rows, err := repositories.Pool.Query(context.Background(), sqlStatement, user.ID, limit, offset)

    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Error getting posts")
        return
    }

    posts := make([]models.Post, 0)
    for rows.Next() {
        var post models.Post
        var typeSpecificId uuid.UUID
        err = rows.Scan(&post.ID, &post.CreatedAt, &post.Caption, &post.Type, &typeSpecificId, &post.Username, &post.LikeCount, &post.HasLiked, &post.HasSaved)
        if err != nil {
            models.Error(w, http.StatusInternalServerError, "Error getting posts")
            return
        }
        content, err := repositories.GetPostContent(post.Type, typeSpecificId)

        if err != nil {
            models.Error(w, http.StatusInternalServerError, "Error getting posts")
            return
        }

        post.Content = content

        processedComments, err := repositories.GetPostComments(post.ID)

        if err != nil {
            models.Error(w, http.StatusInternalServerError, "Error getting posts")
            return
        }

        post.Comments = processedComments

        posts = append(posts, post)
    }

    models.Result(w, posts)
}

func NewPost(w http.ResponseWriter, r *http.Request) {
    sessionId := r.Header.Get("id")
    user, err := auth.GetUserFromSession(uuid.MustParse(sessionId))
    if err != nil {
        models.Error(w, http.StatusUnauthorized, "Invalid session")
        return
    }

    var newPost models.NewPost
    err = json.NewDecoder(r.Body).Decode(&newPost)
    if err != nil {
        models.Error(w, http.StatusBadRequest, "Invalid post")
        return
    }
    fmt.Println(newPost)

    if newPost.Type != "playlist" && newPost.Type != "song" && newPost.Type != "album" && newPost.Type != "artist" {
        models.Error(w, http.StatusBadRequest, "Invalid post type")
        return
    }

    if newPost.Platform != "spotify" {
        models.Error(w, http.StatusNotImplemented, "Invalid platform")
        return
    }

    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Internal server error")
        return
    }

    switch newPost.Type {
    case "playlist":
        playlist, err := repositories.GetPlaylist("spotify", newPost.PlatformSpecificId)
        if err != nil {
            models.Error(w, http.StatusInternalServerError, "Internal server error")
            fmt.Println("playlist1:", err.Error())
            return
        }
        post, err := repositories.CreatePost(user.ID, newPost.Caption, "playlist", playlist.ID)
        if err != nil {
            models.Error(w, http.StatusInternalServerError, "Internal server error")
            fmt.Println(err.Error())
            fmt.Println("playlist2:", err.Error())
            return
        }
        post.Content = playlist
        models.Result(w, post)
        break
    case "song":
        fmt.Println(newPost.PlatformSpecificId)
        song, err := repositories.GetSong("spotify", newPost.PlatformSpecificId)
        if err != nil {
            models.Error(w, http.StatusInternalServerError, "Internal server error")
            fmt.Println("song1:", err.Error())
            return
        }
        post, err := repositories.CreatePost(user.ID, newPost.Caption, "song", song.ID)
        if err != nil {
            models.Error(w, http.StatusInternalServerError, "Internal server error")
            fmt.Println("song2:", err.Error())
            return
        }
        post.Content = song
        models.Result(w, post)
        break
    case "album":
        album, err := repositories.GetAlbum("spotify", newPost.PlatformSpecificId)
        if err != nil {
            models.Error(w, http.StatusInternalServerError, "Internal server error")
            fmt.Println("album1:", err.Error())
            return
        }
        post, err := repositories.CreatePost(user.ID, newPost.Caption, "album", album.ID)
        if err != nil {
            models.Error(w, http.StatusInternalServerError, "Internal server error")
            fmt.Println("album2:", err.Error())
            return
        }
        post.Content = album
        models.Result(w, post)
        break
    case "artist":
        artist, err := repositories.GetArtist("spotify", newPost.PlatformSpecificId)
        if err != nil {
            models.Error(w, http.StatusInternalServerError, "Internal server error")
            fmt.Println("artist1:", err.Error())
            return
        }
        post, err := repositories.CreatePost(user.ID, newPost.Caption, "artist", artist.ID)
        if err != nil {
            models.Error(w, http.StatusInternalServerError, "Internal server error")
            fmt.Println("artist2:", err.Error())
            return
        }
        post.Content = artist
        models.Result(w, post)
        break
    }
}

func GetUserPosts(w http.ResponseWriter, r *http.Request) {
    limit, offset, user, err := GetLimitOffsetSession(r)
    if err != nil {
        models.Error(w, http.StatusUnauthorized, "Invalid session")
        return
    }

    username := r.URL.Query().Get("username")

    if username == "" {
        models.Error(w, http.StatusBadRequest, "Invalid username")
        return
    }

    // Get the users that have the username specified, and username follows user.id

    sqlStatement := `SELECT id FROM users WHERE username = $1 AND id IN (SELECT follower_id FROM follows WHERE followed_id = $2)`

    var followedId uuid.UUID
    err = repositories.Pool.QueryRow(context.Background(), sqlStatement, username, user.ID).Scan(&followedId)
    if err != nil {
        models.Error(w, http.StatusInternalServerError, "You are not followed by this user")
        fmt.Println(err.Error())
        return
    }

    sqlStatement = `SELECT posts.id, posts.created_at, posts.caption, posts.type, posts.type_specific_id, users.username,
                         COUNT(liked_posts.post_id) AS like_count,
                         COUNT(CASE WHEN liked_posts.user_id = users.id THEN 1 END) > 0 AS user_has_liked,
                         COUNT(CASE WHEN saved_posts.user_id = users.id THEN 1 END) > 0 AS user_has_saved
                     FROM posts
                     JOIN users ON posts.user_id = users.id
                     LEFT JOIN liked_posts ON liked_posts.post_id = posts.id
                     LEFT JOIN saved_posts ON saved_posts.post_id = posts.id
                     WHERE users.username = $1
                     GROUP BY posts.id, users.id
                     ORDER BY posts.created_at DESC
                     LIMIT $2
                     OFFSET $3`
    rows, err := repositories.Pool.Query(context.Background(), sqlStatement, username, limit, offset)

    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Error getting posts")
        fmt.Println("here1: ", err.Error())
        return
    }

    posts := make([]models.Post, 0)
    for rows.Next() {
        var post models.Post
        var typeSpecificId uuid.UUID
        err = rows.Scan(&post.ID, &post.CreatedAt, &post.Caption, &post.Type, &typeSpecificId, &post.Username, &post.LikeCount, &post.HasLiked, &post.HasSaved)
        if err != nil {
            models.Error(w, http.StatusInternalServerError, "Error getting posts")
            fmt.Println("here2: ", err.Error())
            return
        }

        content, err := repositories.GetPostContent(post.Type, typeSpecificId)

        if err != nil {
            models.Error(w, http.StatusInternalServerError, "Error getting posts")
            fmt.Println("here3: ", err.Error())
            return
        }

        post.Content = content

        processedComments, err := repositories.GetPostComments(post.ID)

        if err != nil {
            models.Error(w, http.StatusInternalServerError, "Error getting posts")
            return
        }

        post.Comments = processedComments

        posts = append(posts, post)
    }

    models.Result(w, posts)
}

/*

CREATE TABLE IF NOT EXISTS follows(
    id UUID PRIMARY KEY,
    followed_id UUID REFERENCES users (id) NOT NULL,
    follower_id UUID REFERENCES users (id) NOT NULL,
    CONSTRAINT follows_unique_user_combo UNIQUE (followed_id, follower_id),
    date timestamptz NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    email VARCHAR(1024) unique NOT NULL,
    username VARCHAR(1024) unique NOT NULL,
    password_hash VARCHAR(1024) NOT NULL
);

*/
