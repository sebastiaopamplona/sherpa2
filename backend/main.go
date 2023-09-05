package main

import (
	"io"
	"net/http"
	"time"
)

func ListSprints(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second)
	//w.WriteHeader(http.StatusBadRequest)
	io.WriteString(w, `[{"id": 1, "title": "Sprint "}, {"id": 2, "title": "Sprint 2"}, {"id": 3, "title": "Sprint 3"}]`)
}

func main() {
	http.HandleFunc("/sprints", ListSprints)

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		panic(err)
	}
}
