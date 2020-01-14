package http

import (
	"bytes"
	"fmt"
	. "github.com/osuthy/micro-test"
	"io/ioutil"
	"net/http"
)

type RequestDSL struct {
	client      *Client
	differences *Differences
}

func Server(tc TC, serverName string) *RequestDSL {
	return &RequestDSL{
		client:      tc[serverName].(*Client),
		differences: tc["differences"].(*Differences),
	}
}

func (this *RequestDSL) ReceiveRequest(methodType string, path string, requestBody Request) *Response {
	urlWithQueryParam := fmt.Sprintf("%s%s?%s", this.client.Url, path, requestBody.ToQueryParam())
	request, _ := http.NewRequest(methodType, urlWithQueryParam, bytes.NewBuffer([]byte(requestBody.Json)))
	resp, err := http.DefaultClient.Do(request)
	defer resp.Body.Close()

	if err != nil {
		return NewResponse(500, "BadRequest!!")
	}

	body, _ := ioutil.ReadAll(resp.Body)
	r := NewResponse(resp.StatusCode, string(body))
	r.differences = this.differences
	return r
}

type ResponseDSL struct {
	response    *Response
	differences *Differences
}

func (this Response) AndResponseShouldBe(expected *Response) {
	if !this.Equal(expected) {
		this.differences.Push("assert is fail!!!!!!!!!!!1")
	}
}
