package goweb

import (
	"fmt"
	"net/http"
	"testing"
)

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("file")
	if filename == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "BAD REQUEST")
		return
	}

	w.Header().Add("Content-Disposition", "attachment; filename=\""+filename+"\"")
	http.ServeFile(w, r, "./resources/"+filename)
}

func TestDownloadFileServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", DownloadFile)

	server := http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	server.ListenAndServe()
}
