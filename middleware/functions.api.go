package middleware

import (
	"encoding/json"
	"html/template"
	"net/http"
)
var tpl *template.Template
var counter int = 0

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func IncrementCounter(w http.ResponseWriter, r *http.Request) {
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