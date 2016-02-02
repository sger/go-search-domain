package main

import (
	"fmt"
	"log"
	"time"

	"github.com/codegangsta/cli"
)

var marks = map[bool]string{true: "Domain exists", false: "Domain is available"}

func main() {

	d := NewDomain("www.apple.com")

	exist, err := d.Exists()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(marks[exist])
	time.Sleep(1 * time.Second)

}

func validateArgs(c *cli.Context) error {
	fmt.Println(c.Args())
	if len(c.Args()) != 1 {
		return fmt.Errorf("domain value is required")
	}
	return nil
}
