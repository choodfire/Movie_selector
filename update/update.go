package update

import (
	"io/ioutil"
	"net/http"
	"os"
)

func Update() error {
	req, err := http.NewRequest("GET", "https://kinopoiskapiunofficial.tech/api/v2.2/films/top?type=TOP_100_POPULAR_FILMS&page=1", nil)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Api-Key", api)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	//fmt.Println(string(body))

	file, err := os.OpenFile("data.json", os.O_WRONLY|os.O_CREATE, 0444)
	if err != nil {
		return err
	}

	_, err = file.Write([]byte(string(body)))
	if err != nil {
		return err
	}

	return nil
}
