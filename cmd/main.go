package main

import (
	"bff-gatway/internal/helpers"
	"bff-gatway/internal/routes"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	logger := helpers.NewLogger(log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))

	router := routes.NewRouter(logger)

	port := ":8080"
	fmt.Printf("Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
