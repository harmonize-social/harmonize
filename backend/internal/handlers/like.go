package handlers

import (
    "backend/internal/models" // models package where User schema is defined
    "backend/internal/repositories"

    "context"
    "database/sql"
    "encoding/json" // package to encode and decode the json into struct and vice versa
    "fmt"
    "log"
    "net/http" // used to access the request and response object of the api

    "github.com/google/uuid" // uuid
    "github.com/gorilla/mux" // used to get the params from the route
    _ "github.com/lib/pq" // likegres golang driver
)

func CreateLike(w http.ResponseWriter, r *http.Request) {
    // Set headers
    setCommonHeaders(w)
    setAdditionalHeaders(w, "POST")

    // create an empty like of type models.Like
    var like models.Like

    // decode the json request to like
    err := json.NewDecoder(r.Body).Decode(&like)

    if err != nil {
        log.Fatalf("Unable to decode the request body.  %v", err)
    }

    // call insert user function and pass the like
    insertID := insertLike(like)

    // format a response object
    res := response{
        ID:      insertID.String(),
        Message: "Like created successfully",
    }

    // send the response
    json.NewEncoder(w).Encode(res)
}

// Will return a single like by its id
func GetLike(w http.ResponseWriter, r *http.Request) {

    setCommonHeaders(w)

    // get the likeid from the request params
    params := mux.Vars(r)

    // convert the id type from string to uuid.UUID
    likeID, err := uuid.Parse(params["id"])

    if err != nil {
        log.Fatalf("Unable to parse the UUID. %v", err)
    }

    // call the getLike function with likeID to retrieve a single like
    like, err := getLike(likeID)

    if err != nil {
        log.Fatalf("Unable to get like. %v", err)
    }

    // send the response
    json.NewEncoder(w).Encode(like)
}

// Deletes like in the postgres db
func DeleteLike(w http.ResponseWriter, r *http.Request) {

    setCommonHeaders(w)
    setAdditionalHeaders(w, "DELETE")

    // get the likeid from the request params, key is "id"
    params := mux.Vars(r)

    // convert the id in string to uuid.UUID
    likeID, err := uuid.Parse(params["id"])

    if err != nil {
        log.Fatalf("Unable to parse the UUID.  %v", err)
    }

    // call deletePost
    deletedRows := deleteLike(likeID)

    // format the message string
    msg := fmt.Sprintf("Like deleted successfully. Total rows/record affected %v", deletedRows)

    // format the reponse message
    res := response{
        ID:      likeID.String(),
        Message: msg,
    }

    // send the response
    json.NewEncoder(w).Encode(res)
}

// ------------------------- handler functions ----------------
// insert one like in the DB
func insertLike(like models.Like) uuid.UUID {

    // create the insert sql query
    // will return the id of the inserted like
    sqlStatement := `INSERT INTO likes (id, post_id, user_id) VALUES ($1, $2, $3) RETURNING id`

    // generate a new UUID for the like
    likeID := uuid.New()

    // execute the sql statement
    err := repositories.Pool.QueryRow(context.Background(), sqlStatement, likeID, like.PostId, like.UserId).Scan(&likeID)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    fmt.Printf("Inserted a single record %v", likeID)

    // return the inserted id
    return likeID
}

// get one like from the DB by its likeID
func getLike(likeID uuid.UUID) (models.Like, error) {

    // create a like of models.Like type
    var like models.Like

    // create the select sql query
    sqlStatement := `SELECT * FROM likes WHERE id=$1`

    // execute the sql statement
    row := repositories.Pool.QueryRow(context.Background(), sqlStatement, likeID)

    // unmarshal the row object to like
    err := row.Scan(&like.ID, &like.PostId, &like.UserId)

    switch err {
    case sql.ErrNoRows:
        fmt.Println("No rows were returned!")
        return like, nil
    case nil:
        return like, nil
    default:
        log.Fatalf("Unable to scan the row. %v", err)
    }

    // return empty like on error
    return like, err
}

// delete like in the DB
func deleteLike(likeID uuid.UUID) int64 {

    // create the delete sql query
    sqlStatement := `DELETE FROM likes WHERE id=$1`

    // execute the sql statement
    res, err := repositories.Pool.Exec(context.Background(), sqlStatement, likeID)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    // check how many rows affected
    rowsAffected := res.RowsAffected()

    fmt.Printf("Total rows/record affected %v", rowsAffected)

    return rowsAffected
}
