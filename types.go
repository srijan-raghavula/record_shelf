package main

import (
	"fmt"
	"net/http"
)

type serverConfig struct {
	visits int
}

func (sCfg *serverConfig) getVisits(w http.ResponseWriter, r *http.Request) {
	visitCount := sCfg.visits
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`
<html>
<body>
    <h1>Welcome to the Record Shelf</h1>
    <p>The Shelf has been visited %d times since the server started</p>
</body>
</html>
    `, visitCount)))
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
}
