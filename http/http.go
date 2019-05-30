package http

import(
	"fmt"
	"net/http"
	"bytes"
	"io/ioutil"
	. "github.com/ShoichiroKitano/micro_test/json"
)

var clientInformations []*Client

type Client struct {
	Name string
	Url string
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

func (this Client) ReceiveRequest(methodType string, path string, json []byte) (int, string) {
	url := fmt.Sprintf("%s%s",this.Url, path)

	request, _ := http.NewRequest(methodType, url, bytes.NewBuffer([]byte(json)))
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		return 500, "BadRequest!!"
	  }
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return resp.StatusCode, string(body)
}

func WithJson(i Object) []byte {
	return i.ToJson()
}
