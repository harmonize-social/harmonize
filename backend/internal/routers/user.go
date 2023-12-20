package routers

import (
    "backend/internal/handlers"

    "github.com/gorilla/mux"
)

func UserRouter() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/login", handlers.Login).Methods("POST")
    router.HandleFunc("/register", handlers.Register).Methods("POST")
    return router
}

func FollowRouter() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/", handlers.CreateFollow).Methods("POST")
    router.HandleFunc("/{id}", handlers.GetFollow).Methods("GET")
    router.HandleFunc("/{id}", handlers.DeleteFollow).Methods("DELETE")
    return router
}
