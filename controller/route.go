package controller

import (
	"net/http"
)

// ping is a handler function that responds with "pong" when accessed

func Register() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ping())
	mux.HandleFunc("/", crud())
	return mux

}
