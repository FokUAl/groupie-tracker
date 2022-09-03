package errors

import (
	"html/template"
	"net/http"
)

func CheckErrors(w http.ResponseWriter, Msg string) {
	tmpl, err := template.ParseFiles("./ui/html/error.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, Msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
