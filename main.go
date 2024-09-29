package main

import (
	"log"
	apiDB "api.task/db"
	"api.task/routes"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)
//I am told negroni is quite old now but this library is quite amazing.


func main() {
	err := apiDB.ConnectMongo()
	if err != nil {
		log.Fatal("not able to connect")
	}
	log.Print("db connected ....")
	mux := mux.NewRouter().StrictSlash(false)
	routes.AuthRouter(mux)
	routes.TaskRouter(mux)
	n := negroni.Classic()
	n.UseHandler(mux)
	log.Print("listening to server: 8080")
	n.Run(":8080")
}