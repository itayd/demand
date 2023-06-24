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
		Action: func(c *cli.Context) error {
			if c.Args().Len() == 0 {
				return nil
			}

			ok := true

			for _, path := range c.Args().Slice() {
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

				fmt.Println(string(s))
			}

			if !ok {
				return cli.Exit("", 9)
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
