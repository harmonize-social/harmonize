package models

import (
    "time"

    "github.com/google/uuid"
)

type User struct {
    ID uuid.UUID `json:"id"`
    Email string `json:"email"`
    Username string `json:"username"`
    Password string `json:"password_hash"`
}

type Connection struct {
    ID uuid.UUID `json:"id"`
    UserID uuid.UUID `json:"userId"`
    AccessToken string `json:"accessToken"`
    RefreshToken string `json:"refreshToken"`
    Expiry time.Time `json:"expiry"`
}
