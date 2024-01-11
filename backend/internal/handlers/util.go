package handlers

import (
    "backend/internal/auth"
    "backend/internal/models"
    "net/http"
    "strconv"

    "github.com/google/uuid"
)

/*
Extracts the limit and offset from the request URL query and returns them as integers.
*/
func GetLimitOffsetSession(r *http.Request) (int, int, models.User, error) {
    limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
    if err != nil {
        limit = 10
    }
    offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
    if err != nil {
        offset = 0
    }
    user, err := GetUserFromSession(r)
    if err != nil {
        return 0, 0, user, err
    }
    return limit, offset, user, nil
}

/*
Extracts the user from the request header and returns it.
*/
func GetUserFromSession(r *http.Request) (models.User, error) {
    id := uuid.MustParse(r.Header.Get("id"))
    user, err := auth.GetUserFromSession(id)
    if err != nil {
        return user, err
    }
    return user, nil
}
