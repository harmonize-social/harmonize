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

type RootComment struct {
    ID uuid.UUID `json:"id"`
    Username string `json:"username"`
    Message string `json:"message"`
    Replies []Comment `json:"replies"`
    CreatedAt time.Time `json:"createdAt"`
}

type ChildComment struct {
    ID uuid.UUID `json:"id"`
    Username string `json:"username"`
    Message string `json:"message"`
    CreatedAt time.Time `json:"createdAt"`
}

type Comment struct {
    ID uuid.UUID `json:"id"`
    PostId uuid.UUID `json:"postId"`
    ReplyToId uuid.UUID `json:"replyToId"`
    Username string `json:"username"`
    Message string `json:"message"`
    CreatedAt time.Time `json:"createdAt"`
}

type SavedPost struct {
    ID uuid.UUID `json:"id"`
    UserId uuid.UUID `json:"userId"`
    PostId uuid.UUID `json:"postId"`
}
