package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/charmbracelet/glamour"

	"dbut.sh/pages"
)

func main() {
	rend, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(80),
		glamour.WithEmoji(),
	)
	if err != nil {
		panic(err.Error())
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix("/", r.URL.Path)
		if path == "" {
			path = "index"
		}
		bytes, err := pages.MD.ReadFile(path + ".md")
		if err != nil {
			panic(err.Error())
		}
		bytes, err = rend.RenderBytes(bytes)
		if err != nil {
			panic(err.Error())
		}
		w.Write(bytes)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
