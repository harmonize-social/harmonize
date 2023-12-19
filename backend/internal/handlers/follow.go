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
	_ "github.com/lib/pq"    // postgres golang driver
)

func CreateFollow(w http.ResponseWriter, r *http.Request) {
    // Set headers
    setCommonHeaders(w)
    setAdditionalHeaders(w, "POST")

    // create an empty follow of type models.Follow
    var follow models.Follow

    // decode the json request to follow
    err := json.NewDecoder(r.Body).Decode(&follow)

    if err != nil {
        log.Fatalf("Unable to decode the request body.  %v", err)
    }

    // call insert user function and pass the follow
    insertID := insertFollow(follow)

    // format a response object
    res := response{
        ID:      insertID.String(),
        Message: "Follow created successfully",
    }

    // send the response
    json.NewEncoder(w).Encode(res)
}

// Will return a single follow by its id
func GetFollow(w http.ResponseWriter, r *http.Request) {

    setCommonHeaders(w)

    // get the id from the request params
    params := mux.Vars(r)

    // convert the id type from string to uuid.UUID
    followID, err := uuid.Parse(params["id"])

    if err != nil {
        log.Fatalf("Unable to parse the UUID. %v", err)
    }

    // call the getFollow function with followID to retrieve a single follow
    follow, err := getFollow(followID)

    if err != nil {
        log.Fatalf("Unable to get follow. %v", err)
    }

    // send the response
    json.NewEncoder(w).Encode(follow)
}

// Deletes follow in the postgres db
func DeleteFollow(w http.ResponseWriter, r *http.Request) {

    setCommonHeaders(w)
    setAdditionalHeaders(w, "DELETE")

    // get the id from the request params, key is "id"
    params := mux.Vars(r)

    // convert the id in string to uuid.UUID
    followID, err := uuid.Parse(params["id"])

    if err != nil {
        log.Fatalf("Unable to parse the UUID.  %v", err)
    }

    // call deleteFollow
    deletedRows := deleteFollow(followID)

    // format the message string
    msg := fmt.Sprintf("Follow deleted successfully. Total rows/record affected %v", deletedRows)

    // format the reponse message
    res := response{
        ID:      followID.String(),
        Message: msg,
    }

    // send the response
    json.NewEncoder(w).Encode(res)
}

// ------------------------- handler functions ----------------
// insert one follow in the DB
func insertFollow(follow models.Follow) uuid.UUID {

    // create the insert sql query
    // will return the id of the inserted follow
    sqlStatement := `INSERT INTO follows (id, followed_id, follower_id, date) VALUES ($1, $2, $3, $4) RETURNING id`

    // generate a new UUID for the session
    followID := uuid.New()

    // execute the sql statement
    err := repositories.Pool.QueryRow(context.Background(), sqlStatement, followID, follow.FollowedId, follow.FollowedId, follow.Date).Scan(&followID)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    fmt.Printf("Inserted a single record %v", followID)

    // return the inserted id
    return followID
}

// get one follow from the DB by its followID
func getFollow(followID uuid.UUID) (models.Follow, error) {

    // create a follow of models.Follow type
    var follow models.Follow

    // create the select sql query
    sqlStatement := `SELECT * FROM follows WHERE id=$1`

    // execute the sql statement
    row := repositories.Pool.QueryRow(context.Background(), sqlStatement, followID)

    // unmarshal the row object to follow
    err := row.Scan(&follow.ID, &follow.FollowedId, &follow.FollowerId, &follow.Date)

    switch err {
    case sql.ErrNoRows:
        fmt.Println("No rows were returned!")
        return follow, nil
    case nil:
        return follow, nil
    default:
        log.Fatalf("Unable to scan the row. %v", err)
    }

    // return empty follow on error
    return follow, err
}

// delete follow in the DB
func deleteFollow(followID uuid.UUID) int64 {

    // create the delete sql query
    sqlStatement := `DELETE FROM follows WHERE id=$1`

    // execute the sql statement
    res, err := repositories.Pool.Exec(context.Background(), sqlStatement, followID)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    // check how many rows affected
    rowsAffected := res.RowsAffected()

    fmt.Printf("Total rows/record affected %v", rowsAffected)

    return rowsAffected
}