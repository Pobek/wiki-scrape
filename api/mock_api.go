package api

import (
	"fmt"

	"github.com/urfave/cli"
)

func Mock(c *cli.Context) {
	fmt.Printf("Args : %q", c.Args())
}
