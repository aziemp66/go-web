package goweb

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
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

func TemplateDataMap(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))

	t.ExecuteTemplate(w, "name.gohtml", map[string]interface{}{
		"Title": "Template Data Struct",
		"Name":  "Azie",
	})
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://locahost:5000", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

type Page struct {
	Title, Name string
}

func TemplateDataStruct(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))

	t.ExecuteTemplate(w, "name.gohtml", Page{
		Title: "Template Data Struct",
		Name:  "Azie Melza Pratama",
	})
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://locahost:5000", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateIf(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))

	t.ExecuteTemplate(w, "if.gohtml", map[string]interface{}{
		"Name": "Azie",
	})
}

func TemplateComparator(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))

	t.ExecuteTemplate(w, "comparator.gohtml", map[string]interface{}{
		"FinalValue": 100,
	})
}

func TemplateRange(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))

	t.ExecuteTemplate(w, "range.gohtml", map[string]interface{}{
		"Hobbies": []string{
			"Gaming", "Ngoding", "Reading",
		},
	})
}

func TemplateWith(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/with.gohtml"))

	t.ExecuteTemplate(w, "with.gohtml", map[string]interface{}{
		"Name": "Azie",
		"Address": map[string]interface{}{
			"Street": "Sukadana",
			"City":   "Kayuagung",
		},
	})
}

func TemplateLayout(w http.ResponseWriter, r *http.Request) {
	t := template.Must(
		template.ParseFiles(
			"./templates/header.gohtml",
			"./templates/footer.gohtml",
			"./templates/layout.gohtml",
		),
	)

	t.ExecuteTemplate(w, "layout", map[string]any{
		"Name":  "Azie",
		"Title": "Template Layout",
	})
}

type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ", My Name Is " + myPage.Name
}

func TemplateFunction(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/function.gohtml"))
	t.ExecuteTemplate(writer, "function.gohtml", MyPage{
		Name: "Eko",
	})
}

func TemplateFunctionGlobal(w http.ResponseWriter, r *http.Request) {
	t := template.New("function")
	t = t.Funcs(map[string]any{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse(`{{upper .Name}}`))

	t.ExecuteTemplate(w, "function", MyPage{
		Name: "Melza",
	})
}

func TemplateFunctionPipelines(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]any{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
		"sayHello": func(value string) string {
			return "Hello " + value
		},
	})

	t = template.Must(t.Parse(`{{sayHello .Name | upper}}`))

	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Melza",
	})

}

//go:embed templates/*.gohtml
var TemplatesCaching embed.FS

var myTemplates = template.Must(template.ParseFS(TemplatesCaching, "templates/*.gohtml"))

func TemplatesCachingFunction(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "simple.gohtml", "Hello HTML Template")
}

func TestTemplateAction(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/if", TemplateIf)
	mux.HandleFunc("/comparator", TemplateComparator)
	mux.HandleFunc("/range", TemplateRange)
	mux.HandleFunc("/with", TemplateWith)
	mux.HandleFunc("/layout", TemplateLayout)
	mux.HandleFunc("/function", TemplateFunction)
	mux.HandleFunc("/functionglobal", TemplateFunctionGlobal)
	mux.HandleFunc("/functionpipelines", TemplateFunctionPipelines)
	mux.HandleFunc("/cache", TemplatesCachingFunction)

	server := http.Server{
		Addr:    "localhost:5000",
		Handler: mux,
	}

	server.ListenAndServe()
}
