package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

var marks = map[bool]string{true: "Domain exists", false: "Domain is available"}

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: go-search-domain www.golang.com")
		os.Exit(1)
	}

	fmt.Println(os.Args[1])

	d := NewDomain(os.Args[1])

	exist, err := d.Exists()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(marks[exist])
	time.Sleep(1 * time.Second)

}
