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

func CreateSavedPost(w http.ResponseWriter, r *http.Request) {
    // Set headers
    setCommonHeaders(w)
    setAdditionalHeaders(w, "POST")

    // create an empty savedPost of type models.SavedPost
    var savedPost models.SavedPost

    // decode the json request to savedPost
    err := json.NewDecoder(r.Body).Decode(&savedPost)

    if err != nil {
        log.Fatalf("Unable to decode the request body.  %v", err)
    }

    // call insertSavedPost function and pass the savedPost
    insertID := insertSavedPost(savedPost)

    // format a response object
    res := response{
        ID:      insertID.String(),
        Message: "Saved post created successfully",
    }

    // send the response
    json.NewEncoder(w).Encode(res)
}

// Will return a single savedPost by its id
func GetSavedPost(w http.ResponseWriter, r *http.Request) {

    setCommonHeaders(w)

    // get the id from the request params
    params := mux.Vars(r)

    // convert the id type from string to uuid.UUID
    savedPostID, err := uuid.Parse(params["id"])

    if err != nil {
        log.Fatalf("Unable to parse the UUID. %v", err)
    }

    // call the getSavedPost function with savedPostID to retrieve a single savedPost
    savedPost, err := getSavedPost(savedPostID)

    if err != nil {
        log.Fatalf("Unable to get saved post. %v", err)
    }

    // send the response
    json.NewEncoder(w).Encode(savedPost)
}

// Updates the savedPost's details in the postgres db
func UpdateSavedPost(w http.ResponseWriter, r *http.Request) {

    setCommonHeaders(w)
    setAdditionalHeaders(w, "PUT")

    // get the id from the request params, key is "id"
    params := mux.Vars(r)

    // convert the id type from string to uuid.UUID
    savedPostID, err := uuid.Parse(params["id"])

    if err != nil {
        log.Fatalf("Unable to parse the UUID. %v", err)
    }

    // create an empty savedPost of type models.SavedPost
    var savedPost models.SavedPost

    // decode the json request the savedPost
    err = json.NewDecoder(r.Body).Decode(&savedPost)

    if err != nil {
        log.Fatalf("Unable to decode the request body.  %v", err)
    }

    // call updatePost
    updatedRows := updateSavedPost(savedPostID, savedPost)

    // format the message string
    msg := fmt.Sprintf("Saved post updated successfully. Total rows/record affected %v", updatedRows)

    // format the response message
    res := response{
        ID:      savedPostID.String(),
        Message: msg,
    }

    // send the response
    json.NewEncoder(w).Encode(res)
}

// Deletes savedPost in the postgres db
func DeleteSavedPost(w http.ResponseWriter, r *http.Request) {

    setCommonHeaders(w)
    setAdditionalHeaders(w, "DELETE")

    // get the id from the request params, key is "id"
    params := mux.Vars(r)

    // convert the id in string to uuid.UUID
    savedPostID, err := uuid.Parse(params["id"])

    if err != nil {
        log.Fatalf("Unable to parse the UUID.  %v", err)
    }

    // call deleteSavedPost
    deletedRows := deleteSavedPost(savedPostID)

    // format the message string
    msg := fmt.Sprintf("Saved post deleted successfully. Total rows/record affected %v", deletedRows)

    // format the reponse message
    res := response{
        ID:      savedPostID.String(),
        Message: msg,
    }

    // send the response
    json.NewEncoder(w).Encode(res)
}

// ------------------------- handler functions ----------------
// insert one savedPost in the DB
func insertSavedPost(savedPost models.SavedPost) uuid.UUID {

    // create the insert sql query
    // will return the id of the inserted savedPost
    sqlStatement := `INSERT INTO saved_posts (id, user_id, post_id) VALUES ($1, $2, $3) RETURNING id`

    // generate a new UUID for the savedPost
    savedPostID := uuid.New()

    // execute the sql statement
    err := repositories.Pool.QueryRow(context.Background(), sqlStatement, savedPostID, savedPost.UserId, savedPost.PostId).Scan(&savedPostID)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    fmt.Printf("Inserted a single record %v", savedPostID)

    // return the inserted id
    return savedPostID
}

// get one savedPost from the DB by its postID
func getSavedPost(savedPostID uuid.UUID) (models.SavedPost, error) {

    // create a savedPost of models.SavedPost type
    var savedPost models.SavedPost

    // create the select sql query
    sqlStatement := `SELECT * FROM saved_posts WHERE id=$1`

    // execute the sql statement
    row := repositories.Pool.QueryRow(context.Background(), sqlStatement, savedPostID)

    // unmarshal the row object to savedPost
    err := row.Scan(&savedPost.ID, &savedPost.UserId, &savedPost.PostId)

    switch err {
    case sql.ErrNoRows:
        fmt.Println("No rows were returned!")
        return savedPost, nil
    case nil:
        return savedPost, nil
    default:
        log.Fatalf("Unable to scan the row. %v", err)
    }

    // return empty savedPost on error
    return savedPost, err
}

// update savedPost in the DB
func updateSavedPost(savedPostID uuid.UUID, savedPost models.SavedPost) int64 {

    // create the update sql query
    sqlStatement := `UPDATE saved_posts SET user_id=$2, post_id=$3 WHERE id=$1`

    // execute the sql statement
    res, err := repositories.Pool.Exec(context.Background(), sqlStatement, savedPostID, savedPost.UserId, savedPost.PostId)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    // check how many rows affected
    rowsAffected := res.RowsAffected() // how to see/respond if there is an eror with checking the rows?

    fmt.Printf("Total rows/record affected %v", rowsAffected)

    return rowsAffected
}

// delete savedPost in the DB
func deleteSavedPost(savedPostID uuid.UUID) int64 {

    // create the delete sql query
    sqlStatement := `DELETE FROM saved_posts WHERE id=$1`

    // execute the sql statement
    res, err := repositories.Pool.Exec(context.Background(), sqlStatement, savedPostID)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    // check how many rows affected
    rowsAffected := res.RowsAffected()

    fmt.Printf("Total rows/record affected %v", rowsAffected)

    return rowsAffected
}