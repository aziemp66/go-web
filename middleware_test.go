package goweb

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Before Execute Handler")
	middleware.Handler.ServeHTTP(w, r)
	fmt.Println("After Execute Handler")
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Middleware")
	})

	logMiddleware := new(LogMiddleware)
	logMiddleware.Handler = mux

	server := http.Server{
		Addr:    ":3000",
		Handler: logMiddleware,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

type ErrorHandler struct {
	Handler http.Handler
}

func (handler *ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		fmt.Println("RECOVER :", err)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error : %s", err)
		}
	}()

	handler.Handler.ServeHTTP(w, r)
}

func TestErrorMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		panic("Panic Gwe Guys")
	})

	errorHandler := &ErrorHandler{Handler: &LogMiddleware{Handler: mux}}

	server := http.Server{
		Addr:    ":3000",
		Handler: errorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
