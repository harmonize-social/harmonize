package handlers

import (
    "backend/cmd/backend"
	"backend/internal/models" // models package where User schema is defined
	//"backend/internal/repositories"
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

func CreatePost(w http.ResponseWriter, r *http.Request) {
    // Set headers
    setCommonHeaders(w)
    setAdditionalHeaders(w, "POST")

    // create an empty post of type models.Post
    var post models.Post

    // decode the json request to post
    err := json.NewDecoder(r.Body).Decode(&post)

    if err != nil {
        log.Fatalf("Unable to decode the request body.  %v", err)
    }

    // call insert user function and pass the post
    insertID := insertPost(post)

    // format a response object
    res := response{
        ID:      insertID.String(),
        Message: "Post created successfully",
    }

    // send the response
    json.NewEncoder(w).Encode(res)
}

// GetPost will return a single post by its id
func GetPost(w http.ResponseWriter, r *http.Request) {

    setCommonHeaders(w)

    // get the postid from the request params
    params := mux.Vars(r)

    // convert the id type from string to uuid.UUID
    postID, err := uuid.Parse(params["id"])

    if err != nil {
        log.Fatalf("Unable to parse the UUID. %v", err)
    }

    // call the getPost function with postID to retrieve a single post
    post, err := getPost(postID)

    if err != nil {
        log.Fatalf("Unable to get post. %v", err)
    }

    // send the response
    json.NewEncoder(w).Encode(post)
}

// Updates the post's details in the postgres db
func UpdatePost(w http.ResponseWriter, r *http.Request) {

    setCommonHeaders(w)
    setAdditionalHeaders(w, "PUT")

    // get the postid from the request params, key is "id"
    params := mux.Vars(r)

    // convert the id type from string to uuid.UUID
    postID, err := uuid.Parse(params["id"])

    if err != nil {
        log.Fatalf("Unable to parse the UUID. %v", err)
    }

    // create an empty post of type models.Post
    var post models.Post

    // decode the json request the post
    err = json.NewDecoder(r.Body).Decode(&post)

    if err != nil {
        log.Fatalf("Unable to decode the request body.  %v", err)
    }

    // call updatePost
    updatedRows := updatePost(postID, post)

    // format the message string
    msg := fmt.Sprintf("Post updated successfully. Total rows/record affected %v", updatedRows)

    // format the response message
    res := response{
        ID:      postID.String(),
        Message: msg,
    }

    // send the response
    json.NewEncoder(w).Encode(res)
}

// Deletes post in the postgres db
func DeletePost(w http.ResponseWriter, r *http.Request) {

    setCommonHeaders(w)
    setAdditionalHeaders(w, "DELETE")

    // get the postid from the request params, key is "id"
    params := mux.Vars(r)

    // convert the id in string to uuid.UUID
    postID, err := uuid.Parse(params["id"])

    if err != nil {
        log.Fatalf("Unable to parse the UUID.  %v", err)
    }

    // call deletePost
    deletedRows := deletePost(postID)

    // format the message string
    msg := fmt.Sprintf("Post deleted successfully. Total rows/record affected %v", deletedRows)

    // format the reponse message
    res := response{
        ID:      postID.String(),
        Message: msg,
    }

    // send the response
    json.NewEncoder(w).Encode(res)
}

// ------------------------- handler functions ----------------
// insert one user in the DB
func insertPost(post models.Post) uuid.UUID {

    db := main.Pool

    // close the db connection
    defer db.Close()

    // create the insert sql query
    // will return the id of the inserted post
    sqlStatement := `INSERT INTO posts (id, user_id, caption, type, type_specific_id) VALUES ($1, $2, $3, $4, $5) RETURNING id`

    // generate a new UUID for the post
    postID := uuid.New()

    // execute the sql statement
    err := db.QueryRow(context.Background(), sqlStatement, postID, post.UserId, post.Caption, post.Type, post.TypeSpecificId).Scan(&postID)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    fmt.Printf("Inserted a single record %v", postID)

    // return the inserted id
    return postID
}

// get one post from the DB by its postID
func getPost(postID uuid.UUID) (models.Post, error) {
    
    db := main.Pool

    // close the db connection
    defer db.Close()

    // create a post of models.Post type
    var post models.Post

    // create the select sql query
    sqlStatement := `SELECT * FROM posts WHERE id=$1`

    // execute the sql statement
    row := db.QueryRow(context.Background(), sqlStatement, postID)

    // unmarshal the row object to post
    err := row.Scan(&post.ID, &post.UserId, &post.Caption, &post.Type, &post.TypeSpecificId)

    switch err {
    case sql.ErrNoRows:
        fmt.Println("No rows were returned!")
        return post, nil
    case nil:
        return post, nil
    default:
        log.Fatalf("Unable to scan the row. %v", err)
    }

    // return empty post on error
    return post, err
}

// update post in the DB
func updatePost(postID uuid.UUID, post models.Post) int64 {

    db := main.Pool

    // close the db connection
    defer db.Close()

    // create the update sql query
    sqlStatement := `UPDATE posts SET user_id=$2, caption=$3, type=$4, type=$5 WHERE userid=$1`

    // execute the sql statement
    res, err := db.Exec(context.Background(), sqlStatement, postID, post.UserId, post.Caption, post.Type, post.TypeSpecificId)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    // check how many rows affected
    rowsAffected := res.RowsAffected() // how to see/respond if there is an eror with checking the rows?

    fmt.Printf("Total rows/record affected %v", rowsAffected)

    return rowsAffected
}

// delete post in the DB
func deletePost(postID uuid.UUID) int64 {

    db := main.Pool

    // close the db connection
    defer db.Close()

    // create the delete sql query
    sqlStatement := `DELETE FROM posts WHERE id=$1`

    // execute the sql statement
    res, err := db.Exec(context.Background(), sqlStatement, postID)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    // check how many rows affected
    rowsAffected := res.RowsAffected()

    fmt.Printf("Total rows/record affected %v", rowsAffected)

    return rowsAffected
}
