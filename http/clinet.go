package http

import(
	"fmt"
	"net/http"
	"bytes"
	"io/ioutil"
)

var clientInformations []*Client

type Client struct {
	Name string
	Url string
	Request Request
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


func (this Client) ReceiveRequest(methodType string, path string, requestBody Request) *Response {
	url := fmt.Sprintf("%s%s?%s", this.Url, path, requestBody.ToQueryParam())
	fmt.Println(url)
		
	request, _ := http.NewRequest(methodType, url, bytes.NewBuffer([]byte(requestBody.Json)))
	resp, err := http.DefaultClient.Do(request)
	defer resp.Body.Close()
	
	if err != nil {
		fmt.Println(err)
		return NewResponse(500, "BadRequest!!")
	}
	
	body, _ := ioutil.ReadAll(resp.Body)
	return NewResponse(resp.StatusCode, string(body))
}