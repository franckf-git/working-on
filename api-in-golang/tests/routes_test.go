package tests

import (
	"bytes"
	"encoding/json"
	"lite-api-crud/config"
	router "lite-api-crud/routers"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

var apiTest = router.App{}

func TestMain(m *testing.M) {
	apiTest.Initialize()
	code := m.Run()
	os.Exit(code)
}

func Test_WelcomePage(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	responseRec := httptest.NewRecorder()
	apiTest.Router.ServeHTTP(responseRec, request)

	gotBody := responseRec.Body.Bytes()
	gotCode := responseRec.Result().StatusCode
	gotType := responseRec.Header().Get("Content-Type")
	gotJSON := config.Message{}
	json.Unmarshal(gotBody, &gotJSON)

	if gotJSON.Status != "information" {
		t.Errorf("WelcomePage fails, got status: %v", gotJSON.Status)
	}
	if gotJSON.Message == "" {
		t.Errorf("WelcomePage fails, message is empty")
	}
	if gotCode != 200 {
		t.Errorf("WelcomePage fails, got code: %d", gotCode)
	}
	if gotType != "application/json" {
		t.Errorf("WelcomePage fails, got content-type: %v", gotType)
	}
}

func Test_Docs(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/v1/docs", nil)
	responseRec := httptest.NewRecorder()
	apiTest.Router.ServeHTTP(responseRec, request)

	gotCode := responseRec.Result().StatusCode
	if gotCode != 301 {
		t.Errorf("Docs redirect fails, got code: %d", gotCode)
	}
}

func Test_ShowAllPosts(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/v1/posts", nil)
	responseRec := httptest.NewRecorder()
	apiTest.Router.ServeHTTP(responseRec, request)

	gotBody := responseRec.Body.Bytes()
	gotCode := responseRec.Result().StatusCode
	gotType := responseRec.Header().Get("Content-Type")
	gotJSON := []config.Post{}
	json.Unmarshal(gotBody, &gotJSON)

	// carefull very couple to model_test.RegisterPosts
	gotJSON[0].Created = fakeCreatedTime
	gotJSON[1].Created = fakeCreatedTime
	want := []config.Post{
		{Id: 1, Title: "title1", Datas: "datas1", Created: fakeCreatedTime, IdUser: 1},
		{Id: 2, Title: "title2", Datas: "datas2", Created: fakeCreatedTime, IdUser: 2},
	}
	if !reflect.DeepEqual(gotJSON, want) {
		t.Errorf("ShowAllPosts fail, got datas: %v", gotJSON)
	}
	if gotCode != 200 {
		t.Errorf("ShowAllPosts fails, got code: %d", gotCode)
	}
	if gotType != "application/json" {
		t.Errorf("ShowAllPosts fails, got content-type: %v", gotType)
	}
}

func Test_AddPost(t *testing.T) {
	body := []byte(`{"title":"add test post","datas":"datasfill","idUser":99}`)
	request, _ := http.NewRequest("POST", "/api/v1/post", bytes.NewBuffer(body))
	responseRec := httptest.NewRecorder()
	apiTest.Router.ServeHTTP(responseRec, request)

	gotBody := responseRec.Body.Bytes()
	gotCode := responseRec.Result().StatusCode
	gotType := responseRec.Header().Get("Content-Type")
	gotJSON := config.Message{}
	json.Unmarshal(gotBody, &gotJSON)

	if gotJSON.Status != "success" {
		t.Errorf("AddPost fails, got status: %v", gotJSON.Status)
	}
	if gotJSON.Message == "" {
		t.Errorf("AddPost fails, message is empty")
	}
	if gotJSON.Id != 3 {
		t.Errorf("AddPost fails, got id: %d", gotJSON.Id)
	}
	if gotCode != 200 {
		t.Errorf("AddPost fails, got code: %d", gotCode)
	}
	if gotType != "application/json" {
		t.Errorf("AddPost fails, got content-type: %v", gotType)
	}
}
