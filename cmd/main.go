package main

import "10kRps/cmd/server"

func main() {
	serv := server.CreateServer()
	serv.Start()
}
