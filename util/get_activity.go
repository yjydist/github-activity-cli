package util

import (
	"encoding/json"
	"github-activity-cli/model"
	"io"
	"log"
	"net/http"
)

func GetActivity(username string) []model.Activity {
	url := "https://api.github.com/users/"
	url += username
	url += "/events"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(response.Body)
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	
	var Activities []model.Activity

	err = json.Unmarshal(responseData, &Activities)
	if err != nil {
		return nil
	}

	return Activities
}
