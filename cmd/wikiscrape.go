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

	app.Commands = []cli.Command{
		{
			Name:  "generate",
			Usage: "Generates document with scraped data",
			Subcommands: []cli.Command{
				{
					Name:  "simple",
					Usage: "Generates a plain text document",
					Action: func(c *cli.Context) error {
						fmt.Println("Generating plain text document")
						return nil
					},
				},
				{
					Name:  "docx",
					Usage: "Generates a Word document",
					Action: func(c *cli.Context) error {
						fmt.Println("Generating Word document")
						return nil
					},
				},
				{
					Name:  "pdf",
					Usage: "Generates a PDF document",
					Action: func(c *cli.Context) error {
						fmt.Println("Generating PDF document")
						return nil
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
