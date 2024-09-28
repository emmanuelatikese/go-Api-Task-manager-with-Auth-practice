package routes

import (
	"api.task/controllers"
	"github.com/gorilla/mux"
)

func AuthRouter(mux *mux.Router){
	mux.HandleFunc("/signUp", controllers.SignUp).Methods("POST")
	mux.HandleFunc("/login", controllers.Login).Methods("POST")
	mux.HandleFunc("/logout", controllers.Logout).Methods("POST")
}