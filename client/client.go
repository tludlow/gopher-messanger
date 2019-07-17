package client

import (
	"encoding/json"
	"fmt"
	"gopher-messanger/core"
	"net"
	"sync"
)

//Client - The struct used by a client to connect to a server
type Client struct {
	sync.RWMutex
	host       string
	connection net.Conn
	messages   chan core.Message
	encoder    *json.Encoder
	server     string
}

//Create - Create a new client which connects to a server and can subscribe/publish to messages.
func Create(host string) *Client {
	client := &Client{
		messages: make(chan core.Message),
		host:     host,
	}
	return client
}

//ConnectToServer - Makes the connection to the server we have and intercepts/decodes messages.
func (client *Client) ConnectToServer() error {
	serverConn, err := net.Dial("tcp", client.host)

	if err != nil {
		return fmt.Errorf("[GOPHER] Error when connecting client to the server - %s", client.host)
	}

	client.connection = serverConn
	client.encoder = json.NewEncoder(client.connection)

	errPing := client.pingServer()

	if errPing != nil {
		return fmt.Errorf("[GOPHER] Error when pinging the server from the client on connection")
	}

	return nil
}

//pingServer - Checks that the connection has been established and we can actually ping the server.
func (client *Client) pingServer() error {
	//Ping the server seing if we get a connection.
	return client.encoder.Encode(&core.Message{Command: "ping"})
}

//Close - Close the clients connection to the server.
func (client *Client) Close() {
	//Close the connection to the server and close the client messages channel so that we dont keep around unused data in a channel.
	client.connection.Close()
	close(client.messages)
}
