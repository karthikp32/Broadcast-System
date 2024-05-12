package main

import (
	"log"
	"encoding/json"
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
	"slices"
	"strings"
	"fmt"
)

var broadcastedValues []any
var neighborNodes []string


func broadcast(n *maelstrom.Node) {
	n.Handle("broadcast", func(msg maelstrom.Message) error {
		var requestBody map[string]any
		responseBody := make(map[string]any)
		if err := json.Unmarshal(msg.Body, &requestBody); err != nil {
			return err
		}
		responseBody["type"] = "broadcast_ok"
		broadcastedValues = append(broadcastedValues, requestBody["message"])
		return n.Reply(msg, responseBody)
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

func removeCurrNodeFromNeighborNodes(neighborNodes []string, n *maelstrom.Node) []string {
	for idx, neighborNode := range neighborNodes {
		if strings.Compare(neighborNode, n.ID()) == 0 {
			neighborNodes[idx] = neighborNodes[len(neighborNodes)-1] 
			neighborNodes = neighborNodes[:len(neighborNodes)-1]
		} 
	}
	return neighborNodes
}

func topology(n *maelstrom.Node) {
	n.Handle("topology", func(msg maelstrom.Message) error {
		requestBody := make(map[string]any)
		neighborNodes = n.NodeIDs()
		neighborNodes = removeCurrNodeFromNeighborNodes(neighborNodes, n)
		requestBody["type"] = "topology_ok"
		return n.Reply(msg, requestBody)
	})
}


func removeCurrNodeFromNeighborNodes_unitTest(n *maelstrom.Node) {
	n.Init("n0", []string{"n1", "n2"})
	neighborNodes = removeCurrNodeFromNeighborNodes(neighborNodes, n)
	if slices.Contains(neighborNodes, n.ID()) {
		fmt.Println("removeCurrNodeFromNeighborNodes_unitTest failed")
	}
}

func main() {

	neighborNodes = []string{"n0", "n1", "n2"}

	n := maelstrom.NewNode()
	removeCurrNodeFromNeighborNodes_unitTest(n)

	broadcast(n)
	read(n)
	topology(n)

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}

}
