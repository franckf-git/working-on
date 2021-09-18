package tests

import (
	router "lite-api-crud/routers"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var apiTest = router.App{}

func TestMain(m *testing.M) {
	apiTest.Initialize()
	code := m.Run()
	os.Exit(code)
}

func Test_Welcomepage(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	responseRec := httptest.NewRecorder()
	apiTest.Router.ServeHTTP(responseRec, request)

	gotBody := responseRec.Body.String()
	gotCode := responseRec.Result().StatusCode
	gotType := responseRec.Header().Get("Content-Type")

	if gotBody == "" {
		t.Errorf("WelcomePage fails, got body: %v", gotBody)
	}
	if gotCode != 200 {
		t.Errorf("WelcomePage fails, got code: %d", gotCode)
	}
	if gotType != "application/json" {
		t.Errorf("WelcomePage fails, got content-type: %v", gotType)
	}
}
