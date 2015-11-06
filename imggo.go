package main

import (
	"os"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/fatih/color"
)

func main() {
	app := cli.NewApp()
	app.Author = "zujko"
	app.Name = "img-go"
	app.Usage = "INSERT USAGE HERE"
	app.Commands = []cli.Command{
		{
			Name: "resize",
			Aliases: []string{"r"},
			Usage: "resizes the image",
			Action: func(c *cli.Context) {
				if(len(c.Args()) >= 2) {
					//RESIZE IMAGE
				} else {
					fmt.Println("usage: <image name> <size>")
				}
			},

		},
		{
			Name: "blur",
			Aliases: []string{"bl"},
			Usage: "Adds a blur to the image",
			Action: func(c *cli.Context) {
				//BLUR IMAGE
			},
		},
		{
			Name: "brighten",
			Aliases: []string{"br"},
			Usage: "Brighten or darken the image",
			Action: func(c *cli.Context) {
				//BRIGHTEN IMAGE 
			},
		},

	}
	app.Run(os.Args)
}
