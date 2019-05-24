package http

import(
	"fmt"
	"net/http"
	"bytes"
	"io/ioutil"
	. "github.com/ShoichiroKitano/micro_test/json"
)

type Client struct {
	Adress string
}

func Server(adress string) *Client {
	client := new(Client)
	client.Adress = adress
	return client
}

func (this Client) ReceiveRequest(methodType string, path string, json []byte) (int, string) {
	url := fmt.Sprintf("http://%s%s",this.Adress, path)

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

func WithJson(object O) []byte {
	return object.ToJson()
}

