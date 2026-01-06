package main

import (
	"context"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Name:  "gendiff",
		Usage: "Compares two configuration files and shows a difference.",
	}

	app.Run(context.Background(), os.Args)
}
