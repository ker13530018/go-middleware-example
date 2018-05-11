package app

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_handleSomething(t *testing.T) {
	srv := Server{
		Env: "testing",
	}
	srv.Routes()
	request, err := http.NewRequest("GET", "/api", nil)
	if err != nil {
		t.Fatal(err)
	}
	respRecord := httptest.NewRecorder()
	srv.Router.ServeHTTP(respRecord, request)
	checkResponseCode(t, http.StatusOK, respRecord.Code)
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
