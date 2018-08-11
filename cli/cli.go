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
	app.Name = "Raspberry Pi GPIO Api Server"
	app.Usage = "Use HTTP Restful requests to manage your Raspberry Pi GPIO pins."

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
	fmt.Println("This is default.")
	return nil
	}

	err := app.Run(os.Args)
	if err != nil {
	log.Fatal(err)
	}
}
