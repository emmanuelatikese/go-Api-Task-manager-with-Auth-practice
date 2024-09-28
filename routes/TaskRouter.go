package routes

import (
	"api.task/controllers"
	"api.task/middlewares"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func TaskRouter(router *mux.Router) {
	taskRouter := mux.NewRouter().StrictSlash(false)
	taskRouter.HandleFunc("/task/createTask", controllers.CreateTask).Methods("POST")
	taskRouter.HandleFunc("/task/find/{id}", controllers.FindTask).Methods("GET")
	taskRouter.HandleFunc("/task/findAll", controllers.FindAll).Methods("GET")
	taskRouter.HandleFunc("/task/update/{id}", controllers.UpdateTask).Methods("PUT")
	taskRouter.HandleFunc("/task/deleteAll/", controllers.DeleteAll).Methods("DELETE")
	taskRouter.HandleFunc("/task/deleteOne/{id}", controllers.DeleteOne).Methods("DELETE")
	router.PathPrefix("/task").Handler(negroni.New(
		negroni.HandlerFunc(middlewares.ProtectRoutes),
		negroni.Wrap(taskRouter)))
		//created a new router and used the main router with a pathprefix to indicate functions that requires middlewares
		
	}