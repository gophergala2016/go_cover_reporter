// server1 is a minimal "echo" server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler function
	http.HandleFunc("/demo_badge", func(w http.ResponseWriter, r *http.Request) {
		cover_badge(w, 0)
	})
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func dummy_function(i int, j int) int {
	return i + j
}
