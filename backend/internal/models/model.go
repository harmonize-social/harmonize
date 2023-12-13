package models

import (
    "github.com/google/uuid"
)

type User struct {
	ID uuid.UUID `json:"id"`
	Email string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password_hash"`
}