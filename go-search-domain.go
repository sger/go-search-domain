package main

import (
	"fmt"
	"log"
	"time"
)

var marks = map[bool]string{true: "✓", false: "✗"}

func main() {

	d := NewDomain("www.apple.com")

	exist, err := d.Exists()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(exist)
	fmt.Println(marks[!exist])
	time.Sleep(1 * time.Second)
}
