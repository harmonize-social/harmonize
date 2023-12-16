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

type Session struct {
	ID uuid.UUID `json:"id"`
	UserId uuid.UUID `json:"user_id"`
	Expiry time.Time `json:"expiry"`
}

type Follow struct {
	ID uuid.UUID `json:"id"`
	FollowedId uuid.UUID `json:"followed_id"`
	FollowerId uuid.UUID `json:"follower_id"`
	Date time.Time `json:"date"`
}