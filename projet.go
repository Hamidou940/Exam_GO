package main

import (
	"net/http"
)

func main() {
	var task []Task

	http.HandleFunc("/", list)
	http.HandleFunc("/done", done)
	http.HandleFunc("/add", add)
}

func list(rw http.ResponseWriter, _ *http.Request) {

}

func done(rw http.ResponseWriter, _ *http.Request) {

}

func add(rw http.ResponseWriter, _ *http.Request) {

}

type Task struct {
	Description string
	Done        bool
}
