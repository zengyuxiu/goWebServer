package main

import (
	"FileDownload/client"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var clientCommand = cli.Command{
	Name:  "client",
	Usage: `Run Client`,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "d",
			Usage: "Download file",
		},
		cli.BoolFlag{
			Name:  "l",
			Usage: "List",
		},
	},
	Action: func(ctx cli.Context) error {
		if len(ctx.Args()) < 1 {
			return fmt.Errorf("Missing command")
		}
		list := ctx.Bool("l")
		if list == false {
			Path := ctx.String("d")
			if Path != "" {
				log.Infof("Path : %v", Path)
				client.Download(Path)
			}
		} else {
			client.ListDir()
		}
		return nil
	},
}
var serverCommand = cli.Command{
	Name:  "server",
	Usage: `Run Server`,
	Action: func(ctx cli.Context) error {
		if len(ctx.Args()) < 1 {
			return fmt.Errorf("Missing command")
		}
		var cmdArray []string
		for _, arg := range ctx.Args() {
			cmdArray = append(cmdArray, arg)
		}
		Path := cmdArray[0]
		if Path != "" {
			log.Infof("Path : %v", Path)
		}
		return nil
	},
}