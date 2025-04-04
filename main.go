package main

import (
	"log"
	"net/http"

	"github.com/samirkumarbarman/basic-crud-go/routes"
)

func main() {
	router := routes.SetupRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
