package main

import (
	"dancing-pony/cmd/server"
)

func main() {
	server := server.NewApp()

	server.Start()
}
