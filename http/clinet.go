package http

import(
	"fmt"
	"net/http"
	"bytes"
	"io/ioutil"
	"strings"
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

	var params string

	if len(requestBody.Params)>0 {
		var builder strings.Builder
		for _, param := range requestBody.Params {
			if builder.String() != "" {
				p := fmt.Sprintf("&%s=%s", param.Name, param.Value)
				builder.WriteString(p)
			} else {
				p := fmt.Sprintf("%s=%s", param.Name, param.Value)
				builder.WriteString(p)
			}
		}
		params = builder.String()
	}

	url := fmt.Sprintf("%s%s?%s", this.Url, path, params)
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