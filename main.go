package main

import (
	"rest-skeleton/core"
	"rest-skeleton/helloservice"
)

func main() {
	server := core.Server()

	helloservice.SetRouter(server.Router)

	server.Startup()
}
