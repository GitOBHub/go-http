package http_test

import (
	"fmt"
	"log"
	"testing"

	"web/http"
)

type Hello int

func TestListenAndServe(t *testing.T) {
	var hello Hello
	log.Fatal(http.ListenAndServe(":5000", hello))
}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func TestHandleFunc(t *testing.T) {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func TestHandlerFunc(t *testing.T) {
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(sayHello)))
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}
