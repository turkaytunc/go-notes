package frequencycounter

import (
	"strings"
)

func New(arr string) map[string]int {

	freq := make(map[string]int)

	strArr := strings.Split(arr, "")

	for v := range strArr {
		_, ok := freq[strArr[v]]
		if ok {
			freq[strArr[v]]++
			continue
		}
		freq[strArr[v]] = 1
	}

	return freq
}
