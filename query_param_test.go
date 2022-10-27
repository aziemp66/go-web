package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		fmt.Fprint(w, "Hello")
	} else {
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:5000/hello?name=azie", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func SayHelloMultipleParameter(w http.ResponseWriter, r *http.Request) {
	first_name := r.URL.Query().Get("first_name")
	last_name := r.URL.Query().Get("last_name")

	if first_name == "" && last_name == "" {
		fmt.Fprint(w, "Hello")
	} else if first_name == "" {
		fmt.Fprintf(w, "Hello %s", last_name)
	} else if last_name == "" {
		fmt.Fprintf(w, "Hello %s", first_name)
	} else {
		fmt.Fprintf(w, "Hello %s %s", first_name, last_name)
	}
}

func TestMultipleQuery(t *testing.T) {
	request := httptest.NewRequest(
		http.MethodGet,
		"http://localhost:5000/hello?first_name=azie&last_name=melza",
		nil,
	)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
