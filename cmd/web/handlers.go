package main

import (
	"groupie-tracker/cmd/checkRequest"
	"groupie-tracker/cmd/errors"
	"groupie-tracker/pkg"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		check := checkRequest.CheckStatus(http.StatusMethodNotAllowed)
		w.WriteHeader(http.StatusMethodNotAllowed)
		errors.CheckErrors(w, check)
		return
	}

	if r.URL.Path != "/" {
		check := checkRequest.CheckStatus(http.StatusNotFound)
		w.WriteHeader(http.StatusNotFound)
		errors.CheckErrors(w, check)
		return
	}

	ts, err := template.ParseFiles("./ui/html/home.html")
	if err != nil {
		check := checkRequest.CheckStatus(http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		errors.CheckErrors(w, check)
		return
	}

	artists := []Artist{}
	if err = pkg.ConvertJson(ArtistsURL, &artists); err != nil {
		check := checkRequest.CheckStatus(http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		errors.CheckErrors(w, check)
		return
	}

	err = ts.Execute(w, artists)
	if err != nil {
		check := checkRequest.CheckStatus(http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		errors.CheckErrors(w, check)
		return
	}
}

func (app *application) band(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("./ui/html/band.html")
	if err != nil {
		check := checkRequest.CheckStatus(http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		errors.CheckErrors(w, check)
		return
	}

	artists := []Artist{}
	if err = pkg.ConvertJson(ArtistsURL, &artists); err != nil {
		check := checkRequest.CheckStatus(http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		errors.CheckErrors(w, check)
		return
	}

	idText := strings.TrimPrefix(r.URL.Path, "/artist/")
	id, err := strconv.Atoi(idText)
	if err != nil || id < 1 || id > len(artists) {
		check := checkRequest.CheckStatus(http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		errors.CheckErrors(w, check)
		return
	}

	if artists[id-1].InfoConcert() != nil {
		check := checkRequest.CheckStatus(http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		errors.CheckErrors(w, check)
		return
	}

	err = tmp.Execute(w, artists[id-1])
	if err != nil {
		check := checkRequest.CheckStatus(http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		errors.CheckErrors(w, check)
		return
	}
}
