package handlers

import (
    "net/http" // used to access the request and response object of the api
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
    // // Set headers
    // setCommonHeaders(w)
    // setAdditionalHeaders(w, "POST")
    //
    // // create an empty post of type models.Post
    // var post models.Post
    //
    // // decode the json request to post
    // err := json.NewDecoder(r.Body).Decode(&post)
    //
    // if err != nil {
    //     log.Fatalf("Unable to decode the request body.  %v", err)
    // }
    //
    // // call insert user function and pass the post
    // insertID := insertPost(post)
    //
    // // format a response object
    // res := response{
    //     ID:      insertID.String(),
    //     Message: "Post created successfully",
    // }
    //
    // // send the response
    // json.NewEncoder(w).Encode(res)
}

// GetPost will return a single post by its id
func GetPost(w http.ResponseWriter, r *http.Request) {
    //
    // setCommonHeaders(w)
    //
    // // get the postid from the request params
    // params := mux.Vars(r)
    //
    // // convert the id type from string to uuid.UUID
    // postID, err := uuid.Parse(params["id"])
    //
    // if err != nil {
    //     log.Fatalf("Unable to parse the UUID. %v", err)
    // }
    //
    // // call the getPost function with postID to retrieve a single post
    // post, err := getPost(postID)
    //
    // if err != nil {
    //     log.Fatalf("Unable to get post. %v", err)
    // }
    //
    // // send the response
    // json.NewEncoder(w).Encode(post)
}

// Updates the post's details in the postgres db
func UpdatePost(w http.ResponseWriter, r *http.Request) {
    //
    // setCommonHeaders(w)
    // setAdditionalHeaders(w, "PUT")
    //
    // // get the postid from the request params, key is "id"
    // params := mux.Vars(r)
    //
    // // convert the id type from string to uuid.UUID
    // postID, err := uuid.Parse(params["id"])
    //
    // if err != nil {
    //     log.Fatalf("Unable to parse the UUID. %v", err)
    // }
    //
    // // create an empty post of type models.Post
    // var post models.Post
    //
    // // decode the json request the post
    // err = json.NewDecoder(r.Body).Decode(&post)
    //
    // if err != nil {
    //     log.Fatalf("Unable to decode the request body.  %v", err)
    // }
    //
    // // call updatePost
    // updatedRows := updatePost(postID, post)
    //
    // // format the message string
    // msg := fmt.Sprintf("Post updated successfully. Total rows/record affected %v", updatedRows)
    //
    // // format the response message
    // res := response{
    //     ID:      postID.String(),
    //     Message: msg,
    // }
    //
    // // send the response
    // json.NewEncoder(w).Encode(res)
}

// Deletes post in the postgres db
func DeletePost(w http.ResponseWriter, r *http.Request) {
    //
    // setCommonHeaders(w)
    // setAdditionalHeaders(w, "DELETE")
    //
    // // get the postid from the request params, key is "id"
    // params := mux.Vars(r)
    //
    // // convert the id in string to uuid.UUID
    // postID, err := uuid.Parse(params["id"])
    //
    // if err != nil {
    //     log.Fatalf("Unable to parse the UUID.  %v", err)
    // }
    //
    // // call deletePost
    // deletedRows := deletePost(postID)
    //
    // // format the message string
    // msg := fmt.Sprintf("Post deleted successfully. Total rows/record affected %v", deletedRows)
    //
    // // format the reponse message
    // res := response{
    //     ID:      postID.String(),
    //     Message: msg,
    // }
    //
    // // send the response
    // json.NewEncoder(w).Encode(res)
}
