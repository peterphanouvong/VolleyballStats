package main

import (
	"fmt"
	"log"
	"net/http"

	"vballStats/router"
)


func main () {
	r := router.Router()

	fmt.Println("Starting Server on Port 10000")
	log.Fatal(http.ListenAndServe(":10000", r))
}