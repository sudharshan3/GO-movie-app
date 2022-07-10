package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sudharshan3/GO-movie-app/router"
)

func main() {
	fmt.Println("MongoDB CRUD")
	r := router.Router()
	fmt.Println("Starting server on port 4000")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening on port 4000")

}
