package cli

import (
	"os"
	"fmt"
	"log"
	"github.com/urfave/cli"
	"github.com/chrisbeaver/pi-gpio-api/server"
)

func Parse() {
	app := cli.NewApp()
	app.Name = "boom"
	app.Usage = "make an explosive entrance"

	app.Commands = []cli.Command{
		{
		  Name:    "listen",
		  Aliases: []string{"l"},
		  Usage:   "Boot up the GPIO pin server.",
		  Action:  func(c *cli.Context) error {
			// fmt.Println("added task: ", c.Args().First())
			server.Listen()
			return nil
		  },
		},
	}
	app.Action = func(c *cli.Context) error {
	fmt.Println("boom! I say!")
	return nil
	}

	err := app.Run(os.Args)
	if err != nil {
	log.Fatal(err)
	}
}
