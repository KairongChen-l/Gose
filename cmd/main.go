package main

import "github.com/KairongChen-l/Gose/cmd/api"

func main() {
	server := api.NewAPIServer(":8080",nil)
	server.Run()
}
