package handlers

import (
    "backend/internal/auth"
    "backend/internal/models"
    "net/http"
    "strconv"

    "github.com/google/uuid"
)

func GetLimitOffsetSession(r *http.Request) (int, int, models.User, error) {
    limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
    if err != nil {
        limit = 10
    }
    offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
    if err != nil {
        offset = 0
    }
    id := uuid.MustParse(r.Header.Get("id"))
    _, err = auth.GetUserFromSession(id)
    var user models.User
    if err != nil {
        return 0, 0, user, err
    }
    user, err = auth.GetUserFromSession(id)
    if err != nil {
        return 0, 0, user, err
    }
    return limit, offset, user, nil
}
