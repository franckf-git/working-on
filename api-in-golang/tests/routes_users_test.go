package tests

import (
	"bytes"
	"encoding/json"
	"lite-api-crud/config"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Users_Route(t *testing.T) {
	testCases := []struct {
		desc         string
		route        string
		method       string
		body         string
		contenttype  string
		expectedCode int
		expectedRes  string
	}{
		{
			desc:         "Create user",
			route:        "/user",
			method:       "POST",
			body:         `{"email":"user1@mail.lan","password":"VERYstrong&Secur3"}`,
			contenttype:  "",
			expectedCode: 201,
			expectedRes:  `{"status":"success","message":"The user has been saved on id: 1","id":1}`,
		},
		/*
			{
				desc:         "",
				route:        "",
				method:       "",
				body:         nil,
				contenttype:  "",
				expectedCode: 0,
				expectedRes: "",
			},
		*/
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			requestBody := bytes.NewBuffer([]byte(tC.body))
			request, _ := http.NewRequest(tC.method, tC.route, requestBody)
			request.Header.Set("Content-Type", "application/json")
			if tC.contenttype != "" {
				request.Header.Set("Content-Type", tC.contenttype)
			}
			response := httptest.NewRecorder()
			apiTest.Router.ServeHTTP(response, request)

			gotBody := response.Body.Bytes()
			gotCode := response.Result().StatusCode
			gotType := response.Header().Get("Content-Type")
			gotJSON := config.Message{}
			json.Unmarshal(gotBody, &gotJSON)
			expectedJSON := config.Message{}
			json.Unmarshal([]byte(tC.expectedRes), &expectedJSON)

			if gotCode != tC.expectedCode {
				t.Errorf("%v fails, got code: %d", tC.desc, gotCode)
			}
			if gotType != "application/json" {
				t.Errorf("%v fails, got content-type: %v", tC.desc, gotType)
			}
			if gotJSON.Status != expectedJSON.Status {
				t.Errorf("%v fails, got status: %v", tC.desc, gotJSON.Status)
			}
			if gotJSON.Message != expectedJSON.Message {
				t.Errorf("%v fails, got message: %v", tC.desc, gotJSON.Message)
			}
			if gotJSON.Id != expectedJSON.Id {
				t.Errorf("%v fails, got id: %d", tC.desc, gotJSON.Id)
			}
		})
	}
}
