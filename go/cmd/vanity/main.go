package main

import (
	"dbut.sh/pkg/provider"
	"dbut.sh/pkg/server"
)

const tmpl = `<!DOCTYPE html>
<html>
<head>
    <meta name="go-import" content="dbut.sh git https://github.com/dbut2/dbut.sh">
    <meta name="go-source" content="dbut.sh https://github.com/dbut2/dbut.sh https://github.com/dbut2/dbut.sh/tree/main/go{/dir} https://github.com/dbut2/dbut.sh/blob/main/go{/dir}/{file}#L{line}">
</head>
</html>`

func main() {
	p := provider.NewVanityProvider(tmpl)
	err := server.StartHTTPServer(":8080", p)
	if err != nil {
		panic(err)
	}
}
