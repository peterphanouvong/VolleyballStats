package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var tpl *template.Template

var counter int = 0

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", counter)

	fmt.Println("Endpoint Hit: homePage")
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	counter++
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	resp := &struct{
		Message string
		Data int
	}{
		Message: "OK",
		Data: counter,
	}

	json.NewEncoder(w).Encode(resp)
}

func handleRequests() {
	r := mux.NewRouter()
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	r.HandleFunc("/counter/increment", incrementCounter).Methods("GET")
	r.HandleFunc("/", homePage).Methods("GET")
	log.Fatal(http.ListenAndServe(":10000", r))
}

func main () {
	handleRequests()
}