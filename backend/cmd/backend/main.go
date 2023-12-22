package main

import (
    "backend/internal/repositories"
    "backend/internal/routers"
    "fmt"

    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/rs/cors"
)

func main() {
    err := repositories.CreateConnection()
    if err != nil {
        log.Fatal(err)
    }
    err = repositories.GenerateSecret()
    if err != nil {
        log.Fatal(err)
    }

    router := routers.FullRouter()

    c := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://172.20.0.4:5173"},
        AllowCredentials: true,
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    })

    PrintRouter(router)

    handler := c.Handler(router)

    log.Fatal(http.ListenAndServe(":8080", handler))
}

func PrintRouter(router *mux.Router) {
    router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
        t, err := route.GetPathTemplate()
        if err == nil {
            fmt.Println("ROUTE:", t)
        }
        return nil
    })
}
