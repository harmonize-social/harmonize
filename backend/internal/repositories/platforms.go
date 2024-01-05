package repositories

import (
    "backend/internal/models"
    "context"
    "errors"
    "time"

    "github.com/google/uuid"
)

func CreateConnectionAndLibrary(user_id uuid.UUID, platform string, access_token string, refresh_token string, expiry time.Time) (err error) {
    connection := &models.Connection{
        ID:           uuid.New(),
        UserID:       user_id,
        AccessToken:  access_token,
        RefreshToken: refresh_token,
        Expiry:       expiry,
    }
    sqlStatement := `INSERT INTO connections (id, user_id, access_token, refresh_token, expiry) VALUES ($1, $2, $3, $4, $5) RETURNING id`
    var connectionID uuid.UUID
    err = Pool.QueryRow(context.Background(),
        sqlStatement,
        connection.ID,
        connection.UserID,
        connection.AccessToken,
        connection.RefreshToken,
        connection.Expiry.Format(time.RFC3339)).Scan(&connectionID)
    if err != nil {
        return err
    }
    sqlStatement = `
    INSERT INTO libraries (platform_id, id, connection_id) VALUES ($1, uuid_generate_v4(), $2) RETURNING id;
    `
    tag, err := Pool.Exec(context.Background(),
        sqlStatement,
        platform,
        connectionID)
    if err != nil  {
        return err
    }

    if tag.RowsAffected() == 0 {
        return errors.New("Unable to create library")
    }
    return nil
}
