package tests

import (
	"lite-api-crud/controllers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Welcomepage(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil) // args of req don't matter fix it with routes
	recorder := httptest.NewRecorder()
	controllers.WelcomePage(recorder, request)

	gotBody := recorder.Body.String()
	gotCode := recorder.Result().StatusCode
	gotType := recorder.Header().Get("Content-Type")

	if gotBody == "" || gotCode != 200 || gotType != "application/json" {
		t.Errorf("WelcomePage fails, got body: %v - code: %d - content-type: %v", gotBody, gotCode, gotType)
	}
}
