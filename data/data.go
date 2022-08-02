package data

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type MovieResults struct {
	Results []Movie `json:"films"`
}

type Movie struct {
	Title string `json:"nameRu"`
	Score string `json:"rating"`
	//Description string  `json:"overview"`
	PosterPath string `json:"posterUrl"`
}

func LoadData() (MovieResults, error) {
	data, err := ioutil.ReadFile("./data.json")
	if err != nil {
		return MovieResults{}, err
	}
	if json.Valid(data) == false {
		log.Fatal("JSON file isn't valid")
	}
	var Movies MovieResults
	err = json.Unmarshal(data, &Movies)
	if err != nil {
		return MovieResults{}, err
	}

	return Movies, nil
}
