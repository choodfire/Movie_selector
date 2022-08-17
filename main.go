package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"go_gui/data"
	"image/color"
	"strconv"

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

	list.OnSelected = func(id widget.ListItemID) {
		img := canvas.NewImageFromFile(fmt.Sprintf("./temp/%d.jpg", movies.Results[id].FilmId))
		img.FillMode = canvas.ImageFillContain
		img.Resize(fyne.Size{300, 400})
		img.Move(fyne.Position{50, 10})

		text := canvas.NewText(movies.Results[id].Title, color.White)
		text.Resize(fyne.Size{400, 130})
		text.Move(fyne.Position{0, 370})
		text.Alignment = fyne.TextAlignCenter

		score := canvas.NewText(movies.Results[id].Score, color.White)
		score.Resize(fyne.Size{400, 130})
		score.Move(fyne.Position{0, 390})
		score.Alignment = fyne.TextAlignCenter
		scoreInt, _ := strconv.ParseFloat(movies.Results[id].Score, 64)
		if scoreInt < 5.0 {
			score.Color = color.RGBA{255, 0, 0, 255}
		} else if scoreInt > 7.0 {
			score.Color = color.RGBA{0, 255, 0, 255}
		} else {
			score.Color = color.RGBA{128, 128, 128, 255}
		}

		cntnr := container.NewWithoutLayout(img, text, score)

		w.SetContent(container.NewHSplit(list, cntnr))
	}

	w.SetContent(container.NewHSplit(list, container.NewMax(pageText)))

	w.ShowAndRun()
}

// при ресайзе изменять длину и локацию картинки и текста
