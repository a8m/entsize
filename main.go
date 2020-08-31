package main

import (
	"fmt"
	"log"
	"os"

	"github.com/facebook/ent/entc"
	"github.com/facebook/ent/entc/gen"
)

func main() {
	path := "./ent/schema"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	graph, err := entc.LoadGraph(path, &gen.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	var fields, edges int
	for _, n := range graph.Nodes {
		fields += len(n.Fields)
		edges += len(n.Edges)
	}
	fmt.Printf("Types: %d, Fields: %d, Edges: %d\n", len(graph.Nodes), fields, edges)
}
