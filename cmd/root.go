package cmd

import (
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func Run() {

	var in string

	// Option Parser
	app := &cli.App{
		Name:        "taro",
		Version:     "1.0.0",
		Compiled:    time.Now(),
		Description: "upload images to various image sites",
		Usage:       "upload images to various image sites",
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:        "in",
				Usage:       "Image to edit (input)",
				Aliases:     []string{"i"},
				Destination: &in,
				Required:    true,
				TakesFile:   true,
			},
		},
		Action: func(c *cli.Context) error {
			name := Add(in)
			UploadToS3(name, in)
			List()
			return nil
		},
	}

	app.Run(os.Args)
}
