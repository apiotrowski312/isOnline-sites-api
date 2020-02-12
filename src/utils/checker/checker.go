package checker

import "net/http"

func GetStatus(url string) int64 {
	resp, err := http.Get(url)

	if err != nil {
		return -1
	}

	return int64(resp.StatusCode)
}
