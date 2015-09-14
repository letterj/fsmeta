package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"time"
)

// StreamToString is used as a conversion function
func StreamToString(stream io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.String()
}

// Logger is a function to log requesets
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s\n",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)

		if r.Method == "POST" {
			log.Printf("Body: %s\n", StreamToString(r.Body))
		}

	})
}
