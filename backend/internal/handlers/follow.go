package handlers

import (
    "backend/internal/models"
    "backend/internal/repositories"
    "context"
    "net/http"

    "github.com/google/uuid"
)

func PostFollow(w http.ResponseWriter, r *http.Request) {
    id := uuid.MustParse(r.Header.Get("id"))

    user, err := getUserFromSession(id)
    if err != nil {
        models.Error(w, http.StatusInternalServerError, "cannot get session")
        return
    }
    username := r.URL.Query().Get("username")

    if username == "" {
        models.Error(w, http.StatusBadRequest, "username is empty")
        return
    }

    sqlStatement := `INSERT INTO follows (id, followed_id, follower_id, date)
        VALUES (uuid_generate_v4(),
        (SELECT id FROM users WHERE username = $1),
        $2,
        NOW()) ON CONFLICT (followed_id, follower_id) DO NOTHING;`

    tag, err := repositories.Pool.Exec(context.Background(), sqlStatement, username, user.ID)

    if err != nil {
        models.Error(w, http.StatusInternalServerError, "cannot follow")
        return
    }

    if tag.RowsAffected() == 0 {
        models.Error(w, http.StatusBadRequest, "already following")
        return
    }
}

func DeleteFollow(w http.ResponseWriter, r *http.Request) {
    id := uuid.MustParse(r.Header.Get("id"))

    user, err := getUserFromSession(id)
    if err != nil {
        models.Error(w, http.StatusInternalServerError, "cannot get session")
        return
    }
    username := r.URL.Query().Get("username")

    if username == "" {
        models.Error(w, http.StatusBadRequest, "username is empty")
        return
    }

    sqlStatement := `DELETE FROM follows WHERE followed_id = (SELECT id FROM users WHERE username = $1) AND follower_id = $2;`

    tag, err := repositories.Pool.Exec(context.Background(), sqlStatement, username, user.ID)

    if err != nil {
        models.Error(w, http.StatusInternalServerError, "cannot unfollow")
        return
    }

    if tag.RowsAffected() == 0 {
        models.Error(w, http.StatusBadRequest, "not following")
        return
    }
}
