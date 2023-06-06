package main

import (
	"log"
	"net/http"
	"parking/pkg/routes"
)

func main() {
	mux := routes.Routes();

	
	log.Fatal(http.ListenAndServe(":8181",mux))

}