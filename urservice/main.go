package main

import (
	"log"

	"github.com/BhanuKiranChaluvadi/go_example/tree/main/urservice/loader"
)

func main() {
	err := loader.Load("metadata/example.yaml")
	if err != nil {
		log.Fatal(err)
	}
}
