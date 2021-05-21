package middleware

import (
	"fmt"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", counter)

	fmt.Println("Endpoint Hit: homePage")
}

func PlayerProfilePage(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "player-profile.page.gohtml", nil)

	fmt.Println("Endpoint Hit: Player Profile Page")
}