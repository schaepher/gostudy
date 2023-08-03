//
//  Weather update server.
//  Binds PUB socket to tcp://*:5556
//  Publishes random weather updates
//

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	zmq "github.com/pebbe/zmq4"
	"github.com/vmihailenco/msgpack"
)

func main() {
	//  Socket to talk to server
	fmt.Println("Collecting updates from weather server...")
	subscriber, _ := zmq.NewSocket(zmq.REP)
	defer subscriber.Close()
	subscriber.Bind("tcp://127.0.0.1:7012")
	//  Subscribe to zipcode, default is NYC, 10001
	filter := ""
	if len(os.Args) > 1 {
		filter = os.Args[1] + " "
	}
	subscriber.SetSubscribe(filter)

	for {
		msg, _ := subscriber.Recv(0)
		var data map[string]interface{}
		msgpack.Unmarshal([]byte(msg), &data)
		subscriber.Send(data["file_name"].(string), 0)
		j, _ := json.Marshal(data)
		fmt.Println(string(j))
		time.Sleep(2 * time.Second)
	}
}
