package routers

import (
    "backend/internal/handlers"

    "github.com/gorilla/mux"
)

func UserRouter() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/", handlers.CreateUser).Methods("POST")
    router.HandleFunc("/{id}", handlers.GetUser).Methods("GET")
    router.HandleFunc("/{id}", handlers.UpdateUser).Methods("PUT")
    router.HandleFunc("/{id}", handlers.DeleteUser).Methods("DELETE")
    return router
}