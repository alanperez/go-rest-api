package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alanperez/go-rest-api/routes"
)

func main() {
	port := "3000"
	fmt.Printf("App is running on port %s\n", port)
	r := routes.NewRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
