package main

import (
	"log"
	apiDB "api.task/db"
	"api.task/routes"
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
	routes.AuthRouter(mux)
	routes.TaskRouter(mux)
	n := negroni.Classic()
	n.UseHandler(mux)
	log.Print("listening to server: 8080")
	n.Run(":8080")
}