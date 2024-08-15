package main

import (
	"log"
	"net/http"
)

const rootDir = "."

var shelf = &http.Server{
	Addr: ":6969",
}

var sCfg = serverConfig{}

func main() {
	dir := http.Dir(rootDir)
	fs := http.FileServer(dir)

	mux := http.NewServeMux()
	mux.Handle("/", sCfg.trafficCounter(fs))

	shelf.Handler = mux

	mux.HandleFunc("GET /shelf/traffic", sCfg.getVisits)

	// USER PROFILE RELATED ENDPOINTS
	mux.HandleFunc("POST /login", login)
	mux.HandleFunc("POST /add/users", createUser)
	mux.HandleFunc("PUT /update/users/{userID}", updateUser)
	mux.HandleFunc("DELETE /remove/users/{userID}", removeUser)

	// FILES RELATED ENDPOINTS
	mux.HandleFunc("POST /files/users/{userID}", addFile)
	mux.HandleFunc("GET /files/users/{userID}", getFile)
	mux.HandleFunc("PUT /files/users/{userID}", replaceFile)
	mux.HandleFunc("DELETE /files/users/{userID}", removeFile)

	log.Println("serving from", rootDir, "at localhost", shelf.Addr)
	log.Fatal(shelf.ListenAndServe())
}
