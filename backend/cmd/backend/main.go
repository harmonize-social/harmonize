package main

import (
    "fmt"
    "log"
    "net/http"

    "backend/internal/routes"
)

func main() {
    // Get the mux router from the routes package
    router := routes.NewRouter()

    // Start the server
    fmt.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
