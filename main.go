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

	log.Println("serving from", rootDir, "at localhost", shelf.Addr)
	log.Fatal(shelf.ListenAndServe())
}
