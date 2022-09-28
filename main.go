package main

import (
	"Movie_selector/data"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"strconv"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	movies, err := data.UpdateMovieList()
	check(err)

	err = data.GetDescriptions(&movies)
	check(err)

	if len(movies.Results) == 100 {
		err = data.SaveToJSON(movies)
		check(err)
	}

	if len(movies.Results) == 0 {
		movies, err = data.GetFromJSON()
		if err != nil {
			a := app.New()
			w := a.NewWindow("Error")
			w.Resize(fyne.NewSize(200, 100))
			w.SetContent(widget.NewLabel("Количество запросов превышено. Пожалуйста, попробуйте позднее!"))
			w.ShowAndRun()
		}
		return
	}

	err = data.SavePosters(movies)
	check(err)

	a := app.New()
	w := a.NewWindow("Топ 100 самых ожидаемых фильмов")
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

	pageText := widget.NewLabel("Выберите фильм")
	pageText.Alignment = fyne.TextAlignCenter
	pageText.Wrapping = fyne.TextWrapWord

	list.OnSelected = func(id widget.ListItemID) {
		img := canvas.NewImageFromFile(fmt.Sprintf("./temp/%d.jpg", movies.Results[id].FilmId))
		img.SetMinSize(fyne.NewSize(300, 400))
		img.FillMode = canvas.ImageFillContain

		textTitle := canvas.NewText(movies.Results[id].Title, color.White)
		textTitle.Resize(fyne.Size{400, 130})
		textTitle.Alignment = fyne.TextAlignCenter

		textScore := canvas.NewText(movies.Results[id].Score, color.White)
		textScore.Resize(fyne.Size{400, 130})
		textScore.Alignment = fyne.TextAlignCenter

		scoreInt, _ := strconv.ParseFloat(movies.Results[id].Score, 64)
		if scoreInt < 5.0 {
			textScore.Color = color.RGBA{255, 0, 0, 255}
		} else if scoreInt > 7.0 {
			textScore.Color = color.RGBA{0, 255, 0, 255}
		} else {
			textScore.Color = color.RGBA{128, 128, 128, 255}
		}

		textDescription := pageText
		textDescription.SetText(movies.Results[id].Description)
		textDescription.Resize(fyne.Size{400, 130})
		textDescription.Wrapping = fyne.TextWrapWord

		cntnr := container.NewVScroll(container.NewVBox(img, textTitle, textScore, textDescription))

		w.SetContent(container.NewHSplit(
			list,
			cntnr,
		))
	}

	w.SetContent(container.NewHSplit(list, container.NewMax(pageText)))

	w.ShowAndRun()
}
