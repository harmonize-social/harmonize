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

func GetComments(w http.ResponseWriter, r *http.Request) {
    postId := uuid.MustParse(r.URL.Query().Get("id"))
    if postId == uuid.Nil {
        models.Error(w, http.StatusBadRequest, "Invalid post id")
        return
    }

    comments, err := repositories.GetPostComments(postId)

    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Failed to get comments")
        fmt.Println("after get comments", err)
        return
    }

    models.Result(w, comments)
}

func CreateComment(w http.ResponseWriter, r *http.Request) {
    sessionId := r.Header.Get("id")
    user, err := auth.GetUserFromSession(uuid.MustParse(sessionId))
    if err != nil {
        models.Error(w, http.StatusUnauthorized, "Invalid session")
        fmt.Println("after session", err)
        return
    }

    var comment models.Comment
    err = json.NewDecoder(r.Body).Decode(&comment)
    if err != nil {
        models.Error(w, http.StatusBadRequest, "Invalid request body")
        return
    }

    comment.Username = user.Username
    if comment.PostId == uuid.Nil {
        models.Error(w, http.StatusBadRequest, "Invalid post id")
        return
    } else if comment.Message == "" {
        models.Error(w, http.StatusBadRequest, "Invalid message")
        return
    }

    var newComment models.Comment
    if comment.ReplyToId != uuid.Nil {
        sqlStatement := `INSERT INTO comments (id, post_id, username, reply_to_id, message, created_at) VALUES (uuid_generate_v4(), $1, $2, $3, $4, now()) RETURNING id, post_id, username, reply_to_id, message, created_at`
        err = repositories.Pool.QueryRow(context.Background(), sqlStatement, comment.PostId, comment.Username, comment.ReplyToId, comment.Message).Scan(&newComment.ID, &newComment.PostId, &newComment.Username, &newComment.ReplyToId, &newComment.Message, &newComment.CreatedAt)
    } else {
        sqlStatement := `INSERT INTO comments (id, post_id, username, message, created_at) VALUES (uuid_generate_v4(), $1, $2, $3, now()) RETURNING id, post_id, username, message, created_at`
        err = repositories.Pool.QueryRow(context.Background(), sqlStatement, comment.PostId, comment.Username, comment.Message).Scan(&newComment.ID, &newComment.PostId, &newComment.Username, &newComment.Message, &newComment.CreatedAt)
    }

    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Failed to create comment")
        fmt.Println("after insert", err)
        return
    }

    models.Result(w, newComment)
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
    sessionId := r.Header.Get("id")
    user, err := auth.GetUserFromSession(uuid.MustParse(sessionId))
    if err != nil {
        models.Error(w, http.StatusUnauthorized, "Invalid session")
        return
    }

    commentId := uuid.MustParse(r.URL.Query().Get("id"))
    if err != nil {
        models.Error(w, http.StatusBadRequest, "Invalid comment id")
        return
    }

    sqlStatement := `DELETE FROM comments WHERE id = $1 AND username = $2`
    _, err = repositories.Pool.Exec(context.Background(), sqlStatement, commentId, user.Username)
    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Failed to delete comment")
        return
    }

    models.Result(w, "Comment deleted")
}
