package main

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/ozontech/cute"
	"github.com/ozontech/cute/asserts/json"
)

func TestIndexHandler(t *testing.T) {
	cute.NewTestBuilder().
		Title("Positiv test #1").
		Description("This is very simple positiv GET test for /").
		Create().
		RequestBuilder(
			cute.WithURI("http://localhost:8090/"),
			cute.WithMethod(http.MethodGet),
		).
		ExpectExecuteTimeout(3*time.Second).
		ExpectStatus(http.StatusOK).
		AssertBody(
			json.Equal("status", "ok"),
		).
		ExecuteTest(context.Background(), t)
}
