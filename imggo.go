package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
	"github.com/codegangsta/cli"
	"github.com/fatih/color"
	"github.com/disintegration/imaging"
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
			Usage: "Resizes the image",
			Action: func(c *cli.Context) {
				if(len(c.Args()) >= 2) {
					img, err := imaging.Open(c.Args()[0])
					if err != nil {
						color.Red("Error opening image")
					} else {
						size := strings.Split(c.Args()[1],"x")
						width,we := strconv.Atoi(size[0])
						height,he := strconv.Atoi(size[1])
						if (we != nil || he != nil) {
							color.Red("Invalid dimension arguments")
							return
						}
						if(width > 1000 || height > 1000) {
							color.Red("Dimensions are too large")
							return
						}
						resized := imaging.Thumbnail(img, width, height, imaging.CatmullRom)
						resizeName := "resized.jpg"
						if len(c.Args()) == 3 {
							resizeName = c.Args()[2]
						}
						err := imaging.Save(resized,resizeName)
						if err != nil {
							color.Red("Error saving image")
						} else {
							color.Green("Image resized")
						}
					}
				} else {
					fmt.Println("usage: image_name size [output_name]")
				}
			},

		},
		{
			Name: "blur",
			Aliases: []string{"bl"},
			Usage: "Adds a blur to the image",
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fmt.Println("usage: image_name blur_amount [output_name]")
				} else {
					img, err := imaging.Open(c.Args()[0])
					if err != nil {
						color.Red("Error opening Image")
						return
					}

					blurAmt,er := strconv.ParseFloat(c.Args()[1], 64)
					if er != nil {
						color.Red("Invalid arguments")
						return
					}
					dstImage := imaging.Blur(img, blurAmt)
					imgName := "blur.jpg"
					if len(c.Args()) == 3 {
						imgName = c.Args()[2]
					}
					saveErr := imaging.Save(dstImage, imgName)
					if saveErr != nil {
						color.Red("Error saving image")
					} else {
						color.Green("Image blurred")
					}
				}
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
