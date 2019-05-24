package wiremock

import (
	"net/http"
)

func Reset(url string) {
	req, _ := http.NewRequest("POST", "http://" + url +"/__admin/reset", nil)
	http.DefaultClient.Do(req)
}

