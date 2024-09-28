package routes

import (
	"api.task/controllers"
	"github.com/gorilla/mux"
)

func TaskRouter(mux *mux.Router) {
	mux.HandleFunc("/createTask", controllers.CreateTask).Methods("POST")
	mux.HandleFunc("/find/{id}", controllers.FindTask).Methods("GET")
	mux.HandleFunc("/findAll", controllers.FindAll).Methods("GET")
	mux.HandleFunc("/update/{id}", controllers.UpdateTask).Methods("PUT")
	mux.HandleFunc("/deleteAll", controllers.DeleteAll).Methods("DELETE")
	mux.HandleFunc("/deleteOne/{id}", controllers.DeleteOne).Methods("DELETE")
}