package main

import (
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	r.UserAgent()
	w.Write([]byte("hello world"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
