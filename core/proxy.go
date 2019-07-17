package core

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//Proxy - An intermediary between a client and a server which communicates the messages being sent, e.g
//Client->proxy---->proxy->server
type Proxy struct {
	sync.RWMutex
	Pipe           chan Message
	check          chan Message
	done           chan bool
	id             uint32
	subscribedTags []string
}

//CreateProxy - Create a new proxy which is the inbetween the client/server and can handle messaging across networks.
func CreateProxy() (p *Proxy) {

	// create new proxy
	p = &Proxy{
		Pipe:           make(chan Message),
		check:          make(chan Message),
		done:           make(chan bool),
		id:             atomic.AddUint32(&uid, 1),
		subscribedTags: make([]string, 100), //Only has 100 max size, might need to change in future
	}

	go p.handleProxiedMessages()

	return
}

func (p *Proxy) handleProxiedMessages() {
	defer func() {
		fmt.Println("[GOPHER] Received true through the proxy done channel, ending proxy operations.")
		close(p.Pipe)
		close(p.check)
	}()

	for {
		select {
		case message := <-p.check:
			//We have received a message from the server, we need to filter the message to make sure it has a tag we care about (subscribed too)
			fmt.Print("[GOPHER] Proxy check channel triggered.")
			p.RLock()
			doCare := false
			for _, tag := range p.subscribedTags {
				if tag == message.Tag {
					doCare = true
					break
				}
			}
			p.RUnlock()

			//We care about the message, send it down our main pipe which we use.
			if doCare {
				p.Pipe <- message
			}
		case <-p.done:
			//Something has been sent down the done channel, the proxy is over, end it.
			return
		}
	}
}

//Subscribe - Subscribe this proxy to a tag so that when we receive a publish with this tag we can handle it.
func (p *Proxy) Subscribe(tag string) error {
	if len(tag) == 0 {
		return fmt.Errorf("Tag length of 0, not allowed to subscribe to this")
	}

	p.Lock()
	for _, ele := range p.subscribedTags {
		if ele == tag {
			return nil
		}
	}
	p.subscribedTags = append(p.subscribedTags, tag)
	p.Unlock()
	return nil
}

//Close - Close the channels so we dont have unused allocated memory lying around (should be fine if we dont as they will go when the client is no longer used)
func (p *Proxy) Close() {
	close(p.done)
	close(p.check)
	close(p.Pipe)
}
