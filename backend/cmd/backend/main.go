package main

import (
    "backend/internal/repositories"
    "backend/internal/routers"

    "log"
    "net/http"
)

func main() {
    repositories.CreateConnection()
    repositories.GenerateSecret()

    router := routers.FullRouter()

    log.Fatal(http.ListenAndServe(":8080", router))
}
