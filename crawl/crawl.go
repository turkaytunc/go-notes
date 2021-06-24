package crawl

import (
	"io/ioutil"
	"log"
	"net/http"
)

func SiteBodyLength(url string, c chan string) {
	res, httpErr := http.Get(url)
	if httpErr != nil {
		log.Fatal(httpErr)
	}
	defer res.Body.Close()

	body, parseErr := ioutil.ReadAll(res.Body)
	if parseErr != nil {
		log.Fatal(parseErr)
	}
	c <- string(body)
}
