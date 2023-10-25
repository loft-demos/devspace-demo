package main

import (
	"fmt"
	"net/http"
	"flag"
)

func main() {
	text := flag.String("text", "UNSET", "text api response")
	flag.Parse()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "test api response - %s", *text)
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Started server on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
