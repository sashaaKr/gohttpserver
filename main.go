package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// comment
func main() {
  http.Handle("/", loggingMiddlware(http.HandlerFunc(handler)))
  http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello World")
}

func loggingMiddlware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.Infof("uri: %s", r.RequestURI)
    next.ServeHTTP(w, r)
  })
}