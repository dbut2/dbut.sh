package main

import (
	"log"
	"net/http"
	"strings"

	"dbut.sh/pages"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix("/", r.URL.Path)
		if path == "" {
			path = "index"
		}
		bytes, err := pages.HTML.ReadFile(path + ".html")
		if err != nil {
			panic(err.Error())
		}
		w.Header().Add("Content-Type", "text/html")
		w.Write(bytes)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
