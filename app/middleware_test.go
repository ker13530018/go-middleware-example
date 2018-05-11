package app

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_middlewareSomething(t *testing.T) {
	srv := Server{
		Env: "testing",
	}
	srv.Routes()
	request, err := http.NewRequest("GET", "/api/user", nil)
	if err != nil {
		t.Fatal(err)
	}

	request.Header.Add("Authorization", "Basic c25vb2tlcjo0OTVqMzI5OGRnc2RhdA==")
	respRecord := httptest.NewRecorder()
	srv.Router.ServeHTTP(respRecord, request)
	checkResponseCode(t, http.StatusOK, respRecord.Code)
}
