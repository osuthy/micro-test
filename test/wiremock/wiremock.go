package wiremock

import (
	"bytes"
	"net/http"
	"strconv"
	"strings"
)

func Reset(url string) {
	req, _ := http.NewRequest("POST", "http://"+url+"/__admin/reset", nil)
	http.DefaultClient.Do(req)
}

func Stubbing(url string, path string, method string, json string, status int, body string) {
	jsonForWiremock := strings.Replace(json, "\"", "\\\"", -1)
	jsonForWiremock2 := strings.Replace(jsonForWiremock, "\t", "", -1)
	jsonForWiremock3 := strings.Replace(jsonForWiremock2, "\n", "", -1)
	req, _ := http.NewRequest("POST", "http://"+url+"/__admin/mappings/new",
		bytes.NewBuffer([]byte(`
		{
			"request":
				{
					"url": "`+path+`",
					"method": "`+method+`",
					"bodyPatterns": [{
						"equalToJson": "`+jsonForWiremock3+`"
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
