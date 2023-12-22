package models

import (
    "time"

    "github.com/google/uuid"
)

type Post struct {
    ID uuid.UUID `json:"id"`
    Username string `json:"username"`
    Caption string `json:"caption"`
    CreatedAt time.Time `json:"createdAt"`
    Type string `json:"type"`
    Content interface{} `json:"content"`
    LikeCount int `json:"likeCount"`
    HasLiked bool `json:"hasLiked"`
    HasSaved bool `json:"hasSaved"`
}

type Like struct {
    ID uuid.UUID `json:"id"`
    PostId uuid.UUID `json:"postId"`
    UserId uuid.UUID `json:"userId"`
}

type Comment struct {
    ID uuid.UUID `json:"id"`
    PostId uuid.UUID `json:"postId"`
    UserId uuid.UUID `json:"userId"`
    ReplyToId uuid.UUID `json:"replyToId"`
    Message string `json:"message"`
}

type SavedPost struct {
    ID uuid.UUID `json:"id"`
    UserId uuid.UUID `json:"userId"`
    PostId uuid.UUID `json:"postId"`
}
