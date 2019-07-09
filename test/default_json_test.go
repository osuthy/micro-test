package test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/ShoichiroKitano/micro_test/http"
	"github.com/ShoichiroKitano/micro_test/json"
	"github.com/ShoichiroKitano/micro_test/runner"
	"github.com/ShoichiroKitano/micro_test/test/wiremock"
)

func TestDBはデフォルト値を使ってデータのセットアップができる(t *testing.T) {
	t.Run("階層が1つの場合", func(t *testing.T) {
		defer wiremock.Reset("localhost:8080")
		wiremock.Stubbing("localhost:8080", "/test", "GET",
			`{ "o1": "v1", "o2": "d2"}`, 200, "test success")

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
				"o1": {
					"o2": {
						"o3": "v3"
					}
				}
		 }`, 200, "test success")

		http.DefineServer("test_server", "http://localhost:8080")
		defaultJson := json.O{"o1": json.O{"o2": json.O{"o3": "d3"}}}

		http.Server("test_server").
			ReceiveRequest("GET", "/test", http.WithJson(defaultJson.Override("o1", "o2", "o3", "v3"))).
			AndResponseShouldBe(http.Status(200).TextPlain("test success"))
		assert.Equal(t, "success", runner.TestRunner.Result)
	})

	t.Run("jsonで上書きする場合", func(t *testing.T) {
		defer wiremock.Reset("localhost:8080")
		wiremock.Stubbing("localhost:8080", "/test", "GET",
		`{ "o1": { "o12": { "o13": "v13" } }, "o2": "v21"}`,200, "test success")

		http.DefineServer("test_server", "http://localhost:8080")
		defaultJson := json.O{"o1": json.O{"o12": json.O{"o13": "d13"} }, "o2": "d21"}

		http.Server("test_server").
		ReceiveRequest("GET", "/test", http.WithJson(defaultJson.Override(json.O{ "o1": json.O{"o12": json.O{"o13": "v13"}}, "o2": "v21"}))).
		AndResponseShouldBe(http.Status(200).TextPlain("test success"))
		assert.Equal(t, "success", runner.TestRunner.Result)
	})

	t.Run("jsonが複数ある場合", func(t *testing.T) {
		defer wiremock.Reset("localhost:8080")
		wiremock.Stubbing("localhost:8080", "/test", "GET",
		`{
			"array": [
				{"o": "v1"}, {"o": "v2"}
				]
		}`, 200, "test success")

		http.DefineServer("test_server", "http://localhost:8080")

		http.Server("test_server").
		ReceiveRequest("GET", "/test", http.WithJson(json.O{"array": json.O{"o": "v"}.Generate(2)})).
		AndResponseShouldBe(http.Status(200).TextPlain("test success"))
		assert.Equal(t, "success", runner.TestRunner.Result)
	})

	t.Run("jsonの値が数値の場合", func(t *testing.T) {
		defer wiremock.Reset("localhost:8080")
		wiremock.Stubbing("localhost:8080", "/test", "GET",
		`{
			"array": [
				{"o1": 11}, {"o1": 12}
				]
		}`, 200, "test success")

		http.DefineServer("test_server", "http://localhost:8080")

		http.Server("test_server").
		ReceiveRequest("GET", "/test", http.WithJson(json.O{"array": json.O{"o1": 10}.Generate(2)})).
		AndResponseShouldBe(http.Status(200).TextPlain("test success"))
		assert.Equal(t, "success", runner.TestRunner.Result)
	})
}
