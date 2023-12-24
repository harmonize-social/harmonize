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
