package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/maxkueng/go-streamable"
)

func main() {
	app := cli.NewApp()

	app.Name = "streamable-cli"
	app.Usage = "upload videos to streamable.com"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "username, u",
			Usage: "streamable.com username",
		},
	}

	app.Action = func(c *cli.Context) {
		println("Greetings " + c.String("username"))
		res, _ := streamable.Upload("/home/max/Downloads/feedremove2.mp4")
		fmt.Printf("%+v\n", res)
	}

	app.Run(os.Args)
}
