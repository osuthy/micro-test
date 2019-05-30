package test

import(
	"testing"
	"github.com/stretchr/testify/assert"

	"github.com/ShoichiroKitano/micro_test/http"
	"github.com/ShoichiroKitano/micro_test/json"
	"github.com/ShoichiroKitano/micro_test/runner"

	"github.com/ShoichiroKitano/micro_test/test/wiremock"
)
func TestHttpはサーバーにJSONを送ることができる(t *testing.T) {
	wiremock.Reset("localhost:8080")

	t.Run("Objectがトップレベルの階層の場合", func(t *testing.T) {
		defer wiremock.Reset("localhost:8080")
		wiremock.Stubbing("localhost:8080", "/test", "GET",
		`{ \"object\": \"value\" }`, 200, "test success")

		http.DefineServer("test_server", "http://localhost:8080")
		status, body := http.Server("test_server").
		ReceiveRequest("GET", "/test", http.WithJson(json.O{"object": "value"})).
			AndResponseShouldBe(Status(200).Text("test success"))
		assert.Equal(t, "success", runner.TestRunner.Result)
	})

	t.Run("Arrayがトップレベルの階層の場合", func(t *testing.T) {
		runner.TestRunner.Result = "init"
		defer wiremock.Reset("localhost:8080")
		wiremock.Stubbing("localhost:8080", "/test", "GET",
		`[1, 2, 3]`, 200, "test success")

		http.DefineServer("test_server", "http://localhost:8080")
		status, body := http.Server("test_server").
		ReceiveRequest("GET", "/test", http.WithJson(json.A{1, 2, 3})).
			AndResponseShouldBe(Status(200).Text("test success"))
		assert.Equal(t, "success", runner.TestRunner.Result)
	})
}

func TestHttpはサーバーにPOSTでリクエストを送ることができる(t *testing.T) {
	wiremock.Reset("localhost:8080")
	runner.TestRunner.Result = "init"
	defer wiremock.Reset("localhost:8080")

	wiremock.Stubbing("localhost:8080", "/test", "POST",
	`{ \"object\": \"value\" }`, 200, "test success")

	http.DefineServer("test_server", "http://localhost:8080")
	http.Server("test_server").
		ReceiveRequest("POST", "/test", http.WithJson(json.O{"object": "value"})).
		AndResponseShouldBe(Status(200).Text("test success"))
	assert.Equal(t, "success", runner.TestRunner.Result)
}

func TestHttpはサーバーはレスポンスが期待と異なる場合はテストを失敗させる(t *testing.T) {
	wiremock.Reset("localhost:8080")
	runner.TestRunner.Result = "init"
	defer wiremock.Reset("localhost:8080")

	http.DefineServer("test_server", "http://localhost:8080")
	http.Server("test_server").
		ReceiveRequest("POST", "/test", http.WithJson(json.O{"object": "value"})).
		AndResponseShouldBe(Status(200).Text("test success"))
	assert.Equal(t, "", runner.TestRunner.Result)
}

