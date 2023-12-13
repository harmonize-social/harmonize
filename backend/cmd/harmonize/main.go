package main

import (
	"fmt"
	"log"
	"net/http"
	"backend/internal/routers"
)

func main() {
	r := routers.Router()

	fmt.Println("Starting server on port 8080")

	log.Fatal(http.ListenAndServe(":8080", r))
}