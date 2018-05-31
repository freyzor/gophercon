package routing_test

import (
	"testing"
	"net/http"
	"net/http/httptest"

	"github.com/freyzor/gophercon/pkg/routing"
)

func TestBaseRouter(t *testing.T) {
	handler := routing.BaseRouter()

	ts := httptest.NewServer(handler)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/home")
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Error("")
	}
}
