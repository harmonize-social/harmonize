package routers

import (
	"github.com/gorilla/mux"
	"backend/internal/handlers"
)

// Router is used in main.go
func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/user/{id}", handlers.GetUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/user", handlers.GetAllUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/newuser", handlers.CreateUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/user/{id}", handlers.UpdateUser).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deleteuser/{id}", handlers.DeleteUser).Methods("DELETE", "OPTIONS")

	return router
}