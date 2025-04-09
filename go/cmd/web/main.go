package main

import (
	"dbut.sh/pkg/provider"
	"dbut.sh/pkg/server"
)

func main() {
	p := &provider.HTMLProvider{}
	var err = server.StartHTTPServer(":8080", p)
	if err != nil {
		panic(err)
	}
}
