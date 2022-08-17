package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"go_gui/data"
	"image/color"

	//"image/color"
	"log"

	"fyne.io/fyne/v2/canvas"
)

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

	list.OnSelected = func(id widget.ListItemID) {
		img := canvas.NewImageFromFile(fmt.Sprintf("./temp/%d.jpg", movies.Results[id].FilmId))
		img.FillMode = canvas.ImageFillContain
		img.Resize(fyne.Size{300, 300})
		img.Move(fyne.Position{50, 10})

		text := canvas.NewText(movies.Results[id].Title, color.White)
		text.Resize(fyne.Size{400, 130})
		text.Move(fyne.Position{0, 270})
		text.Alignment = fyne.TextAlignCenter

		score := canvas.NewText(movies.Results[id].Score, color.White)
		score.Resize(fyne.Size{400, 130})
		score.Move(fyne.Position{0, 290})
		score.Alignment = fyne.TextAlignCenter

		cntnr := container.NewWithoutLayout(img, text, score)
		//cntnr := container.NewGridWithRows(2, img, text)

		w.SetContent(container.NewHSplit(list, cntnr))
	}

	w.SetContent(container.NewHSplit(list, container.NewMax(pageText)))

	w.ShowAndRun()
}

// 		image := canvas.NewImageFromFile(fmt.Sprintf("./temp/%d.jpg", movies.Results[id].FilmId))
//		image.FillMode = canvas.ImageFillOriginal
//		image.Resize(fyne.NewSize(200, 200))
//		content := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), image, layout.NewSpacer())
