package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pobek/wiki-scrape/api"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "WikiScrape"
	app.Usage = "Will scrape wikipedia for knowledge"

	app.Commands = []cli.Command{
		{
			Name:   "print",
			Usage:  "Prints 'Hello `Name`'",
			Action: testOutput,
		},
		{
			Name:   "api",
			Usage:  "Trigger API call",
			Action: api.Mock,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func testOutput(c *cli.Context) error {
	name := "Marco"
	if c.NArg() > 0 {
		name = c.Args().Get(0)
	}
	fmt.Printf("Hello %s", name)
	return nil
}
