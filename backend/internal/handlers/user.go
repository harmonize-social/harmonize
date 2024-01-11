package handlers

import (
    "backend/internal/models" // models package where User schema is defined
    "backend/internal/repositories"
    "context"
    "fmt"
    "time"

    "encoding/json" // package to encode and decode the json into struct and vice versa
    "net/http"      // used to access the request and response object of the api

    "github.com/alexedwards/argon2id"
    "github.com/golang-jwt/jwt"
    "github.com/google/uuid" // uuid
    "github.com/jackc/pgx/v4"
)

/*
Returned when a user logs in or registers
*/
type TokenResponse struct {
    Token string `json:"token"`
}

/*
Representation of a user's login request
*/
type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

/*
Login handler

POST /users/login

{
    "username": <username|email>,
    "password": "password"
}
*/
func Login(w http.ResponseWriter, r *http.Request) {
    var loginRequest LoginRequest
    err := json.NewDecoder(r.Body).Decode(&loginRequest)

    if err != nil {
        models.Error(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    // get user from db with password, email and/or username
    sqlStatement := `SELECT * FROM users WHERE username=$1 OR email=$1`
    var user models.User
    err = repositories.Pool.QueryRow(context.Background(), sqlStatement, loginRequest.Username).Scan(&user.ID, &user.Email, &user.Username, &user.Password)

    if err == pgx.ErrNoRows {
        models.Error(w, http.StatusUnauthorized, "Username/Email or Password wrong")
        return
    }

    valid, err := argon2id.ComparePasswordAndHash(loginRequest.Password, user.Password)

    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Error comparing password")
        return
    }

    if !valid {
        models.Error(w, http.StatusUnauthorized, "Username/Email or Password wrong")
        return
    }

    sessionID, err := insertSession(user.ID)

    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Error creating session")
    }

    t, err := generateJWT(sessionID)

    models.Result(w, &TokenResponse{
        Token: t,
    })
}

/*
Representation of a user's registration request
*/
type RegisterRequest struct {
    Email    string `json:"email"`
    Username string `json:"username"`
    Password string `json:"password"`
}

/*
Register handler

POST /users/register

{
    "email": "email",
    "username": "username",
    "password": "password"
}
*/
func Register(w http.ResponseWriter, r *http.Request) {
    // create an empty user of type models.User
    var registerRequest RegisterRequest

    // decode the json request to user
    err := json.NewDecoder(r.Body).Decode(&registerRequest)

    if err != nil {
        models.Error(w, http.StatusBadRequest, "Invalid request payload")
    }

    var user models.User

    user.Email = registerRequest.Email
    user.Username = registerRequest.Username

    user.Password, err = argon2id.CreateHash(user.Password, argon2id.DefaultParams)
    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Error creating account")
    }

    sqlStatement := `INSERT INTO users (id, email, username, password_hash) VALUES ($1, $2, $3, $4) RETURNING id`
    userId := uuid.New()

    err = repositories.Pool.QueryRow(context.Background(), sqlStatement, userId, user.Email, user.Username, user.Password).Scan(&userId)

    if err != nil {
        fmt.Println(err)
        models.Error(w, http.StatusInternalServerError, "Username or Email already exists")
        return
    }

    sessionID, err := insertSession(userId)

    t, err := generateJWT(sessionID)

    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Error creating session")
    }

    models.Result(w, &TokenResponse{
        Token: t,
    })
}

/*
User's info handler
*/
func GetUserInfo(w http.ResponseWriter, r *http.Request) {
    user, err := GetUserFromSession(r)

    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Error getting user")
        return
    }

    user.Password = ""

    models.Result(w, user)
}

/*
Create a session for a user in the database
*/
func insertSession(userId uuid.UUID) (uuid.UUID, error) {

    sqlStatement := `INSERT INTO sessions (id, user_id, expiry) VALUES ($1, $2, $3) RETURNING id`
    sessionID := uuid.New()
    err := repositories.Pool.QueryRow(context.Background(), sqlStatement, sessionID, userId, time.Now().AddDate(0, 0, 7)).Scan(&sessionID)

    return sessionID, err
}

/*
Generate a JWT for a session
*/
func generateJWT(sessionID uuid.UUID) (string, error) {
    t := jwt.NewWithClaims(jwt.SigningMethodHS256,
        jwt.MapClaims{
            "id":     sessionID.String(),
            "expiry": time.Now().AddDate(0, 0, 7).String(),
        },
    )
    token, err := t.SignedString(repositories.Secret)

    return token, err
}
