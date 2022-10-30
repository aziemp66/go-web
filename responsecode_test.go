package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(400)
	} else {
		w.WriteHeader(200)
		fmt.Fprintf(w, "Hi %s", name)
	}
}

func TestResponseCode(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:5000", nil)
	rec := httptest.NewRecorder()

	ResponseCode(rec, req)

	response := rec.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}
