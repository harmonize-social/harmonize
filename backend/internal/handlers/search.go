package handlers

import (
    "backend/internal/models"
    "backend/internal/repositories"
    "context"
    "fmt"
    "net/http"
    "strconv"
)

/*
Search users by username

GET /search?username=string&limit=<int>&offset=<int>
*/
func Search(w http.ResponseWriter, r *http.Request) {
    limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
    if err != nil {
        limit = 10
    }
    offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
    if err != nil {
        offset = 0
    }

    username := r.URL.Query().Get("username")

    if username == "" {
        models.Error(w, http.StatusBadRequest, "Username is required")
        return
    }

    sqlStatement := `SELECT username
                     FROM users
                     WHERE SIMILARITY(username, $1) >= 0.4
                     ORDER BY SIMILARITY(username, $1) DESC
                     LIMIT $2 OFFSET $3;`

    rows, err := repositories.Pool.Query(context.Background(), sqlStatement, username, limit, offset)

    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Internal server error")
        fmt.Println(err)
        return
    }

    usernames := make([]string, 0)
    for rows.Next() {
        var username string
        err = rows.Scan(&username)
        if err != nil {
            fmt.Println(err)
            models.Error(w, http.StatusInternalServerError, "Internal server error")
        }
        usernames = append(usernames, username)
    }

    models.Result(w, usernames)
}
