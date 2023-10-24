package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	response, error := http.Get("http://api.api.svc.cluster.local:8080")
	if error != nil {
	   fmt.Println(error)
	}

	// read response body
	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
	   fmt.Println(error)
	}
	// close response body
	response.Body.Close()
	
	API_response := string(body)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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
		`, API_response)
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Started server on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
