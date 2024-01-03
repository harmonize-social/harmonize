package repositories

import (
    "backend/internal/models"
    "context"

    "github.com/google/uuid"
)

func CreatePost(userID uuid.UUID, caption string, postType string, typeSpecificID uuid.UUID) (models.Post, error) {
    var post models.Post

    sqlStatement := "INSERT INTO posts (id, user_id, caption, type, created_at, type_specific_id) VALUES (uuid_generate_v4() ,$1, $2, $3, now(), $4) RETURNING id, created_at, caption, type;"
    err := Pool.QueryRow(context.Background(), sqlStatement, userID, caption, postType, typeSpecificID).Scan(&post.ID, &post.CreatedAt, &post.Caption, &post.Type)

    if err != nil {
        return post, err
    }

    return post, nil
}
