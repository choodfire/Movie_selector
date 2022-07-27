package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"go_gui/data"
	"io"
	"log"
	"net/http"
	"os"
)

func savePoster(link string) error {
	err := os.Mkdir("temp", 0750)
	if err != nil {
		return err
	}
	fileURL := fmt.Sprintf("https://image.tmdb.org/t/p/w500%s", link)
	filePath := fmt.Sprintf("temp%s", link)
	img, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer img.Close()

	resp, err := http.Get(fileURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(img, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	movies, err := data.LoadData()
	if err != nil {
		log.Fatal(err)
	}

	a := app.New()
	w := a.NewWindow("List of most popular movies")
	w.Resize(fyne.NewSize(600, 400))

	list := widget.NewList(
		func() int {
			return len(movies.Results)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(id widget.ListItemID, object fyne.CanvasObject) {
			object.(*widget.Label).SetText(movies.Results[id].Title)
		},
	)

	pageText := widget.NewLabel("Select movie")
	pageText.Alignment = fyne.TextAlignCenter
	pageText.Wrapping = fyne.TextWrapWord

	list.OnSelected = func(id widget.ListItemID) {
		pageText.SetText(movies.Results[id].Description)
	}

	w.SetContent(container.NewHSplit(
		list,
		container.NewMax(pageText),
	))

	w.ShowAndRun()
}
