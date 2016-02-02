package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tj/go-spin"
)

var marks = map[bool]string{true: "✓", false: "✗"}

func main() {

	s := spin.New()

	for i := 0; i < 30; i++ {
		fmt.Printf("\r  \033[36mcomputing\033[m %s ", s.Next())
		time.Sleep(100 * time.Millisecond)
	}

	d := NewDomain("www.apple.com")

	exist, err := d.Exists()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(exist)
	fmt.Println(marks[exist])
	time.Sleep(1 * time.Second)
}
