package main

import (
	"log"
	"net/http"
)

const tmpl = `<!DOCTYPE html>
<html>
<head>
    <meta name="go-import" content="dbut.sh git https://github.com/dbut2/dbut.sh">
    <meta name="go-source" content="dbut.sh https://github.com/dbut2/dbut.sh https://github.com/dbut2/dbut.sh/tree/main/go{/dir} https://github.com/dbut2/dbut.sh/blob/main/go{/dir}/{file}#L{line}">
</head>
</html>`

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(tmpl))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
