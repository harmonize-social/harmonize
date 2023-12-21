package handlers

import (
    "backend/internal/models"
    "backend/internal/repositories"
    "context"
    "fmt"
    "net/http"
    "strconv"

    "github.com/google/uuid"
)

func GetSavedPosts(w http.ResponseWriter, r *http.Request) {
    limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
    if err != nil {
        limit = 10
    }
    offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
    if err != nil {
        offset = 0
    }

    id := uuid.MustParse(r.Header.Get("id"))

    user, err := getUserFromSession(id)
    if err != nil {
        models.Error(w, http.StatusInternalServerError, "cannot get session")
        return
    }

    fmt.Println(user.ID)

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

    var posts []models.Post
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
        posts = append(posts, post)
    }
    models.Result(w, posts)
}

/*
CREATE TABLE IF NOT EXISTS saved_posts(
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES posts (id) NOT NULL,
    post_id UUID REFERENCES posts (id) NOT NULL
);
*/

func PostSavedPost(w http.ResponseWriter, r *http.Request) {
    id := r.Header.Get("id")
    fmt.Println(id)
}
