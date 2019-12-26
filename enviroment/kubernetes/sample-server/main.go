package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/ping", ping)
	http.ListenAndServe("0.0.0.0:5000", nil)
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
	w.WriteHeader(http.StatusOK)
}
