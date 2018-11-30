package main

import (
	"fmt"
	"log"
)

func main() {
	tree, err := ParseFile("foo.tless")
	if err != nil {
		log.Fatalf("%+v", err)
	}
	fmt.Println(tree.Text())
}
