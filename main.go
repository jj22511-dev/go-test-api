package main

import (
	"net/http"
)

func main() {
	// Request Multiplexer
	mux := http.NewServeMux()

	// Register Routes
	mux.Handle("/home", &homeHandler{})

	// Run Server
	http.ListenAndServe(":5000", mux)
}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("This is my home page"))
}
