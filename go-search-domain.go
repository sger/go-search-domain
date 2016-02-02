package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/brettweavnet/gosync/version"
	"github.com/codegangsta/cli"
)

var marks = map[bool]string{true: "Domain exists", false: "Domain is available"}

func main() {

	app := cli.NewApp()
	app.Version = version.Version()
	app.Name = "go-search-domain"
	app.Usage = "go-search-domain DOMAIN"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "domain",
			Value: "",
			Usage: "",
		},
	}

	app.Action = func(c *cli.Context) {
		domain := c.String("domain")
		fmt.Println(domain)
	}

	app.Run(os.Args)

	d := NewDomain("www.apple.com")

	exist, err := d.Exists()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(marks[exist])
	time.Sleep(1 * time.Second)
}
