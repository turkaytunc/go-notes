package crawl

import (
	"testing"
)

type testData struct {
	url  string
	want int
}

func TestSiteBodyLength(t *testing.T) {

	testCases := []testData{
		{url: "https://google.com/", want: 5100},
		{url: "https://amazon.com/", want: 1000},
		{url: "https://aws.com/", want: 5000},
	}

	for _, v := range testCases {
		c := make(chan int)

		go SiteBodyLength(v.url, c)
		got := <-c

		if got < v.want {
			t.Errorf("error: got %v , want %v", got, v.want)
		}

	}

}
