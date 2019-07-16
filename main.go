package main

import (
	"fmt"
	"gopher-messanger/server"
)

func main() {
	fmt.Println("[GOPHER] Starting program.")
	srv := server.Create("127.0.0.1:1452")
	err := srv.Start()

	if err != nil {
		panic("ERROR STARTING SERVER IN MAIN")
	}
}
