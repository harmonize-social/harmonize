package routers

import (
    "backend/internal/handlers"
    "net/http"

    "github.com/gorilla/mux"
)

func UserRouter() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/login", handlers.Login).Methods(http.MethodPost, http.MethodOptions)
    router.HandleFunc("/register", handlers.Register).Methods(http.MethodPost, http.MethodOptions)
    return router
}

func FollowRouter() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/", handlers.CreateFollow).Methods("POST")
    router.HandleFunc("/{id}", handlers.GetFollow).Methods("GET")
    router.HandleFunc("/{id}", handlers.DeleteFollow).Methods("DELETE")
    return router
}
