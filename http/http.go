package http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	. "github.com/osuthy/micro-test"
)


func DefineHttpServer(config C) {
	AppendConnectionDefinable(&HttpServerDefinition{
		config: config,
	})
}

type HttpServerDefinition struct {
	config C
}

func (this *HttpServerDefinition) SetConnectionForK8S(tc TC, namespace string) TC {
	return tc
}

func (this *HttpServerDefinition) SetConnectionForLocal(tc TC) TC {
	localConfig := this.config["local"].(C)
	url := fmt.Sprintf("http://%s:%d", localConfig["host"].(string), localConfig["port"].(int))
	tc[this.config["name"].(string)] = &Client{url: url}
	return tc
}


func (this Response) AndResponseShouldBe(expected *Response) {
	if !this.Equal(expected) {
		this.differences.Push("assert is fail!!!!!!!!!!!1")
	}
}

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
	urlWithQueryParam := fmt.Sprintf("%s%s?%s", this.client.url, path, requestBody.ToQueryParam())
	request, _ := http.NewRequest(methodType, urlWithQueryParam, bytes.NewBuffer([]byte(requestBody.Json)))
	resp, err := http.DefaultClient.Do(request)
	p(err)
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

type ResponseDSL struct {
	response    *Response
	differences *Differences
}

