package handlers

import (
    "backend/internal/auth"
    "backend/internal/models"
    "backend/internal/repositories"
    "context"
    "encoding/json"
    "fmt"
    "net/http" // used to access the request and response object of the api
    "strconv"

    "github.com/google/uuid"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
    limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
    if err != nil {
        limit = 10
    }
    offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
    if err != nil {
        offset = 0
    }
    sessionId := r.Header.Get("id")
    user, err := auth.GetUserFromSession(uuid.MustParse(sessionId))
    if err != nil {
        models.Error(w, http.StatusUnauthorized, "Invalid session")
        return
    }

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
        sqlStatement := `SELECT id, post_id, username, reply_to_id, message, created_at FROM comments WHERE post_id = $1 ORDER BY created_at DESC`
        rows, err := repositories.Pool.Query(context.Background(), sqlStatement, post.ID)
        if err != nil {
            models.Error(w, http.StatusInternalServerError, "Failed to get comments")
            return
        }
        defer rows.Close()

        comments := make([]models.Comment, 0)
        for rows.Next() {
            var comment models.Comment
            err = rows.Scan(&comment.ID, &comment.PostId, &comment.Username, &comment.ReplyToId, &comment.Message, &comment.CreatedAt)
            if err != nil {
                models.Error(w, http.StatusInternalServerError, "Failed to get comments")
                return
            }
            comments = append(comments, comment)
        }

        processedComments := make([]models.RootComment, 0)
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
        post, err := repositories.CreatePlaylistPost(playlist, user.ID, newPost.Caption)
        if err != nil {
            models.Error(w, http.StatusInternalServerError, "Internal server error")
            fmt.Println(err.Error())
            fmt.Println("playlist2:", err.Error())
            return
        }
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
        post, err := repositories.CreateSongPost(song, user.ID, newPost.Caption)
        if err != nil {
            models.Error(w, http.StatusInternalServerError, "Internal server error")
            fmt.Println("song2:", err.Error())
            return
        }
        models.Result(w, post)
        break
    case "album":
        album, err := repositories.GetAlbum("spotify", newPost.PlatformSpecificId)
        if err != nil {
            models.Error(w, http.StatusInternalServerError, "Internal server error")
            fmt.Println("album1:", err.Error())
            return
        }
        post, err := repositories.CreateAlbumPost(album, user.ID, newPost.Caption)
        if err != nil {
            models.Error(w, http.StatusInternalServerError, "Internal server error")
            fmt.Println("album2:", err.Error())
            return
        }
        models.Result(w, post)
        break
    case "artist":
        artist, err := repositories.GetArtist("spotify", newPost.PlatformSpecificId)
        if err != nil {
            models.Error(w, http.StatusInternalServerError, "Internal server error")
            fmt.Println("artist1:", err.Error())
            return
        }
        post, err := repositories.CreateArtistPost(artist, user.ID, newPost.Caption)
        if err != nil {
            models.Error(w, http.StatusInternalServerError, "Internal server error")
            fmt.Println("artist2:", err.Error())
            return
        }
        models.Result(w, post)
        break
    }
}

func GetUserPosts(w http.ResponseWriter, r *http.Request) {
    limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
    if err != nil {
        limit = 10
    }
    offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
    if err != nil {
        offset = 0
    }
    sessionId := r.Header.Get("id")
    user, err := auth.GetUserFromSession(uuid.MustParse(sessionId))
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
        sqlStatement := `SELECT id, post_id, username, reply_to_id, message, created_at FROM comments WHERE post_id = $1 ORDER BY created_at DESC`
        rows, err := repositories.Pool.Query(context.Background(), sqlStatement, post.ID)
        if err != nil {
            models.Error(w, http.StatusInternalServerError, "Failed to get comments")
            fmt.Println(err.Error())
            return
        }
        defer rows.Close()

        comments := make([]models.Comment, 0)
        for rows.Next() {
            var comment models.Comment
            err = rows.Scan(&comment.ID, &comment.PostId, &comment.Username, &comment.ReplyToId, &comment.Message, &comment.CreatedAt)
            if err != nil {
                models.Error(w, http.StatusInternalServerError, "Failed to get comments")
                fmt.Println(err.Error())
                return
            }
            comments = append(comments, comment)
        }

        processedComments := make([]models.RootComment, 0)
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
