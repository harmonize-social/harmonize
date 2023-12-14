package models

import (
	"github.com/google/uuid"
)

type Post struct {
	ID uuid.UUID `json:"id"`
	UserId uuid.UUID `json:"user_id"`
	Caption string `json:"caption"`
	Type string `json:"type"`
	TypeSpecificId uuid.UUID `json:"type_specific_id"`
}

type Like struct {
	ID uuid.UUID `json:"id"`
	PostId uuid.UUID `json:"post_id"`
	UserId uuid.UUID `json:"user_id"`
}

type Comment struct {
	ID uuid.UUID `json:"id"`
	PostId uuid.UUID `json:"post_id"`
	UserId uuid.UUID `json:"user_id"`
	ReplyToId uuid.UUID `json:"reply_to_id"`
	Message string `json:"message"`
}

type SavedPost struct {
	ID uuid.UUID `json:"id"`
	UserId uuid.UUID `json:"user_id"`
	PostId uuid.UUID `json:"post_id"`
}