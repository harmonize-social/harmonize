package platforms

import (
    "backend/internal/auth"
    "backend/internal/models"
    "backend/internal/repositories"
    "context"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "time"

    "github.com/google/uuid"
    "github.com/markbates/goth/providers/deezer"
)

type Tokens struct {
    AccessToken  string `json:"accessToken"`
    RefreshToken string `json:"refreshToken"`
}

func DeezerProvider() *deezer.Provider {
    id := repositories.DeezerClientId
    secret := repositories.DeezerSecret
    provider := deezer.New(id, secret, repositories.DeezerRedirect, "basic_access", "email", "offline_access", "manage_library", "manage_community", "delete_library", "listening_history")
    return provider
}

func DeezerURL(csrf string) (string, error) {
    provider := DeezerProvider()
    session, err := provider.BeginAuth(csrf)
    if err != nil {
        return "", err
    }
    url, err := session.GetAuthURL()
    if err != nil {
        return "", err
    }
    return url, nil
}

type DeezerAccessToken struct {
    AccessToken string `json:"access_token"`
}

func GetDeezerSession(provider *deezer.Provider, code string) (*deezer.Session, error) {
    var token DeezerAccessToken
    url := "https://connect.deezer.com/oauth/access_token.php?app_id=" + provider.ClientKey + "&secret=" + provider.Secret + "&code=" + code + "&output=json"
    response, err := provider.Client().Get(url)
    if err != nil {
        return nil, err
    }
    defer response.Body.Close()
    body, err := io.ReadAll(response.Body)
    json.Unmarshal(body, &token)
    expiresAt := time.Now().Add(time.Hour * 24 * 365 * 100)
    session := &deezer.Session{
        AuthURL:     "",
        AccessToken: token.AccessToken,
        ExpiresAt:   expiresAt,
    }
    return session, nil
}

func DeezerCallback(w http.ResponseWriter, r *http.Request) {
    state := r.URL.Query().Get("state")
    provider := DeezerProvider()
    user, err := auth.GetUserFromSession(uuid.MustParse(state))
    if err != nil {
        fmt.Println("Error: ", err.Error())
        models.Error(w, http.StatusUnauthorized, "Invalid token")
        return
    }
    code := r.URL.Query().Get("code")
    session, err := GetDeezerSession(provider, code)

    connection := &models.Connection{
        ID:           uuid.New(),
        UserID:       user.ID,
        AccessToken:  session.AccessToken,
        RefreshToken: "",
        Expiry:       time.Now().Add(time.Hour * 24 * 365 * 100),
    }
    sqlStatement := `INSERT INTO connections (id, user_id, access_token, refresh_token, expiry) VALUES ($1, $2, $3, $4, $5) RETURNING id`
    var connectionID uuid.UUID
    err = repositories.Pool.QueryRow(context.Background(),
        sqlStatement,
        connection.ID,
        connection.UserID,
        connection.AccessToken,
        connection.RefreshToken,
        connection.Expiry.Format(time.RFC3339)).Scan(&connectionID)
    if err != nil {
        fmt.Fprintf(w, "connection: %s\n\r", connection.UserID)
        fmt.Fprintf(w, "Unable to execute the query. %s", err)
    }
    sqlStatement = `
    INSERT INTO libraries (platform_id, id, connection_id) VALUES ('deezer', uuid_generate_v4(), $1) RETURNING id;
    `
    tag, err := repositories.Pool.Exec(context.Background(),
        sqlStatement,
        connectionID)
    if err != nil || tag.RowsAffected() == 0 {
        fmt.Printf("error: %v", err)
        models.Error(w, http.StatusInternalServerError, "Internal server error")
        return
    }
    models.Result(w, "Ok")
}
