package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/codegangsta/cli"
	"github.com/sger/go-search-domain/version"
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
			Usage: "Search for a domain like www.apple.com",
		},
	}

	app.Action = func(c *cli.Context) {
		err := validateArgs(c)

		if err != nil {
			fmt.Println(err)
		} else {
			domain := c.String("domain")
			fmt.Println(domain)
			d := NewDomain(domain)

			exist, err := d.Exists()

			if err != nil {
				log.Fatalln(err)
			}

			fmt.Println(marks[exist])
			time.Sleep(1 * time.Second)
		}
	}

	app.Run(os.Args)
}

func validateArgs(c *cli.Context) error {
	if len(c.Args()) != 1 {
		return fmt.Errorf("domain value is required")
	}
	return nil
}
