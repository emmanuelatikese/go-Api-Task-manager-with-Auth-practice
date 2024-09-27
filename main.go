package main

import (
	"log"

	"api.task/controllers"
	apiDB "api.task/db"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)



func main() {
	err := apiDB.ConnectMongo()
	if err != nil {
		log.Fatal("not able to connect")
	}
	log.Print("db connected ....")
	mux := mux.NewRouter().StrictSlash(false)
	mux.HandleFunc("/signUp", controllers.SignUp).Methods("POST")
	mux.HandleFunc("/login", controllers.Login).Methods("POST")
	n := negroni.Classic()
	n.UseHandler(mux)
	log.Print("listening to server: 8080")
	n.Run(":8080")
}