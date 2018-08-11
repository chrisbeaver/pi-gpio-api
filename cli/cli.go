package cli

import (
	"os"
	"fmt"
	"log"
	"github.com/urfave/cli"
	"github.com/chrisbeaver/pi-gpio-api/server"
)

func Parse() {
	// Flag
	var port string
	app := cli.NewApp()
	app.Name = "boom"
	app.Usage = "make an explosive entrance"

	// app.Flags = []cli.Flag {
	// 	cli.StringFlag{
	// 	  Name: "port, p",
	// 	  Value: "8000",
	// 	  Usage: "Set the port for the server to listen on.",
	// 	  Destination: &port,
	// 	},
	// }

	app.Commands = []cli.Command{
		{
		    Name:    "listen",
		    Aliases: []string{"l"},
		    Usage:   "Boot up the GPIO pin server.",
		    Action:  func(c *cli.Context) error {
				fmt.Println("Server listening on port: ", port)
				server.Listen(port)
				return nil
		    },
		    Flags: []cli.Flag {
			    cli.StringFlag{
					Name: "port, p",
					Value: "8000",
					Usage: "Set the port for the server to listen on.",
					Destination: &port,
				},
			},
		},
	}
	app.Action = func(c *cli.Context) error {
	fmt.Println(port)
	return nil
	}

	err := app.Run(os.Args)
	if err != nil {
	log.Fatal(err)
	}
}
