package main

import (
	"fmt"
	"gopher-messanger/client"
	"gopher-messanger/server"
)

func main() {

	//TODO make the server keep running until terminated by the user using ctrl-c

	fmt.Println("[GOPHER] Starting program.")
	srv := server.Create("127.0.0.1:1452")
	err := srv.Start()

	if err != nil {
		panic("ERROR STARTING SERVER IN MAIN")
	}

	clnt := client.Create("127.0.0.1:1452")
	clnt.ConnectToServer()

	for {
	}

}
