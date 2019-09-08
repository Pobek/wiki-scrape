package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "WikiScrape"
	app.Usage = "Will scrape wikipedia for knowledge"
	app.Action = func(c *cli.Context) error {
		fmt.Println("WikiScrape test#1")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
