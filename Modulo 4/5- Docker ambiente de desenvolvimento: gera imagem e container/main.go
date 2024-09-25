package main

import "net/http"

// To run this application, use the following command:
// go run main.go
// Then, open your web browser and visit http://localhost:8080 OR IN THE TERMINAL USE curl http://localhost:8080
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	http.ListenAndServe(":8080", nil)
}
