package apiserver

import (
	"fmt"
	"github.com/SerhiiCho/reciper/backend/model"
	"github.com/SerhiiCho/reciper/backend/store/teststore"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServer_authenticateUser(t *testing.T) {
	store := teststore.New()
	user := model.TestUser(t)
	store.User().CreateUser(user)

	testCases := []struct {
		name         string
		cookieValue  map[interface{}]interface{}
		expectedCode int
	}{
		{
			name: "authenticated",
			cookieValue: map[interface{}]interface{}{
				"user_id": user.ID,
			},
			expectedCode: http.StatusOK,
		},
		{
			name:         "not authenticated",
			cookieValue:  nil,
			expectedCode: http.StatusUnauthorized,
		},
	}

	secretKey := []byte("secret")

	serv := newServer(store, sessions.NewCookieStore(secretKey))
	secureCookie := securecookie.New(secretKey, nil)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/", nil)
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			cookieStr, _ := secureCookie.Encode(sessionName, tc.cookieValue)
			req.Header.Set("Cookie", fmt.Sprintf("%s=%s", sessionName, cookieStr))

			serv.authenticateUser(handler).ServeHTTP(rec, req)

			if rec.Code != tc.expectedCode {
				t.Errorf("Expected response code must be %d but %d returned", tc.expectedCode, rec.Code)
			}
		})
	}
}

func TestServer_userCreate(t *testing.T) {
	t.Parallel()

	serv := newServer(teststore.New(), sessions.NewCookieStore([]byte("secret")))

	testCases := []struct {
		name         string
		data         string
		expectedCode int
	}{
		{
			name:         "valid",
			data:         "email=annakototchaeva@mail.ru&password=somevalidpassword",
			expectedCode: http.StatusCreated,
		},
		{
			name:         "no fields",
			data:         "other=some&diff=null",
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name:         "email is invalid",
			data:         "email=somemail@mail&password=12345678910",
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name:         "email is missing",
			data:         "password=12345635435243",
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name:         "password is short",
			data:         "email=somemail@mail.ru&password=1234567",
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name:         "password is missing",
			data:         "email=somemail@mail.ru",
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/api/users", strings.NewReader(tc.data))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			serv.ServeHTTP(rec, req)

			if rec.Code != tc.expectedCode {
				t.Errorf("Expected response code must be %d but %d returned", tc.expectedCode, rec.Code)
			}
		})
	}
}

func TestServer_sessionCreate(t *testing.T) {
	t.Parallel()

	user := model.TestUser(t)
	store := teststore.New()
	serv := newServer(store, sessions.NewCookieStore([]byte("secret")))

	if err := store.User().CreateUser(user); err != nil {
		t.Fatal("can't create user in test store")
	}

	testCases := []struct {
		name         string
		data         string
		expectedCode int
	}{
		{
			name:         "valid",
			data:         fmt.Sprintf("email=%s&password=%s", user.Email, user.Password),
			expectedCode: http.StatusOK,
		},
		{
			name:         "no fields",
			data:         "other=some&diff=null",
			expectedCode: http.StatusUnauthorized,
		},
		{
			name:         "email is invalid",
			data:         "email=somemail@mail&password=12345678910",
			expectedCode: http.StatusUnauthorized,
		},
		{
			name:         "email is missing",
			data:         "password=12345635435243",
			expectedCode: http.StatusUnauthorized,
		},
		{
			name:         "password is short",
			data:         "email=somemail@mail.ru&password=1234567",
			expectedCode: http.StatusUnauthorized,
		},
		{
			name:         "password is missing",
			data:         "email=somemail@mail.ru",
			expectedCode: http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/api/sessions", strings.NewReader(tc.data))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			serv.ServeHTTP(rec, req)

			if rec.Code != tc.expectedCode {
				t.Errorf("Expected response code must be %d but %d returned", tc.expectedCode, rec.Code)
			}
		})
	}
}
