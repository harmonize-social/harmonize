package models

import (
    "time"

    "github.com/google/uuid"
)

/*
Represents what the frontend sends to the backend
when a user posts a new post.
*/
type NewPost struct {
    Caption string `json:"caption"`
    Platform string `json:"platform"`
    Type string `json:"type"`
    PlatformSpecificId string `json:"id"`
}

/*
Represents a post that is returned to the frontend
*/
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
    Comments []RootComment `json:"comments"`
}

/*
Represents a root comment which can have child comments
*/
type RootComment struct {
    ID uuid.UUID `json:"id"`
    Username string `json:"username"`
    Message string `json:"message"`
    Replies []Comment `json:"replies"`
    CreatedAt time.Time `json:"createdAt"`
}

/*
Represents a child comment which is a reply to a root comment,
this can never be a root comment itself
*/
type ChildComment struct {
    ID uuid.UUID `json:"id"`
    Username string `json:"username"`
    Message string `json:"message"`
    CreatedAt time.Time `json:"createdAt"`
}

/*
Represents a comment in the database
*/
type Comment struct {
    ID uuid.UUID `json:"id"`
    PostId uuid.UUID `json:"postId"`
    ReplyToId uuid.UUID `json:"replyToId"`
    Username string `json:"username"`
    Message string `json:"message"`
    CreatedAt time.Time `json:"createdAt"`
}

/*
Represents a saved post in the database
*/
type SavedPost struct {
    ID uuid.UUID `json:"id"`
    UserId uuid.UUID `json:"userId"`
    PostId uuid.UUID `json:"postId"`
}
