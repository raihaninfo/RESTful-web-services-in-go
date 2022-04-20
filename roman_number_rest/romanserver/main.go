package main

import (
	romannumber "RESTful-web-services-in-go/roman_number_rest/roman_number"
	"fmt"
	"html"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlPathElements := strings.Split(r.URL.Path, "/")
		// If request is GET with correct syntax
		if urlPathElements[1] == "roman_number" {
			number, _ := strconv.Atoi(strings.TrimSpace(urlPathElements[2]))
			if number == 0 || number > 10 {
				// If resource is not in the list, send Not Found status
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404 - Not Found"))
			} else {
				fmt.Fprintf(w, "%q",
					html.EscapeString(romannumber.Numberals[number]))
			}
		} else {
			// For all other requests, tell that Client sent a bad request
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad request"))
		}
	})
	
	// Create a server and run it on 8000 port
	s := &http.Server{
		Addr:           ":8002",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
