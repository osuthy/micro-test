package test

import (
	"testing"
	"github.com/stretchr/testify/assert"

	"github.com/ShoichiroKitano/micro_test/test/wiremock"
	"github.com/ShoichiroKitano/micro_test/http"
	"github.com/ShoichiroKitano/micro_test/json"
	"github.com/ShoichiroKitano/micro_test/runner"
)

func TestDBはデフォルト値を使ってデータのセットアップができる(t *testing.T) {
	t.Run("階層が1つの場合", func(t *testing.T) {
		defer wiremock.Reset("localhost:8080")
		wiremock.Stubbing("localhost:8080", "/test", "GET",
			`{ \"o1\": \"v1\", \"o2\": \"d2\"}`, 200, "test success")

		http.DefineServer("test_server", "http://localhost:8080")
		defaultJson := json.O{"o1": "d1", "o2": "d2"}

		http.Server("test_server").
			ReceiveRequest("GET", "/test", http.WithJson(defaultJson.Override("o1", "v1"))).
			AndResponseShouldBe(http.Status(200).TextPlain("test success"))
		assert.Equal(t, "success", runner.TestRunner.Result)
	})

	t.Run("階層が2つの場合", func(t *testing.T) {
		defer wiremock.Reset("localhost:8080")
		wiremock.Stubbing("localhost:8080", "/test", "GET",
			`{
					\"o1\": {
						\"o2\": {
							\"o3"\: \"v3\"
						}
					},
			 }`,200, "test success")

		http.DefineServer("test_server", "http://localhost:8080")
		defaultJson := json.O{"o1": json.O{"o2": json.O{"o3": "d3"} }}

		http.Server("test_server").
			ReceiveRequest("GET", "/test", http.WithJson(defaultJson.Override("o1", "o2", "o3", "v3"))).
			AndResponseShouldBe(http.Status(200).TextPlain("test success"))
		assert.Equal(t, "success", runner.TestRunner.Result)
	})
}
