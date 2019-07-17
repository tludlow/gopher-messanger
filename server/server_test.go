package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPServerCreation(t *testing.T) {
	//Create just a standard server on the default URI.
	srv1 := Create("127.0.0.1:1452")
	err1 := srv1.Start()

	assert.Nil(t, err1, "Creating a server on the default URI has a problem on start.")
}
