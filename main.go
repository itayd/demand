package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/itayd/demand/demand"
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

			ok := true

			for _, path := range c.Args().Slice() {
				debug("processing %q", path)

				r, err := demand.DemandPath(path)
				if err != nil {
					return fmt.Errorf("%s: %w", path, err)
				}

				if !r.OK {
					ok = false
				}

				s, err := json.MarshalIndent(r, "", "  ")
				if err != nil {
					return fmt.Errorf("json marshal: %w", err)
				}

				print("%s", s)
			}

			if !ok {
				return cli.Exit("", 9)
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		hiss("%v", err)
		os.Exit(1)
	}
}
