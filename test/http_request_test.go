package test

import(
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/ShoichiroKitano/micro_test/http"
	"github.com/ShoichiroKitano/micro_test/json"
)

func TestHttpはサーバーにメソッドGETでリクエストを送ることができる(t *testing.T) {
	var body string
	var status int
	status, body = http.Server("localhost:8080").
	ReceiveRequest("GET", "/test",
	http.WithJson(json.O{
		"object": "value",
		"array": json.A{1, 2, 3},
	}))
	assert.Equal(t, 200, status)
	assert.Equal(t, "test success", body)
}
