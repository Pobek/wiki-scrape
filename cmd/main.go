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

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name: "verbose",
		},
	}

	app.Action = testOutput

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
	if c.Bool("verbose") {
		fmt.Println("Verbose enabled!")
		fmt.Println("Will output hello <name>")
	}
	fmt.Printf("Hello %s", name)
	return nil
}
