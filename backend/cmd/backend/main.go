package main

import (
    "backend/internal/repositories"
    "backend/internal/routers"

    "log"
    "net/http"

    "github.com/rs/cors"
)

func main() {
    err := repositories.CreateConnection()
    if err != nil {
        log.Fatal(err)
    }
    repositories.LoadEnv()

    router := routers.FullRouter()

    c := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://172.20.0.4:5173"},
        AllowCredentials: true,
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    })

    handler := c.Handler(router)

    log.Fatal(http.ListenAndServe(":8080", handler))
}
