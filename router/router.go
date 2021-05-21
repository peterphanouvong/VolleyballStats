package router

import (
	"net/http"
	"vballStats/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	router.HandleFunc("/counter/increment", middleware.IncrementCounter).Methods("GET")
	router.HandleFunc("/", middleware.HomePage)
	router.HandleFunc("/player", middleware.PlayerProfilePage)
	return router
}