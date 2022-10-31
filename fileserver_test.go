package goweb

import (
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {
	dir := http.Dir("./resources/")
	fileserver := http.FileServer(dir)

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	server := http.Server{
		Addr:    "localhost:5000",
		Handler: mux,
	}

	server.ListenAndServe()
}
