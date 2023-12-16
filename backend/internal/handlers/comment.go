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

func CreateComment(w http.ResponseWriter, r *http.Request) {
    // Set headers
    setCommonHeaders(w)
    setAdditionalHeaders(w, "POST")

    // create an empty comment of type models.Comment
    var comment models.Comment

    // decode the json request to comment
    err := json.NewDecoder(r.Body).Decode(&comment)

    if err != nil {
        log.Fatalf("Unable to decode the request body.  %v", err)
    }

    // call insert user function and pass the comment
    insertID := insertComment(comment)

    // format a response object
    res := response{
        ID:      insertID.String(),
        Message: "Comment created successfully",
    }

    // send the response
    json.NewEncoder(w).Encode(res)
}

// GetComment will return a single comment by its id
func GetComment(w http.ResponseWriter, r *http.Request) {

    setCommonHeaders(w)

    // get the commentid from the request params
    params := mux.Vars(r)

    // convert the id type from string to uuid.UUID
    commentID, err := uuid.Parse(params["id"])

    if err != nil {
        log.Fatalf("Unable to parse the UUID. %v", err)
    }

    // call the getComment function with postID to retrieve a single post
    comment, err := getComment(commentID)

    if err != nil {
        log.Fatalf("Unable to get comment. %v", err)
    }

    // send the response
    json.NewEncoder(w).Encode(comment)
}

// Updates the comment's details in the postgres db
func UpdateComment(w http.ResponseWriter, r *http.Request) {

    setCommonHeaders(w)
    setAdditionalHeaders(w, "PUT")

    // get the postid from the request params, key is "id"
    params := mux.Vars(r)

    // convert the id type from string to uuid.UUID
    commentID, err := uuid.Parse(params["id"])

    if err != nil {
        log.Fatalf("Unable to parse the UUID. %v", err)
    }

    // create an empty comment of type models.Comment
    var comment models.Comment

    // decode the json request of the comment
    err = json.NewDecoder(r.Body).Decode(&comment)

    if err != nil {
        log.Fatalf("Unable to decode the request body.  %v", err)
    }

    // call updateComment
    updatedRows := updateComment(commentID, comment)

    // format the message string
    msg := fmt.Sprintf("Comment updated successfully. Total rows/record affected %v", updatedRows)

    // format the response message
    res := response{
        ID:      commentID.String(),
        Message: msg,
    }

    // send the response
    json.NewEncoder(w).Encode(res)
}

// Deletes comment in the postgres db
func DeleteComment(w http.ResponseWriter, r *http.Request) {

    setCommonHeaders(w)
    setAdditionalHeaders(w, "DELETE")

    // get the id from the request params, key is "id"
    params := mux.Vars(r)

    // convert the id in string to uuid.UUID
    commentID, err := uuid.Parse(params["id"])

    if err != nil {
        log.Fatalf("Unable to parse the UUID.  %v", err)
    }

    // call deleteComment
    deletedRows := deleteComment(commentID)

    // format the message string
    msg := fmt.Sprintf("Comment deleted successfully. Total rows/record affected %v", deletedRows)

    // format the reponse message
    res := response{
        ID:      commentID.String(),
        Message: msg,
    }

    // send the response
    json.NewEncoder(w).Encode(res)
}

// ------------------------- handler functions ----------------
// insert one comment in the DB
func insertComment(comment models.Comment) uuid.UUID {

    // create the insert sql query
    // will return the id of the inserted comment
    sqlStatement := `INSERT INTO comments (id, post_id, user_id, reply_to_id, message) VALUES ($1, $2, $3, $4, $5) RETURNING id`

    // generate a new UUID for the comment
    commentID := uuid.New()

    // execute the sql statement
    err := repositories.Pool.QueryRow(context.Background(), sqlStatement, commentID, comment.PostId, comment.UserId, comment.ReplyToId, comment.Message).Scan(&commentID)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    fmt.Printf("Inserted a single record %v", commentID)

    // return the inserted id
    return commentID
}

// get one comment from the DB by its commentID
func getComment(commentID uuid.UUID) (models.Comment, error) {

    // create a comment of models.Comment type
    var comment models.Comment

    // create the select sql query
    sqlStatement := `SELECT * FROM comments WHERE id=$1`

    // execute the sql statement
    row := repositories.Pool.QueryRow(context.Background(), sqlStatement, commentID)

    // unmarshal the row object to comment
    err := row.Scan(&comment.ID, &comment.PostId, &comment.UserId, &comment.ReplyToId, &comment.Message)

    switch err {
    case sql.ErrNoRows:
        fmt.Println("No rows were returned!")
        return comment, nil
    case nil:
        return comment, nil
    default:
        log.Fatalf("Unable to scan the row. %v", err)
    }

    // return empty comment on error
    return comment, err
}

// update comment in the DB
func updateComment(commentID uuid.UUID, comment models.Comment) int64 {

    // create the update sql query
    sqlStatement := `UPDATE comments SET post_id=$2, user_id=$3, reply_to_id=$4, message=$5 WHERE id=$1`

    // execute the sql statement
    res, err := repositories.Pool.Exec(context.Background(), sqlStatement, commentID, comment.PostId, comment.UserId, comment.ReplyToId, comment.Message)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    // check how many rows affected
    rowsAffected := res.RowsAffected() // how to see/respond if there is an eror with checking the rows?

    fmt.Printf("Total rows/record affected %v", rowsAffected)

    return rowsAffected
}

// delete comment in the DB
func deleteComment(commentID uuid.UUID) int64 {

    // create the delete sql query
    sqlStatement := `DELETE FROM comments WHERE id=$1`

    // execute the sql statement
    res, err := repositories.Pool.Exec(context.Background(), sqlStatement, commentID)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    // check how many rows affected
    rowsAffected := res.RowsAffected()

    fmt.Printf("Total rows/record affected %v", rowsAffected)

    return rowsAffected
}