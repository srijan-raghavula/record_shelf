package main

import (
	"net/http"
)

func (sCfg *serverConfig) trafficCounter(base http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sCfg.visits++
		base.ServeHTTP(w, r)
	})
}
