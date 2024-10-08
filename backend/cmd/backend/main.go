package main

import (
    "backend/internal/repositories"
    "backend/internal/routers"

    "log"
    "net/http"

    "github.com/joho/godotenv"
    "github.com/rs/cors"
)

/*
Main entry point for the backend server.
*/
func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    err = repositories.CreateConnection()
    if err != nil {
        log.Fatal(err)
    }
    repositories.LoadEnv()

    router := routers.FullRouter()

    c := cors.New(cors.Options{
        AllowedOrigins:   []string{repositories.FrontendUrl},
        AllowCredentials: true,
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    })

    handler := c.Handler(router)

    log.Fatal(http.ListenAndServe(":8080", handler))
}
