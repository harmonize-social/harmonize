package models

import (
	"time"

	"github.com/google/uuid"
)

/*
Represents a user in the database.
*/
type User struct {
	ID		 uuid.UUID `json:"id,omitempty"`
	Email	 string    `json:"email"`
	Username string    `json:"username"`
	Password string    `json:"password_hash,omitempty"`
}

/*
Represents a connection in the database. A connection is a user's connection to
a third-party service, such as Spotify. This always has a library associated.
*/
type Connection struct {
	ID			 uuid.UUID `json:"id"`
	UserID		 uuid.UUID `json:"userId"`
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
	Expiry		 time.Time `json:"expiry"`
}

/*
Represents a session in the database.
*/
type Session struct {
	ID	   uuid.UUID `json:"id"`
	UserId uuid.UUID `json:"user_id"`
	Expiry time.Time `json:"expiry"`
}

/*
Represents a follow in the database.
*/
type Follow struct {
	ID		   uuid.UUID `json:"id"`
	FollowedId uuid.UUID `json:"followed_id"`
	FollowerId uuid.UUID `json:"follower_id"`
	Date	   time.Time `json:"date"`
}
