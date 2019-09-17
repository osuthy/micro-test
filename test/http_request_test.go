package test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/ShoichiroKitano/micro_test/http"
	"github.com/ShoichiroKitano/micro_test/json"
	"github.com/ShoichiroKitano/micro_test/runner"
	. "github.com/ShoichiroKitano/micro_test"

	"github.com/ShoichiroKitano/micro_test/test/wiremock"
)

func TestHttpはサーバーにJSONを送ることができる(t *testing.T) {
	wiremock.Reset("localhost:8080")

	t.Run("Objectがトップレベルの階層の場合", func(t *testing.T) {
		defer wiremock.Reset("localhost:8080")
		defer resetSuites()
		resetSuites()
		spy := setUpSpy()

		Describe("A", func() {
			It("B", func() {
				wiremock.Stubbing("localhost:8080", "/test", "GET",
					`{ "object": "value" }`, 200, "test success")

				http.DefineServer("test_server", "http://localhost:8080")
				http.Server("test_server").
					ReceiveRequest("GET", "/test", http.WithJson(json.O{"object": "value"})).
					AndResponseShouldBe(http.Status(200).TextPlain("test success"))
			})
		})

		runner.Run()
		assert.Equal(t, 0, len(spy.results))
	})

	t.Run("Arrayがトップレベルの階層の場合", func(t *testing.T) {
		runner.TestRunner.Result = "init"
		defer wiremock.Reset("localhost:8080")
		wiremock.Stubbing("localhost:8080", "/test", "GET",
			`[1, 2, 3]`, 200, "test success")

		http.DefineServer("test_server", "http://localhost:8080")
		http.Server("test_server").
			ReceiveRequest("GET", "/test", http.WithJson(json.A{1, 2, 3})).
			AndResponseShouldBe(http.Status(200).TextPlain("test success"))
		assert.Equal(t, "success", runner.TestRunner.Result)
	})
}

func TestHttpはサーバーにPOSTでリクエストを送ることができる(t *testing.T) {
	wiremock.Reset("localhost:8080")
	runner.TestRunner.Result = "init"
	defer wiremock.Reset("localhost:8080")

	wiremock.Stubbing("localhost:8080", "/test", "POST",
		`{ "object": "value" }`, 200, "test success")

	http.DefineServer("test_server", "http://localhost:8080")
	http.Server("test_server").
		ReceiveRequest("POST", "/test", http.WithJson(json.O{"object": "value"})).
		AndResponseShouldBe(http.Status(200).TextPlain("test success"))
	assert.Equal(t, "success", runner.TestRunner.Result)
}

func TestHttpはサーバーはレスポンスが期待と異なる場合はテストを失敗させる(t *testing.T) {
	wiremock.Reset("localhost:8080")
	runner.TestRunner.Result = "init"
	defer wiremock.Reset("localhost:8080")

	http.DefineServer("test_server", "http://localhost:8080")
	http.Server("test_server").
		ReceiveRequest("POST", "/test", http.WithJson(json.O{"object": "value"})).
		AndResponseShouldBe(http.Status(200).TextPlain("test success"))
	assert.Equal(t, "", runner.TestRunner.Result)
}

func TestHttpはパラメータ付きのPOSTのリクエストを送ることができる(t *testing.T) {
	wiremock.Reset("localhost:8080")
	runner.TestRunner.Result = "init"
	defer wiremock.Reset("localhost:8080")

	wiremock.Stubbing("localhost:8080", "/test?param1=p1&param2=p2", "POST",
		`{ "object": "value" }`, 200, "test success")

	http.DefineServer("test_server", "http://localhost:8080")
	http.Server("test_server").
		ReceiveRequest("POST", "/test",
			http.WithParam("param1", "p1").
				WithJson(json.O{"object": "value"}).
				WithParam("param2", "p2")).AndResponseShouldBe(http.Status(200).TextPlain("test success"))
	assert.Equal(t, "success", runner.TestRunner.Result)
}

func TestHttpはパラメータ付きのPATCHのリクエストを送ることができる(t *testing.T) {
	wiremock.Reset("localhost:8080")
	runner.TestRunner.Result = "init"
	defer wiremock.Reset("localhost:8080")

	wiremock.Stubbing("localhost:8080", "/test?param1=p1&param2=p2", "PATCH",
		`{ "object": "value" }`, 200, "test success")

	http.DefineServer("test_server", "http://localhost:8080")
	http.Server("test_server").
		ReceiveRequest("PATCH", "/test",
			http.WithParam("param1", "p1").
				WithJson(json.O{"object": "value"}).
				WithParam("param2", "p2")).AndResponseShouldBe(http.Status(200).TextPlain("test success"))
	assert.Equal(t, "success", runner.TestRunner.Result)
}

func TestHttpはパラメータ付きのPUTのリクエストを送ることができる(t *testing.T) {
	wiremock.Reset("localhost:8080")
	runner.TestRunner.Result = "init"
	defer wiremock.Reset("localhost:8080")

	wiremock.Stubbing("localhost:8080", "/test?param1=p1&param2=p2", "PUT",
		`{ "object": "value" }`, 200, "test success")

	http.DefineServer("test_server", "http://localhost:8080")
	http.Server("test_server").
		ReceiveRequest("PUT", "/test",
			http.WithParam("param1", "p1").
				WithJson(json.O{"object": "value"}).
				WithParam("param2", "p2")).AndResponseShouldBe(http.Status(200).TextPlain("test success"))
	assert.Equal(t, "success", runner.TestRunner.Result)
}

func TestHttpはパラメータ付きのDELETEのリクエストを送ることができる(t *testing.T) {
	wiremock.Reset("localhost:8080")
	runner.TestRunner.Result = "init"
	defer wiremock.Reset("localhost:8080")

	wiremock.Stubbing("localhost:8080", "/test?param1=p1&param2=p2", "DELETE",
		`{ "object": "value" }`, 200, "test success")

	http.DefineServer("test_server", "http://localhost:8080")
	http.Server("test_server").
		ReceiveRequest("DELETE", "/test",
			http.WithParam("param1", "p1").
				WithJson(json.O{"object": "value"}).
				WithParam("param2", "p2")).AndResponseShouldBe(http.Status(200).TextPlain("test success"))
	assert.Equal(t, "success", runner.TestRunner.Result)
}
