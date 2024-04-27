package main

import (
	// "encoding/json"
	"encoding/json"
	"log"
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

var broadcastedValues []any
var neighborNodes []string


func broadcast(n *maelstrom.Node) {
	n.Handle("broadcast", func(msg maelstrom.Message) error {
		var requestBody map[string]any

		if err := json.Unmarshal(msg.Body, &requestBody); err != nil {
			return err
		}

		requestBody["type"] = "broadcast_ok"
		broadcastedValues = append(broadcastedValues, requestBody["message"])
		requestBody["message"] = nil

		return n.Reply(msg, requestBody)
	})
}


func read(n *maelstrom.Node) {
	n.Handle("read", func(msg maelstrom.Message) error {
		var requestBody map[string]any


		if err := json.Unmarshal(msg.Body, &requestBody); err != nil {
			return err
		}
		
		requestBody["type"] = "read_ok"
		requestBody["messages"] = broadcastedValues

		return n.Reply(msg, requestBody)
	})
}

func topology(n *maelstrom.Node) {
	n.Handle("topology", func(msg maelstrom.Message) error {
		var requestBody map[string]any
		var topology map[string]string 

		if err := json.Unmarshal(msg.Body, &requestBody); err != nil {
			return err
		}
		
		requestBody["type"] = "topology_ok"
		topology = requestBody["topology"].(map[string]string)
		neighborNodes = append(neighborNodes, topology[n.ID()])
		requestBody["topology"] = nil

		return n.Reply(msg, requestBody)
	})
}


func main() {

	n := maelstrom.NewNode()

	broadcast(n)
	read(n)
	topology(n)

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}

}
