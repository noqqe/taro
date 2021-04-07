package cmd

import (
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func Run() {

	// Option Parser
	app := &cli.App{
		Name:        "taro",
		Version:     "1.0.0",
		Compiled:    time.Now(),
		Description: "upload images to various image sites",
		Usage:       "upload images to various image sites",
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a photo to the queue",
				Action: func(c *cli.Context) error {
					name := Add(c.Args().Get(0))
					UploadToS3(name, c.Args().Get(0))
					return nil
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "list all photos",
				Action: func(c *cli.Context) error {
					List()
					return nil
				},
			},
			{
				Name:    "show",
				Aliases: []string{"s"},
				Usage:   "add a photo to the queue",
				Action: func(c *cli.Context) error {
					Show(c.Args().Get(0))
					return nil
				},
			},
			{
				Name:    "upload",
				Aliases: []string{"u"},
				Usage:   "upload a photo to sites",
				Action: func(c *cli.Context) error {
					UploadToFlickr(c.Args().Get(0))
					return nil
				},
			},
		},
	}

	app.Run(os.Args)
}
