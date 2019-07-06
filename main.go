package main

import (
	"flag"
	"fmt"

	"github.com/ypapax/btree"
)

func main() {
	var nodes, maxVal int
	flag.IntVar(&nodes, "nodes", 10, "amount of nodes")
	flag.IntVar(&maxVal, "maxval", 10, "maximum value in a node")
	flag.Parse()

	fmt.Println(btree.Random(nodes, maxVal))

}
