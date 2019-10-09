package apiserver

import (
	"github.com/SerhiiCho/reciper/backend/app"
	"github.com/SerhiiCho/reciper/backend/store"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIServer_RecipeIndex(t *testing.T) {
	application := app.NewApp(store.NewStore(store.NewConfig()))
	server := NewAPIServer(application, NewConfig())

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/recipes", nil)

	server.recipeIndex().ServeHTTP(rec, req)

	if rec.Body.String() != "Hello" {
		t.Error("response must be Hello")
	}
}
