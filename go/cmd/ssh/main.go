package main

import (
	"dbut.sh/pkg/server"
)

func main() {
	err := server.StartSSHServer(":8022", "host_key")
	if err != nil {
		panic(err)
	}
}
