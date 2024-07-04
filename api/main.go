package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Loading config...")
	config.Load()

	fmt.Println("Setting up routes...")
	r := router.Generate()

	fmt.Printf("Listening on port %d!\n", config.ApiPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.ApiPort), r))
}
