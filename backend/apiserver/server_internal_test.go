package apiserver

import (
	"github.com/SerhiiCho/reciper/backend/store/teststore"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServer_RecipeIndex(t *testing.T) {
	t.Parallel()

	serv := newServer(teststore.New())

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
