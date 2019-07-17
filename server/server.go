package server

import (
	"encoding/json"
	"fmt"
	"gopher-messanger/core"
	"io"
	"net"
	"sync"
)

//At first the only protocol supported will be HTTP messanging, which will be expanded in the future to more protocols.

type Server struct {
	sync.RWMutex
	uri string
}

//(default uri 127.0.0.1:1452)
func Create(uri string) *Server {
	server := &Server{
		uri: uri,
	}

	return server
}

func (s *Server) Start() error {
	listen, err := net.Listen("tcp", s.uri)

	if err != nil {
		return fmt.Errorf("Problem occured whilst starting the TCP server - %s", err.Error())
	}

	fmt.Printf("[GOPHER] TCP server listening on: %s", s.uri)

	go func() {
		for {
			connection, err := listen.Accept()

			if err != nil {
				panic("[GOPHER] Error occured on the server when receiving TCP connection from client")
			}

			go handleTCPConnection(connection)
		}
	}()

	return nil
}

func handleTCPConnection(connection net.Conn) {

	defer connection.Close()

	//Create JSON encodere/decoder to translate the connection message to something the server can understand.
	//encoder := json.NewEncoder(connection)
	decoder := json.NewDecoder(connection)

	proxy := core.CreateProxy()
	defer proxy.Close()

	//Handle the message incoming from the client
	for {
		message := core.Message{}

		//Decode the JSON message to a message object in the form of the struct called Message within core
		if err := decoder.Decode(&message); err != nil {
			switch err {
			case io.EOF:
				panic("[GOPHER] Client disconnected")
				return
			case io.ErrUnexpectedEOF:
				panic("[GOPHER] Client disconnected (unexpected)")
				return
			default:
				panic("[GOPHER] Failed to decode TCP connection message from client")
				return
			}

		}

		fmt.Println(message.Command)

		//We now need to see what message was sent.
		switch message.Command {
		case "publish":
			//handle publish
			fmt.Println("[GOPHER] Publish command received on server")
		case "subscribe":
			//handle publish
			fmt.Println("[GOPHER] Subscribe command received on server")
		case "ping":
			//handle ping
			fmt.Println("[GOPHER] Received ping from client (IN SERVER SWITCH)")
		default:
			panic("Unknown command issued to the gopher server.")
		}
	}
}
