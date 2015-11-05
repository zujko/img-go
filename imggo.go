package main

import (
	"os"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "img-go"
	app.Usage = "INSERT USAGE HERE"
	app.Action = func(c *cli.Context) {
	}
	app.Run(os.Args)
}
