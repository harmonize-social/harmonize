package main

import (
	"fmt"
	"log"
	"net/http"
	"backend/internal/routers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// User routes
	path := "/api/user"
	routers.RouterUser(path, router)

	/* Just an example right now
	// Album routes
	path := "/api/album"
	routers.RouterAlbum(path, router)
	*/

	fmt.Println("Starting server on port 8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}