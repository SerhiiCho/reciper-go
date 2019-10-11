package apiserver

import (
	"github.com/SerhiiCho/reciper/backend/store/teststore"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_RecipeIndex(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/recipes", nil)

	newServer(teststore.New()).recipeIndex().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Error("response must be Hello")
	}
}
