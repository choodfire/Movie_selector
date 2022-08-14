package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"go_gui/data"
	"log"
)

func updateWindow(w fyne.Window, list fyne.CanvasObject, cntnr fyne.CanvasObject) {
	w.SetContent(container.NewHSplit(list, cntnr))
}

func main() {
	movies, err := data.Update()
	if err != nil {
		log.Fatal(err)
	}
	err = data.SavePosters(movies)
	if err != nil {
		log.Fatal(err)
	}

	a := app.New()
	w := a.NewWindow("List of most popular movies")
	w.Resize(fyne.NewSize(800, 600))

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

	//image := canvas.NewImageFromFile(fmt.Sprintf("./temp/%d.jpg", movies.Results[11].FilmId))
	//cntnr := container.NewMax(image, pageText)

	cntnr := container.NewMax(pageText)

	list.OnSelected = func(id widget.ListItemID) {
		image := canvas.NewImageFromFile(fmt.Sprintf("./temp/%d.jpg", movies.Results[id].FilmId))
		image.FillMode = canvas.ImageFillOriginal

		pageText.SetText(fmt.Sprintf("%s\n%s", movies.Results[id].Title, movies.Results[id].Score))

		cntnr := container.NewMax(image, pageText)
		cntnr.Refresh()
		fmt.Println("cntnr changed")

		updateWindow(w, list, cntnr)
	}

	updateWindow(w, list, cntnr)

	w.ShowAndRun()
}
