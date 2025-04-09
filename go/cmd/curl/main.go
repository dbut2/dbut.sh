package main

import (
	"dbut.sh/pkg/provider"
	"dbut.sh/pkg/server"
	"dbut.sh/pkg/utils"
)

func main() {
	p := utils.Must(provider.NewMarkdownProvider(80))
	err := server.StartHTTPServer(":8080", p)
	if err != nil {
		panic(err)
	}
}
