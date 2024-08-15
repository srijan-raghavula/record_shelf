package main

import (
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("TODO: write login logic"))
}
