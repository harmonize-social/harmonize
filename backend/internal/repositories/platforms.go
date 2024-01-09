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
		ID:			  uuid.New(),
		UserID:		  user_id,
		AccessToken:  access_token,
		RefreshToken: refresh_token,
		Expiry:		  expiry,
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

func GetTokens(platform string, user_id uuid.UUID) (models.Tokens, error) {
	var tokens models.Tokens
	sqlStatement := `SELECT c.access_token, c.refresh_token, c.expiry FROM connections c INNER JOIN libraries l ON c.id = l.connection_id WHERE l.platform_id = $1 AND c.user_id = $2`
	err := Pool.QueryRow(context.Background(), sqlStatement, user_id).Scan(&tokens.AccessToken, &tokens.RefreshToken, &tokens.Expiry)
	if err != nil {
		return tokens, err
	}
	return tokens, nil
}
