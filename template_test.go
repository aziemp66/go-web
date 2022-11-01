package goweb

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func SimpleHTML(w http.ResponseWriter, r *http.Request) {
	templateText := `<html><body>{{.}}</body></html>`
	t, err := template.New("SIMPLE").Parse(templateText)
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(w, "SIMPLE", "Hello Html Template")
}

func TestHtmlTemplate(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:5000", nil)

	SimpleHTML(rec, req)

	res := rec.Result()
	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))
}

func SimpleHtmlFile(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/simple.gohtml")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(w, "simple.gohtml", "Hello Html File Template")
}

//go:embed templates/*.gohtml
var templates embed.FS

func TemplateEmbed(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(w, "simple.gohtml", "Hello HTML Templates")
}

func TestTemplating(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", TemplateEmbed)

	server := http.Server{
		Addr:    "localhost:5000",
		Handler: mux,
	}

	server.ListenAndServe()
}
