package handlers

import (
    "backend/internal/models"
    "backend/internal/repositories"
    "context"
    "net/http" // used to access the request and response object of the api
    "strconv"

    "github.com/google/uuid"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {

    setCommonHeaders(w)
    setAdditionalHeaders(w, "GET")

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

    sqlStatement := `SELECT posts.*, users.username FROM posts JOIN users ON posts.user_id = users.id WHERE user_id = $1 ORDER BY posts.created_at DESC LIMIT $2 OFFSET $3;`
    rows, err := repositories.Pool.Query(context.Background(), sqlStatement, user.ID, limit, offset)

    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Error getting posts")
        return
    }

    posts := make([]models.Post, 0)
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
