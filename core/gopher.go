package core

//Each programming running a gopher-messanger has a gopher instance, this handles both the client and server aspects
//of the program. The client handles TCP connections and the server has a proxy which mediates this.

var uid uint32

//Message - The structure being sent between the client and the server through their corresponding proxies.
type Message struct {
	Command string `json:"command"`
	Tag     string `json:"tag,omitempty"`
	Data    string `json:"data,omitempty"`
}
