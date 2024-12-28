package main

import "backendcalc/cmd/api"

func main() {
	server := api.NewAPIServer(":8080")
	server.Run()
}
