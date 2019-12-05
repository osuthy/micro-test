package http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	. "github.com/osuthy/micro-test"
)

var clientInformations []*Client

type Client struct {
	Name    string
	Url     string
	Request Request
	differences *Differences
}

func DefineServer(serverName string, url string) {
	clientInformations = append(clientInformations, &Client{Name: serverName, Url: url})
}

func Server(tc TC, serverName string) *Client {
	for _, clientInfo := range clientInformations {
		if clientInfo.Name == serverName {
			clientInfo.differences = tc["differences"].(*Differences)
			return clientInfo
		}
	}
	return nil
}

func (this Client) ReceiveRequest(methodType string, path string, requestBody Request) *Response {
	url := fmt.Sprintf("%s%s?%s", this.Url, path, requestBody.ToQueryParam())

	request, _ := http.NewRequest(methodType, url, bytes.NewBuffer([]byte(requestBody.Json)))
	resp, err := http.DefaultClient.Do(request)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println(err)
		return NewResponse(500, "BadRequest!!")
	}

	body, _ := ioutil.ReadAll(resp.Body)
	r := NewResponse(resp.StatusCode, string(body))
	r.differences = this.differences
	return r
}
