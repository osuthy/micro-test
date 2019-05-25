package test

import(
	"testing"
	"github.com/stretchr/testify/assert"

	"github.com/ShoichiroKitano/micro_test/http"
	"github.com/ShoichiroKitano/micro_test/json"

	"github.com/ShoichiroKitano/micro_test/test/wiremock"
)

func TestHttpはサーバーにJSONを送ることができる(t *testing.T) {
	t.Run("Objectがトップレベルの階層の場合", func(t *testing.T) {
		defer wiremock.Reset("localhost:8080")
		wiremock.Stubbing("localhost:8080", "/test", "GET",
		`{ \"object\": \"value\", \"array\": [1, 2, 3] }`, 200, "test success")
		status, body := http.Server("localhost:8080").
		ReceiveRequest("GET", "/test",
		http.WithJson(json.O{
			"object": "value",
			"array": json.A{1, 2, 3},
		}))
		assert.Equal(t, 200, status)
		assert.Equal(t, "test success", body)
	})
}
