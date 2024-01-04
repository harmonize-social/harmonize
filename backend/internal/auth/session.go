package auth

import (
    "backend/internal/models"
    "backend/internal/repositories"
    "context"

    "github.com/google/uuid"
)

func GetUserFromSession(sessionID uuid.UUID) (models.User, error) {
    var user models.User

    sqlStatement := `SELECT users.id, users.email, users.username, users.password_hash FROM sessions LEFT JOIN users ON users.id = sessions.user_id WHERE sessions.id = $1;`
    row := repositories.Pool.QueryRow(context.Background(), sqlStatement, sessionID)
    err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Password)

    return user, err
}
