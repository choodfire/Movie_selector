package data

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

type MovieResults struct {
	Results []Movie `json:"films"`
}

func (current *MovieResults) addItems(other MovieResults) {
	current.Results = append(current.Results, other.Results...)
}

type Movie struct {
	Title      string `json:"nameRu"`
	Score      string `json:"rating"`
	FilmId     int    `json:"filmId"`
	PosterPath string `json:"posterUrl"`
	//Description string  `json:"overview"`
}

func Update() (MovieResults, error) {
	var movies MovieResults
	for page := 1; page < 6; page++ {
		link := fmt.Sprintf("https://kinopoiskapiunofficial.tech/api/v2.2/films/top?type=TOP_100_POPULAR_FILMS&page=%d", page)
		req, err := http.NewRequest("GET", link, nil)
		if err != nil {
			return MovieResults{}, err
		}

		req.Header.Set("Accept", "application/json")
		req.Header.Set("X-Api-Key", api)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return MovieResults{}, err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return MovieResults{}, err
		}

		var temp MovieResults
		err = json.Unmarshal([]byte(string(body)), &temp)
		if err != nil {
			return MovieResults{}, err
		}

		movies.addItems(temp)
	}

	return movies, nil
}

func SavePosters(movies MovieResults) error {
	if _, err := os.Stat("temp"); os.IsNotExist(err) {
		err := os.Mkdir("temp", 0777)
		if err != nil {
			return err
		}
	}

	wg := sync.WaitGroup{}
	ch := make(chan error)

	for _, movie := range movies.Results {
		wg.Add(1)
		movie := movie
		go func() {
			filePath := fmt.Sprintf("temp/%d.jpg", movie.FilmId)
			if _, err := os.Stat(filePath); !os.IsNotExist(err) {
				wg.Done()
				return
			}

			img, err := os.Create(filePath)
			if err != nil {
				//return err
				ch <- err
			}
			defer img.Close()

			resp, err := http.Get(movie.PosterPath)
			if err != nil {
				//return err
				ch <- err
			}
			defer resp.Body.Close()

			_, err = io.Copy(img, resp.Body)
			if err != nil {
				//return err
				ch <- err
			}
			wg.Done()
		}()
	}
	wg.Wait()

	select {
	case err := <-ch:
		return err
	default:
		return nil
	}
}
