package httputil

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var client = &http.Client{Timeout: 15 * time.Second}

func GetJSON(url string) (string, error) {
	res, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	return getPayload(res), nil
}

func GetJSONWithAuth(url string, token string) (string, error) {
	var bearer = "Bearer " + token
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", bearer)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	return getPayload(res), nil
}

func getPayload(response *http.Response) string {
	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	return string(body)
}
