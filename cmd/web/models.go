package main

import (
	"encoding/json"
	"groupie-tracker/pkg"
)

const (
	ArtistsURL = "https://groupietrackers.herokuapp.com/api/artists"
)

type Artist struct {
	Id           int                 `json:"id"`
	Image        string              `json:"image"`
	Name         string              `json:"name"`
	Members      []string            `json:"members"`
	CreationDate int                 `json:"creationDate"`
	ConcertDates string              `json:"concertDates"`
	FirstAlbum   string              `json:"firstAlbum"`
	Relation     string              `json:"relations"`
	Concerts     map[string][]string `json:"datesLocations"`
}

type Relation struct {
	Concerts map[string][]string `json:"datesLocations"`
}

func (artist *Artist) InfoConcert() error {
	r := Relation{}
	body, err := pkg.GetJson(artist.Relation)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, &r); err != nil {
		return err
	}

	artist.Concerts = r.Concerts
	return nil
}
