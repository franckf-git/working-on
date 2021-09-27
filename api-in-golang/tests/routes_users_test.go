package tests

import (
	"bytes"
	"encoding/json"
	"lite-api-crud/config"
	"lite-api-crud/controllers"
	"lite-api-crud/models"
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
		{
			desc:         "Bad content type",
			route:        "/user",
			method:       "POST",
			body:         "",
			contenttype:  "jason",
			expectedCode: 406,
			expectedRes:  `{"status":"error","message":"error bad content-type formating:map[Content-Type:[jason]]","id":0}`,
		},
		{
			desc:         "Bad formating",
			route:        "/user",
			method:       "POST",
			body:         `{"name":"user1@mail.lan","password":"VERYstrong&Secur3"}`,
			contenttype:  "",
			expectedCode: 415,
			expectedRes:  `{"status":"error","message":"error while decoding payload <nil>","id":0}`,
		},
		{
			desc:         "Bad email",
			route:        "/user",
			method:       "POST",
			body:         `{"email":"user1@mail?lan","password":"VERYstrong&Secur3"}`,
			contenttype:  "",
			expectedCode: 428,
			expectedRes:  `{"status":"error","message":"error in email or password validator - email must be a valid email and password must be at least 8 characters, uppercase, lowercase, numbers and specials included","id":0}`,
		},
		{
			desc:         "Bad password",
			route:        "/user",
			method:       "POST",
			body:         `{"email":"user1@mail.lan","password":"VERYstrong&Secur"}`,
			contenttype:  "",
			expectedCode: 428,
			expectedRes:  `{"status":"error","message":"error in email or password validator - email must be a valid email and password must be at least 8 characters, uppercase, lowercase, numbers and specials included","id":0}`,
		},
		/*
			{
				desc:         "",
				route:        "",
				method:       "",
				body:         "",
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

func Test_checkEmailPassword(t *testing.T) {
	type args struct {
		user models.User
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Good mail - Good Password",
			args: args{
				user: models.User{
					Email:    "user1@mail.lan",
					Password: "VERYstrong&Secur3",
				},
			},
			want: true,
		},
		{
			name: "Invalid email",
			args: args{
				user: models.User{
					Email:    "user1mail.lan",
					Password: "VERYstrong&Secur3",
				},
			},
			want: false,
		},
		{
			name: "Password - bad length",
			args: args{
				user: models.User{
					Email:    "user1@mail.lan",
					Password: "rR&4567",
				},
			},
			want: false,
		},
		{
			name: "Password - no numbers",
			args: args{
				user: models.User{
					Email:    "user1@mail.lan",
					Password: "rR&aaaab",
				},
			},
			want: false,
		},
		{
			name: "Password - no uppercase",
			args: args{
				user: models.User{
					Email:    "user1@mail.lan",
					Password: "rr&45678",
				},
			},
			want: false,
		},
		{
			name: "Password - no lowercase",
			args: args{
				user: models.User{
					Email:    "user1@mail.lan",
					Password: "RR&45678",
				},
			},
			want: false,
		},
		{
			name: "Password - no special",
			args: args{
				user: models.User{
					Email:    "user1@mail.lan",
					Password: "rRf45678",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := controllers.CheckEmailPassword(tt.args.user); got != tt.want {
				t.Errorf("checkEmailPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_JWT_Route(t *testing.T) {
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
			desc:         "Create JWT",
			route:        "/user/jwt",
			method:       "POST",
			body:         `{"email":"user1@mail.lan","password":"VERYstrong&Secur3"}`,
			contenttype:  "",
			expectedCode: 202,
			expectedRes:  `{"status":"success","message":"Successfull auth, JWT created, it is valid for 24H","id":0}`,
		},
		{
			desc:         "Bad content type",
			route:        "/user/jwt",
			method:       "POST",
			body:         "",
			contenttype:  "jason",
			expectedCode: 406,
			expectedRes:  `{"status":"error","message":"error bad content-type formating:map[Content-Type:[jason]]","id":0}`,
		},
		{
			desc:         "Bad formating",
			route:        "/user/jwt",
			method:       "POST",
			body:         `{"name":"user1@mail.lan","password":"VERYstrong&Secur3"}`,
			contenttype:  "",
			expectedCode: 415,
			expectedRes:  `{"status":"error","message":"error while decoding payload <nil>","id":0}`,
		},
		{
			desc:         "Wrong email",
			route:        "/user/jwt",
			method:       "POST",
			body:         `{"email":"user2@mail.lan","password":"VERYstrong&Secur3"}`,
			contenttype:  "",
			expectedCode: 401,
			expectedRes:  `{"status":"error","message":"This email doesn't exist or the password is wrong","id":0}`,
		},
		{
			desc:         "Wrong password",
			route:        "/user/jwt",
			method:       "POST",
			body:         `{"email":"user1@mail.lan","password":"VERYstrong&Secur4"}`,
			contenttype:  "",
			expectedCode: 401,
			expectedRes:  `{"status":"error","message":"This email doesn't exist or the password is wrong","id":0}`,
		},
		/*
			{
				desc:         "",
				route:        "",
				method:       "",
				body:         "",
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
