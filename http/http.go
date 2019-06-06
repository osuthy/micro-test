package http

import(
	"fmt"
	"net/http"
	"bytes"
	"io/ioutil"
	. "github.com/ShoichiroKitano/micro_test/json"
	"github.com/ShoichiroKitano/micro_test/runner"
)

var clientInformations []*Client

type Client struct {
	Name string
	Url string
}

type Response struct {
	Status int
	Body string
}

type ExpectedResponse struct {
	Status int
	Body string
}

func DefineServer(serverName string, url string) {
	client := new(Client)
	client.Name = serverName
	client.Url = url
	clientInformations = append(clientInformations, client)
}

func Server(serverName string) *Client {
	for _, clientInfo := range clientInformations {
		if clientInfo.Name == serverName {
			return clientInfo
		}
	}
	return nil
}

func (this Client) ReceiveRequest(methodType string, path string, json []byte) Response {
	url := fmt.Sprintf("%s%s",this.Url, path)

	request, _ := http.NewRequest(methodType, url, bytes.NewBuffer([]byte(json)))
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		return Response{500, "BadRequest!!"}
	  }
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return Response{resp.StatusCode, string(body)}
}

func (this Response) AndResponseShouldBe(expected ExpectedResponse) (int, string) {
	if this.IsSame(expected) {
		runner.TestRunner.Result = "success"
		return 200, "test success"
	} else {
		runner.TestRunner.Result = ""
		return 500, ""
	}
}

func (this Response) IsSame(expected ExpectedResponse) bool {
	if this.Status == expected.Status && this.Body == expected.Body {
		return true
	} else {
		return false
	}
}

func Status(expectedStatus int) *ExpectedResponse {
	resp := new(ExpectedResponse)
	resp.Status = expectedStatus
	return resp
}

func (this ExpectedResponse) TextPlain(expectedBody string) ExpectedResponse {
	this.Body = expectedBody
	return this
}

func WithJson(i Object) []byte {
	return i.ToJson()
}
