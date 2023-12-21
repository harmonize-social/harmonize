package models

import (
    "time"

    "github.com/google/uuid"
)

type Post struct {
    ID uuid.UUID `json:"id"`
    Username string `json:"username"`
    Caption string `json:"caption"`
    CreatedAt time.Time `json:"created_at"`
    Type string `json:"type"`
    Content interface{} `json:"content"`
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
