//
//  Weather update client.
//  Connects SUB socket to tcp://localhost:5556
//  Collects weather updates and finds avg temp in zipcode
//

package main

import (
	"math/rand"
	"time"

	zmq "github.com/pebbe/zmq4"

	"fmt"
)

func main() {

	//  Prepare our publisher
	publisher, _ := zmq.NewSocket(zmq.PUB)
	defer publisher.Close()
	publisher.Connect("tcp://localhost:5556")

	//  Initialize random number generator
	rand.Seed(time.Now().UnixNano())

	// loop for a while apparently
	for {

		//  Get values that will fool the boss
		// zipcode := rand.Intn(100000)
		// zipcode := ""
		temperature := rand.Intn(215) - 80
		relhumidity := rand.Intn(50) + 10

		//  Send message to all subscribers
		msg := fmt.Sprintf("%d %d",  temperature, relhumidity)
		publisher.Send(msg, 0)
	}
}
