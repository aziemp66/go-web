package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(w http.ResponseWriter, r http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	fName := r.PostForm.Get("firstname")
	lName := r.PostForm.Get("lastname")

	fmt.Fprintf(w, "Hello %s %s", fName, lName)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("firstname=Azie&lastname=Melza")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPost(recorder, *request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
