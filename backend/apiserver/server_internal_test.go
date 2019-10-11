package apiserver

import (
	"fmt"
	"github.com/SerhiiCho/reciper/backend/store/teststore"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServer_RecipeIndex(t *testing.T) {
	serv := newServer(teststore.New())

	testCases := []struct {
		name         string
		data         map[string]string
		expectedCode int
	}{
		{"valid", map[string]string{
			"email":    "annakototchaeva@mail.ru",
			"password": "somevalidpassword",
		}, http.StatusCreated},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			dataString := fmt.Sprintf("email=%s&password=%s", tc.data["email"], tc.data["password"])

			req, _ := http.NewRequest(http.MethodPost, "/api/users", strings.NewReader(dataString))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			serv.ServeHTTP(rec, req)

			if rec.Code != tc.expectedCode {
				t.Errorf("Expected response code must be %d but %d returned", tc.expectedCode, rec.Code)
			}
		})
	}
}
