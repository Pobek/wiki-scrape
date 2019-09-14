package api

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/urfave/cli"
)

type Wiki struct {
	Topic   string
	URL     string
	Content string
}

func ParseWiki(c *cli.Context) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	response, err := client.Get("https://en.wikipedia.org/wiki/Go_(programming_language)")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Copy data from the response to standard output
	n, err := io.Copy(os.Stdout, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Number of bytes copied to STDOUT:", n)
}
