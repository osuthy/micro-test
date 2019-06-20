package test

import (
	"testing"

	. "github.com/ShoichiroKitano/micro_test/db/infra"

	"github.com/ShoichiroKitano/micro_test/test/wiremock"
	"github.com/ShoichiroKitano/micro_test/db"
)

func TestDBはデフォルト値を使ってデータのセットアップができる(t *testing.T) {
	defer wiremock.Reset("localhost:8080")
	wiremock.Stubbing("localhost:8080", "/test", "GET",
		`{ \"o1\": \"v1\", \"o2\": \"d2\"}`, 200, "test success")

	http.DefineServer("test_server", "http://localhost:8080")
	defaultJson := json.O{"o1": "d1", "o2": "d2"}

	http.Server("test_server").
		ReceiveRequest("GET", "/test", http.WithJson(defaultJson.Override("o1", "v1"))).
		AndResponseShouldBe(http.Status(200).TextPlain("test success"))
	assert.Equal(t, "success", runner.TestRunner.Result)
}
