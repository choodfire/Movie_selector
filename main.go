package main

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"io/ioutil"
	"log"
)

type MovieResults struct {
	Results []Movie `json:"results"`
}

type Movie struct {
	Title string  `json:"title"`
	Score float64 `json:"vote_average"`
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

type Page struct {
	Page int `json:"page"`
}

func LoadPage() (Page, error) {
	var page Page
	data, err := ioutil.ReadFile("./data.json")
	if err != nil {
		return Page{}, err
	}
	if json.Valid(data) == false {
		log.Fatal("JSON file isn't valid")
	}

	err = json.Unmarshal(data, &page)
	if err != nil {
		return Page{}, err
	}

	return page, nil
}

func main() {
	movies, err := LoadData()
	if err != nil {
		log.Fatal(err)
	}

	for _, movie := range movies.Results {
		fmt.Printf("%+v\n", movie)
	}

	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	))

	w.ShowAndRun()
}
