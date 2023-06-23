package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	"go.autokitteh.dev/demand/demand"
)

func main() {
	app := cli.App{
		Name:      "demand",
		Usage:     "check for installed versions of various tools",
		UsageText: "demand [-v] [-q] [spec-path1 [spec-path2 ...]]",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "verbose",
				Aliases:     []string{"v"},
				Destination: &logFlags.verbose,
			},
			&cli.BoolFlag{
				Name:        "quiet",
				Aliases:     []string{"q"},
				Destination: &logFlags.quiet,
			},
		},
		Action: func(c *cli.Context) error {
			if c.Args().Len() == 0 {
				info("no spec files specified")
				return nil
			}

			for _, path := range c.Args().Slice() {
				info("processing %q", path)

				if err := demand.DemandPath(path); err != nil {
					return fmt.Errorf("%s: %w", path, err)
				}
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		hiss("%v", err)
		os.Exit(1)
	}
}
