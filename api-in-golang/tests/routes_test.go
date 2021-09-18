package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func Test_AddPosts(t *testing.T) {
	body := []byte(`{"title":"add test post","datas":"datasfill","idUser":99}`)
	request, _ := http.NewRequest("POST", "/api/v1/post", bytes.NewBuffer(body))
	responseRec := httptest.NewRecorder()
	apiTest.Router.ServeHTTP(responseRec, request)

	gotBody := responseRec.Body.Bytes()
	gotCode := responseRec.Result().StatusCode
	gotType := responseRec.Header().Get("Content-Type")
	fmt.Println(gotBody)
	if gotCode != 200 {
		t.Errorf("AddPosts fails, got code: %d", gotCode)
	}
	if gotType != "application/json" {
		t.Errorf("AddPosts fails, got content-type: %v", gotType)
	}
}
