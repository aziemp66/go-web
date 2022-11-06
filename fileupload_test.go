package goweb

import (
	"io"
	"net/http"
	"os"
	"testing"
)

func UploadForm(w http.ResponseWriter, r *http.Request) {
	err := myTemplates.ExecuteTemplate(w, "upload.form.gohtml", nil)
	if err != nil {
		panic(err)
	}
}

func Upload(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}
	name := r.PostFormValue("name")
	myTemplates.ExecuteTemplate(w, "upload.success.gohtml", map[string]string{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}

func TestUploadServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
