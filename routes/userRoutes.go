package routes

import (
	"github.com/gorilla/mux"
	"github.com/samirkumarbarman/basic-crud-go/controllers"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")

	return router
}
