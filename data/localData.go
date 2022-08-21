package data

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func SaveToJSON(movies MovieResults) error {
	if _, err := os.Stat("./data/data.json"); !os.IsNotExist(err) {
		err := os.Remove("./data/data.json")
		if err != nil {
			return err
		}
	}

	file, err := json.MarshalIndent(movies, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("./data/data.json", file, 0644)
	if err != nil {
		return err
	}

	return nil
}

func GetFromJSON() (MovieResults, error) {
	data, err := ioutil.ReadFile("./data/data.json")
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
