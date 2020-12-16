package main_test

import (

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	t.Run("Non 200 Response", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(500)
		}))

		defer ts.Close()

		DefaultHTTPGetAddress = ts.URL

		_, err := handler(events.APIGatewayProxyRequest{})
		if err != nil && err.Error() != ErrNon200Response.Error() {
			t.Fatal("Error failed to trigger with an invalid request: %v", err)
		}
	})
}