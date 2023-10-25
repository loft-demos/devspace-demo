package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response, err := http.Get("http://api.api")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer response.Body.Close()

		// read response body
		responseBody, err := io.ReadAll(response.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, `
		<html>
			<head>
				<title>auth</title>
				<link rel="stylesheet" href="https://devspace.sh/css/quickstart.css">
			</head>
			<body>
				<section>
					<div class="container">
						<div class="left">
							<h1>api response: %s</h1>
						</div>
						<div><img src="https://static.loft.sh/devspace/quickstarts/devspace-astronaut.gif" /></div>
					</div>
				</section>
			</body>
		</html>
		`, string(responseBody))
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Started server on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
