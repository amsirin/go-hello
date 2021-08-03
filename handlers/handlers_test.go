package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouterHomeWithSuccess(t *testing.T) {
	r := Router()
	ts := httptest.NewServer(r)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/hello")
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Status code for /hello is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusOK)
	}
}
