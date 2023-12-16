package handlers

import (
    "backend/internal/models"
    "backend/internal/repositories"

    "context"
    "database/sql"
    "encoding/json" // package to encode and decode the json into struct and vice versa
    "fmt"
    "github.com/google/uuid"   // uuid
    "github.com/gorilla/mux"   // used to get the params from the route
    _ "github.com/lib/pq"      // postgres golang driver
    "log"
    "net/http" // used to access the request and response object of the api
)

// used https://codesource.io/build-a-crud-application-in-golang-with-postgresql/

func CreateUser(w http.ResponseWriter, r *http.Request) {
    // Set headers
    setCommonHeaders(w)
    setAdditionalHeaders(w, "POST")

    // create an empty user of type models.User
    var user models.User

    // decode the json request to user
    err := json.NewDecoder(r.Body).Decode(&user)

    if err != nil {
        log.Fatalf("Unable to decode the request body.  %v", err)
    }

    // call insert user function and pass the user
    insertID := insertUser(user)

    // format a response object
    res := response{
        ID:      insertID.String(),
        Message: "User created successfully",
    }

    // send the response
    json.NewEncoder(w).Encode(res)
}

// Will return a single user by its id
func GetUser(w http.ResponseWriter, r *http.Request) {

    setCommonHeaders(w)

    // get the userid from the request params
    params := mux.Vars(r)

    // convert the id type from string to uuid.UUID
    userID, err := uuid.Parse(params["id"])

    if err != nil {
        log.Fatalf("Unable to parse the UUID. %v", err)
    }

    // call the getUser function with user id to retrieve a single user
    user, err := getUser(userID)

    if err != nil {
        log.Fatalf("Unable to get user. %v", err)
    }

    // send the response
    json.NewEncoder(w).Encode(user)
}

// Update user's detail in the postgres db
func UpdateUser(w http.ResponseWriter, r *http.Request) {

    setCommonHeaders(w)
    setAdditionalHeaders(w, "PUT")

    // get the userid from the request params, key is "id"
    params := mux.Vars(r)

    // convert the id type from string to uuid.UUID
    userID, err := uuid.Parse(params["id"])

    if err != nil {
        log.Fatalf("Unable to convert the string into int.  %v", err)
    }

    // create an empty user of type models.User
    var user models.User

    // decode the json request to user
    err = json.NewDecoder(r.Body).Decode(&user)

    if err != nil {
        log.Fatalf("Unable to decode the request body.  %v", err)
    }

    // call update user to update the user
    updatedRows := updateUser(userID, user)

    // format the message string
    msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", updatedRows)

    // format the response message
    res := response{
        ID:      userID.String(),
        Message: msg,
    }

    // send the response
    json.NewEncoder(w).Encode(res)
}

// Delete user's detail in the postgres db
func DeleteUser(w http.ResponseWriter, r *http.Request) {

    setCommonHeaders(w)
    setAdditionalHeaders(w, "DELETE")

    // get the userid from the request params, key is "id"
    params := mux.Vars(r)

    // convert the id in string to uuid.UUID
    userID, err := uuid.Parse(params["id"])

    if err != nil {
        log.Fatalf("Unable to parse the UUID.  %v", err)
    }

    // call deleteUser, convert the int
    deletedRows := deleteUser(userID)

    // format the message string
    msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", deletedRows)

    // format the reponse message
    res := response{
        ID:      userID.String(),
        Message: msg,
    }

    // send the response
    json.NewEncoder(w).Encode(res)
}

// ------------------------- handler functions ----------------
// insert one user in the DB
func insertUser(user models.User) uuid.UUID {

    // create the postgres db connection
    db, errDB := repositories.CreateConnection()
    if errDB != nil {
        log.Fatalf("Error creating database connection: %v", errDB)
    }

    // close the db connection
    defer db.Close()

    // Create the insert sql query
    // Will return the id of the inserted user
    sqlStatement := `INSERT INTO users (id, email, username, password_hash) VALUES ($1, $2, $3, $4) RETURNING id`

    // generate a new UUID for the user
    userID := uuid.New()

    // execute the sql statement
    err := db.QueryRow(context.Background(), sqlStatement, userID, user.Email, user.Username, user.Password).Scan(&userID)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    fmt.Printf("Inserted a single record %v", userID)

    // return the inserted id
    return userID
}

// get one user from the DB by its userid
func getUser(userID uuid.UUID) (models.User, error) {
    // create the postgres db connection
    db, errDB := repositories.CreateConnection()
    if errDB != nil {
        log.Fatalf("Error creating database connection: %v", errDB)
    }

    // close the db connection
    defer db.Close()

    // create a user of models.User type
    var user models.User

    // create the select sql query
    sqlStatement := `SELECT * FROM users WHERE id=$1`

    // execute the sql statement
    row := db.QueryRow(context.Background(), sqlStatement, userID)

    // unmarshal the row object to user
    err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Password)

    switch err {
    case sql.ErrNoRows:
        fmt.Println("No rows were returned!")
        return user, nil
    case nil:
        return user, nil
    default:
        log.Fatalf("Unable to scan the row. %v", err)
    }

    // return empty user on error
    return user, err
}

// update user in the DB
func updateUser(userID uuid.UUID, user models.User) int64 {

    // create the postgres db connection
    db, errDB := repositories.CreateConnection()
    if errDB != nil {
        log.Fatalf("Error creating database connection: %v", errDB)
    }

    // close the db connection
    defer db.Close()

    // create the update sql query
    sqlStatement := `UPDATE users SET email=$2, username=$3, password_hash=$4 WHERE id=$1`

    // execute the sql statement
    res, err := db.Exec(context.Background(), sqlStatement, userID, user.Email, user.Username, user.Password)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    // check how many rows affected
    rowsAffected := res.RowsAffected()

    fmt.Printf("Total rows/record affected %v", rowsAffected)

    return rowsAffected
}

// delete user in the DB
func deleteUser(userID uuid.UUID) int64 {

    // create the postgres db connection
    db, errDB := repositories.CreateConnection()
    if errDB != nil {
        log.Fatalf("Error creating database connection: %v", errDB)
    }

    // close the db connection
    defer db.Close()

    // create the delete sql query
    sqlStatement := `DELETE FROM users WHERE id=$1`

    // execute the sql statement
    res, err := db.Exec(context.Background(), sqlStatement, userID)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    // check how many rows affected
    rowsAffected := res.RowsAffected()

    fmt.Printf("Total rows/record affected %v", rowsAffected)

    return rowsAffected
}
