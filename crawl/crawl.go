package crawl

import (
	"io/ioutil"
	"net/http"
)

func New() (string, error) {
	res, httpErr := http.Get("https://google.com")
	if httpErr != nil {
		return "", httpErr
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
