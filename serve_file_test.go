package goweb

import (
	"net/http"
	"testing"
)

func ServeFile(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("name") != "" {
		http.ServeFile(w, r, "./serve/ok.html")
	} else {
		http.ServeFile(w, r, "./serve/error.html")
	}
}

func TestServeFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:5000",
		Handler: http.HandlerFunc(ServeFile),
	}

	server.ListenAndServe()
}
