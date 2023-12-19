package handlers

import (
    "backend/internal/models" // models package where User schema is defined
    "backend/internal/repositories"
    "context"
    "time"

    "database/sql"
    "encoding/json" // package to encode and decode the json into struct and vice versa
    "fmt"
    "log"
    "net/http" // used to access the request and response object of the api

    "github.com/alexedwards/argon2id"
    "github.com/google/uuid" // uuid
    "github.com/gorilla/mux" // used to get the params from the route
    "github.com/jackc/pgx/v4"
)

type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func CreateSession(w http.ResponseWriter, r *http.Request) {
    // Set headers
    setCommonHeaders(w)
    setAdditionalHeaders(w, "POST")

    var loginRequest LoginRequest
    err := json.NewDecoder(r.Body).Decode(&loginRequest)

    if err != nil {
        fmt.Fprintf(w, "Error: %s", err)
        return
    }

    if err != nil {
        fmt.Fprintf(w, "Hash error: %s", err)
        return
    }

    // get user from db with password, email and/or username
    sqlStatement := `SELECT * FROM users WHERE username=$1 OR email=$1`
    var user models.User
    err = repositories.Pool.QueryRow(context.Background(), sqlStatement, loginRequest.Username).Scan(&user)

    if err == pgx.ErrNoRows {
        fmt.Fprintf(w, "Username or Password wrong: %s", err)
        return
    }

    valid, err := argon2id.ComparePasswordAndHash(loginRequest.Password, user.Password)

    if err != nil {
        fmt.Fprintf(w, "Hash error: %s", err)
        return
    }

    if !valid {
        fmt.Fprintf(w, "Username or Password wrong: %s", err)
        return
    }

    if err != nil {
        fmt.Fprintf(w, "Error DB: %s", err)
        return
    }

    if user.ID == uuid.Nil {
        fmt.Fprintf(w, "Username or Password wrong: %s", err)
        return
    }

    // insert session in the DB and return the newly inserted session
    sqlStatement = `INSERT INTO sessions (id, user_id, expiry) VALUES ($1, $2, $3) RETURNING id`

    // generate a new UUID for the session
    sessionID := uuid.New()

    err = repositories.Pool.QueryRow(context.Background(), sqlStatement, sessionID, user.ID, time.Now().AddDate(0, 0, 7)).Scan(&sessionID)

    if err != nil {
        fmt.Fprintf(w, "Error DB: %s", err)
        return
    }

    frontendSession := models.FrontendSession{
        ID:      sessionID,
        Expiry:  time.Now().AddDate(0, 0, 7),
    }

    res := models.ApiResponse{
        Value: frontendSession,
    }

    json.NewEncoder(w).Encode(res)
}

// GetSession will return a single session by its id
func GetSession(w http.ResponseWriter, r *http.Request) {

    setCommonHeaders(w)

    // get the id from the request params
    params := mux.Vars(r)

    // convert the id type from string to uuid.UUID
    sessionID, err := uuid.Parse(params["id"])

    if err != nil {
        log.Fatalf("Unable to parse the UUID. %v", err)
    }

    // call the getSession function with sessionID to retrieve a single session
    session, err := getSession(sessionID)

    if err != nil {
        log.Fatalf("Unable to get session. %v", err)
    }

    // send the response
    json.NewEncoder(w).Encode(session)
}

// Updates the session's details in the postgres db
func UpdateSession(w http.ResponseWriter, r *http.Request) {

    setCommonHeaders(w)
    setAdditionalHeaders(w, "PUT")

    // get the id from the request params, key is "id"
    params := mux.Vars(r)

    // convert the id type from string to uuid.UUID
    sessionID, err := uuid.Parse(params["id"])

    if err != nil {
        log.Fatalf("Unable to parse the UUID. %v", err)
    }

    // create an empty session of type models.Session
    var session models.Session

    // decode the json request of the session
    err = json.NewDecoder(r.Body).Decode(&session)

    if err != nil {
        log.Fatalf("Unable to decode the request body.  %v", err)
    }

    // call updateSession
    updatedRows := updateSession(sessionID, session)

    // format the message string
    msg := fmt.Sprintf("Session updated successfully. Total rows/record affected %v", updatedRows)

    // format the response message
    res := response{
        ID:      sessionID.String(),
        Message: msg,
    }

    // send the response
    json.NewEncoder(w).Encode(res)
}

// Deletes session in the postgres db
func DeleteSession(w http.ResponseWriter, r *http.Request) {

    setCommonHeaders(w)
    setAdditionalHeaders(w, "DELETE")

    // get the id from the request params, key is "id"
    params := mux.Vars(r)

    // convert the id in string to uuid.UUID
    sessionID, err := uuid.Parse(params["id"])

    if err != nil {
        log.Fatalf("Unable to parse the UUID.  %v", err)
    }

    // call deleteSession
    deletedRows := deleteSession(sessionID)

    // format the message string
    msg := fmt.Sprintf("Session deleted successfully. Total rows/record affected %v", deletedRows)

    // format the reponse message
    res := response{
        ID:      sessionID.String(),
        Message: msg,
    }

    // send the response
    json.NewEncoder(w).Encode(res)
}

// ------------------------- handler functions ----------------
// insert one session in the DB
func insertSession(session models.Session) uuid.UUID {

    // create the insert sql query
    // will return the id of the inserted session
    sqlStatement := `INSERT INTO sessions (id, user_id, expiry) VALUES ($1, $2, $3) RETURNING id`

    // generate a new UUID for the session
    sessionID := uuid.New()

    // execute the sql statement
    err := repositories.Pool.QueryRow(context.Background(), sqlStatement, sessionID, session.UserId, session.Expiry).Scan(&sessionID)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    fmt.Printf("Inserted a single record %v", sessionID)

    // return the inserted id
    return sessionID
}

// get one session from the DB by its sessionID
func getSession(sessionID uuid.UUID) (models.Session, error) {

    // create a session of models.Session type
    var session models.Session

    // create the select sql query
    sqlStatement := `SELECT * FROM sessions WHERE id=$1`

    // execute the sql statement
    row := repositories.Pool.QueryRow(context.Background(), sqlStatement, sessionID)

    // unmarshal the row object to session
    err := row.Scan(&session.ID, &session.UserId, &session.Expiry)

    switch err {
    case sql.ErrNoRows:
        fmt.Println("No rows were returned!")
        return session, nil
    case nil:
        return models.Session{}, nil
    default:
        log.Fatalf("Unable to scan the row. %v", err)
    }

    // return empty session on error
    return session, err
}

// update session in the DB
func updateSession(sessionID uuid.UUID, session models.Session) int64 {

    // create the update sql query
    sqlStatement := `UPDATE sessions SET user_id=$2, expiry=$3 WHERE id=$1`

    // execute the sql statement
    res, err := repositories.Pool.Exec(context.Background(), sqlStatement, sessionID, session.UserId, session.Expiry)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    // check how many rows affected
    rowsAffected := res.RowsAffected() // how to see/respond if there is an eror with checking the rows?

    fmt.Printf("Total rows/record affected %v", rowsAffected)

    return rowsAffected
}

// delete session in the DB
func deleteSession(sessionID uuid.UUID) int64 {

    // create the delete sql query
    sqlStatement := `DELETE FROM sessions WHERE id=$1`

    // execute the sql statement
    res, err := repositories.Pool.Exec(context.Background(), sqlStatement, sessionID)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    // check how many rows affected
    rowsAffected := res.RowsAffected()

    fmt.Printf("Total rows/record affected %v", rowsAffected)

    return rowsAffected
}
