package wiremock

import (
	"bytes"
	"net/http"
	"strconv"
)

func Reset(url string) {
	req, _ := http.NewRequest("POST", "http://"+url+"/__admin/reset", nil)
	http.DefaultClient.Do(req)
}

func Stubbing(url string, path string, method string, json string, status int, body string) {
	req, _ := http.NewRequest("POST", "http://"+url+"/__admin/mappings/new",
		bytes.NewBuffer([]byte(`
		{
			"request":
				{
					"url": "`+path+`",
					"method": "`+method+`",
					"bodyPatterns": [{
						"equalToJson": "`+json+`"
					}]
				},
			"response":
				{
					"status": `+strconv.Itoa(status)+`,
					"body": "`+body+`"
				}
		}`)))
	http.DefaultClient.Do(req)
}
