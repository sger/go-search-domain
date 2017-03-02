package main

import (
	"fmt"
	"log"
	"os"
)

var result = map[bool]string{true: "Domain exists", false: "Domain is available"}

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: go-search-domain golang")
		os.Exit(1)
	}

	d := NewDomain(os.Args[1])

	exist, err := d.Exists()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(result[exist])
}
