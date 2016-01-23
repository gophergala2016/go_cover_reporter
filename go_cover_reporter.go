// server1 is a minimal "echo" server.
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type message struct {
	Name string
	Body string
}

const (
	filename = "persist.txt"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler function
	http.HandleFunc("/receiver", receiver)
	http.HandleFunc("/demo_badge", func(w http.ResponseWriter, r *http.Request) {
		coverBadge(w, 0)
	})
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	buffer, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprintf(w, string(buffer))
}

func receiver(rw http.ResponseWriter, req *http.Request) {

	file, err := os.Create(filename)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))

	var t message
	err = json.Unmarshal(body, &t)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = io.WriteString(file, t.Body)
	if err != nil {
		log.Fatalln(err)
	}
	file.Close()

	log.Println(t.Body)
}

func dummyFunction(i int, j int) int {
	return i + j
}
