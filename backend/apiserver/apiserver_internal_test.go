package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIServer_RecipeIndex(t *testing.T) {
	server := NewAPIServer(NewConfig())

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/recipes", nil)

	server.recipeIndex().ServeHTTP(rec, req)

	if rec.Body.String() != "Hello" {
		t.Error("response must be Hello")
	}
}
