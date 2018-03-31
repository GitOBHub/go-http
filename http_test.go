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
	fmt.Fprintf(w, "fuck you")
}
