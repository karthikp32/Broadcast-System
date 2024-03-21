package main

import (
	// "encoding/json"
	"encoding/json"
	"log"
	"math"
	"math/rand"
	"slices"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

var broadcastedValues []any
var neighborNodes []string
// Approach 1:
//iterate through neighborNodes
//and send value
func broadcast(value any, neighborNodes []string) {

}

// Approach 1:
//iterate through neighborNodes
//and send value
func read() []any {
	return broadcastedValues
}

// Approach 1:
//set neighborNodes = requestBody["topology"][currentNode]
func topology() {

}


func main() {

	n := maelstrom.NewNode()

	n.Handle("broadcast", func(msg maelstrom.Message) error {
		var requestBody map[string]any

		if err := json.Unmarshal(msg.Body, &requestBody); err != nil {
			return err
		}

		// requestBody["type"] = "generate_ok"
		// requestBody["id"] = generateUniqueId()

		return n.Reply(msg, requestBody)
	})

	n.Handle("read", func(msg maelstrom.Message) error {

	})

	n.Handle("topology", func(msg maelstrom.Message) error {

	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}

}
