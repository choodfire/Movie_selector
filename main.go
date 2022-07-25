package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"go_gui/data"
	"log"
)

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

	//list.OnSelected = func(id widget.ListItemID) {
	//	widget.NewLabel()
	//}

	welcomeText := widget.NewLabel("Select movie")
	welcomeText.Alignment = fyne.TextAlignCenter

	w.SetContent(container.NewHSplit(
		list,
		container.NewMax(welcomeText),
	))

	w.ShowAndRun()
}
