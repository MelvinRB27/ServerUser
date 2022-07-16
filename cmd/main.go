package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MelvinRB27/server-user/authorization"
	"github.com/MelvinRB27/server-user/handlers"
	"github.com/MelvinRB27/server-user/storage"
)

func main() {
	err := authorization.LoadFiles("./certificates/app.rsa", "./certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("Error loading certificates: %v", err) 
	}

	mux := http.NewServeMux()
	handlers.Route(mux)

	//driver := storage.MySQL
	storage.NewMySQL()
	storage.Migrate()

	//server for handler
	fmt.Print("Starting server on port \n", "http://localhost:8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Printf("Error starting server on port %v\n", err)
	}

}
