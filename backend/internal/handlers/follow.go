package handlers

import (
    "backend/internal/auth"
    "backend/internal/models"
    "backend/internal/repositories"
    "context"
    "fmt"
    "net/http"
    "strconv"

    "github.com/google/uuid"
)

/*
Follow a user

POST /follow?username=<username>
*/
func PostFollow(w http.ResponseWriter, r *http.Request) {
    id := uuid.MustParse(r.Header.Get("id"))

    user, err := auth.GetUserFromSession(id)
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

    models.Result(w, "OK");
}

/*
Unfollow a user

DELETE /follow?username=<username>
*/
func DeleteFollow(w http.ResponseWriter, r *http.Request) {
    id := uuid.MustParse(r.Header.Get("id"))

    user, err := auth.GetUserFromSession(id)
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

    models.Result(w, "OK");
}

/*
Get usernames that follow a user

GET /followers?limit=<limit>&offset=<offset>
*/
func GetFollowers(w http.ResponseWriter, r *http.Request) {
    limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
    if err != nil {
        limit = 10
    }

    offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
    if err != nil {
        offset = 0
    }

    id := uuid.MustParse(r.Header.Get("id"))

    user, err := auth.GetUserFromSession(id)

    if err != nil {
        models.Error(w, http.StatusInternalServerError, "cannot get session")
        return
    }

    sqlStatement := `SELECT u.username FROM follows f JOIN users u ON f.follower_id = u.id WHERE f.followed_id = $1 LIMIT $2 OFFSET $3;`

    rows, err := repositories.Pool.Query(context.Background(), sqlStatement, user.ID, limit, offset)

    if err != nil {
        models.Error(w, http.StatusInternalServerError, "cannot get followers")
        return
    }

    defer rows.Close()

    users := make([]string, 0)

    for rows.Next() {
        username := "NULL"
        err = rows.Scan(&username)
        if err != nil {
            models.Error(w, http.StatusInternalServerError, "cannot get followers")
            return
        }
        users = append(users, username)
    }

    fmt.Println(users)
    models.Result(w, users)
}

/*
Get usernames that a user is following

GET /following?limit=<limit>&offset=<offset>
*/
func GetFollowing(w http.ResponseWriter, r *http.Request) {
    limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
    if err != nil {
        limit = 10
    }

    offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
    if err != nil {
        offset = 0
    }

    id := uuid.MustParse(r.Header.Get("id"))

    user, err := auth.GetUserFromSession(id)

    if err != nil {
        models.Error(w, http.StatusInternalServerError, "cannot get session")
        return
    }

    sqlStatement := `SELECT username FROM users WHERE id IN (SELECT followed_id FROM follows WHERE follower_id = (SELECT id FROM users WHERE username = $1)) LIMIT $2 OFFSET $3;`

    rows, err := repositories.Pool.Query(context.Background(), sqlStatement, user.Username, limit, offset)

    if err != nil {
        models.Error(w, http.StatusInternalServerError, "cannot get following")
        fmt.Println(err)
        return
    }

    defer rows.Close()

    users := make([]string, 0)

    for rows.Next() {
        var username string
        err = rows.Scan(&username)
        if err != nil {
            fmt.Println(err)
            return
        }
        users = append(users, username)
    }

    models.Result(w, users)
}
