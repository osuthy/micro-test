package http

import(
	"fmt"
	"net/http"
	"bytes"
	"reflect"
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

func (this Client) ReceiveRequest(methodType string, path string, json []byte) {
	url := fmt.Sprintf("%s %s ",this.Adress, path)
	fmt.Println("######################")
	fmt.Println(url)

	reqest, _ := http.NewRequest(methodType, url, bytes.NewBuffer([]byte(json)))
	fmt.Println(reflect.TypeOf(reqest))
}

func WithJson(object O) []byte {
	return object.ToJson()
}

