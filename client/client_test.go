package client

import (
	"gopher-messanger/server"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPClientPIng(t *testing.T) {
	//Create a server for the client to ping
	srv := server.Create("127.0.0.1:1452")
	srvError := srv.Start()

	assert.Nil(t, srvError, "Error creating server in client test ping")

	clnt := Create("127.0.0.1:1452")
	clientConnectionError := clnt.ConnectToServer()

	assert.Nil(t, clientConnectionError, "Error pinging server in client connection")

	pingError := clnt.pingServer()

	assert.Nil(t, pingError, "Error pinging server in client test ping")

	clnt.Close()
}
