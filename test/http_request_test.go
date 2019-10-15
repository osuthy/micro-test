package test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/osuthy/micro-test/http"
	"github.com/osuthy/micro-test/json"
	"github.com/osuthy/micro-test/runner"
	. "github.com/osuthy/micro-test"

	"github.com/osuthy/micro-test/test/wiremock"
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
		defer wiremock.Reset("localhost:8080")
		defer resetSuites()
		resetSuites()
		spy := setUpSpy()

		Describe("A", func() {
			It("B", func() {
				wiremock.Stubbing("localhost:8080", "/test", "GET",
					`[1, 2, 3]`, 200, "test success")

				http.DefineServer("test_server", "http://localhost:8080")
				http.Server("test_server").
					ReceiveRequest("GET", "/test", http.WithJson(json.A{1, 2, 3})).
					AndResponseShouldBe(http.Status(200).TextPlain("test success"))
			})
		})

		runner.Run()
		assert.Equal(t, 0, len(spy.results))
	})
}

func TestHttpはサーバーにPOSTでリクエストを送ることができる(t *testing.T) {
	defer wiremock.Reset("localhost:8080")
	defer resetSuites()

	wiremock.Reset("localhost:8080")
	resetSuites()
	spy := setUpSpy()

	Describe("A", func() {
		It("B", func() {
			wiremock.Stubbing("localhost:8080", "/test", "POST",
				`{ "object": "value" }`, 200, "test success")

			http.DefineServer("test_server", "http://localhost:8080")
			http.Server("test_server").
				ReceiveRequest("POST", "/test", http.WithJson(json.O{"object": "value"})).
				AndResponseShouldBe(http.Status(200).TextPlain("test success"))
		})
	})

	runner.Run()
	assert.Equal(t, 0, len(spy.results))
}

func TestHttpはサーバーはレスポンスが期待と異なる場合はテストを失敗させる(t *testing.T) {
	defer wiremock.Reset("localhost:8080")
	defer resetSuites()

	wiremock.Reset("localhost:8080")
	resetSuites()
	spy := setUpSpy()

	Describe("A", func() {
		It("B", func() {
			http.DefineServer("test_server", "http://localhost:8080")
			http.Server("test_server").
				ReceiveRequest("POST", "/test", http.WithJson(json.O{"object": "value"})).
				AndResponseShouldBe(http.Status(200).TextPlain("test success"))
		})
	})
	runner.Run()
	assert.Equal(t, "A B", spy.results[0])
	assert.Equal(t, 2, len(spy.results))
}

func TestHttpはパラメータ付きのPOSTのリクエストを送ることができる(t *testing.T) {
	defer wiremock.Reset("localhost:8080")
	defer resetSuites()

	wiremock.Reset("localhost:8080")
	resetSuites()
	spy := setUpSpy()

	Describe("A", func() {
		It("B", func() {
			wiremock.Stubbing("localhost:8080", "/test?param1=p1&param2=p2", "POST",
				`{ "object": "value" }`, 200, "test success")

			http.DefineServer("test_server", "http://localhost:8080")
			http.Server("test_server").
				ReceiveRequest("POST", "/test",
					http.WithParam("param1", "p1").
						WithJson(json.O{"object": "value"}).
						WithParam("param2", "p2")).AndResponseShouldBe(http.Status(200).TextPlain("test success"))
		})
	})
	runner.Run()
	assert.Equal(t, 0, len(spy.results))
}

func TestHttpはパラメータ付きのPATCHのリクエストを送ることができる(t *testing.T) {
	defer wiremock.Reset("localhost:8080")
	defer resetSuites()

	wiremock.Reset("localhost:8080")
	resetSuites()
	spy := setUpSpy()

	Describe("A", func() {
		It("B", func() {
			wiremock.Stubbing("localhost:8080", "/test?param1=p1&param2=p2", "PATCH",
				`{ "object": "value" }`, 200, "test success")

			http.DefineServer("test_server", "http://localhost:8080")
			http.Server("test_server").
				ReceiveRequest("PATCH", "/test",
					http.WithParam("param1", "p1").
						WithJson(json.O{"object": "value"}).
						WithParam("param2", "p2")).AndResponseShouldBe(http.Status(200).TextPlain("test success"))
		})
	})
	runner.Run()
	assert.Equal(t, 0, len(spy.results))
}

func TestHttpはパラメータ付きのPUTのリクエストを送ることができる(t *testing.T) {
	defer wiremock.Reset("localhost:8080")
	defer resetSuites()

	wiremock.Reset("localhost:8080")
	resetSuites()
	spy := setUpSpy()

	Describe("A", func() {
		It("B", func() {
			wiremock.Stubbing("localhost:8080", "/test?param1=p1&param2=p2", "PUT",
				`{ "object": "value" }`, 200, "test success")

			http.DefineServer("test_server", "http://localhost:8080")
			http.Server("test_server").
				ReceiveRequest("PUT", "/test",
					http.WithParam("param1", "p1").
						WithJson(json.O{"object": "value"}).
						WithParam("param2", "p2")).AndResponseShouldBe(http.Status(200).TextPlain("test success"))
		})
	})
	runner.Run()
	assert.Equal(t, 0, len(spy.results))
}

func TestHttpはパラメータ付きのDELETEのリクエストを送ることができる(t *testing.T) {
	defer wiremock.Reset("localhost:8080")
	defer resetSuites()

	wiremock.Reset("localhost:8080")
	resetSuites()
	spy := setUpSpy()

	Describe("A", func() {
		It("B", func() {
			wiremock.Stubbing("localhost:8080", "/test?param1=p1&param2=p2", "DELETE",
				`{ "object": "value" }`, 200, "test success")

			http.DefineServer("test_server", "http://localhost:8080")
			http.Server("test_server").
				ReceiveRequest("DELETE", "/test",
					http.WithParam("param1", "p1").
						WithJson(json.O{"object": "value"}).
						WithParam("param2", "p2")).AndResponseShouldBe(http.Status(200).TextPlain("test success"))
		})
	})
	runner.Run()
	assert.Equal(t, 0, len(spy.results))
}
