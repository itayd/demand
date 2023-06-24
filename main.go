package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/itayd/demand/demand"
)

func main() {
	var list, quiet, fail, onlyincompat bool

	app := cli.App{
		Name:      "demand",
		Usage:     "check for installed versions of various tools",
		UsageText: "demand [-v] [-q] [spec-path1 [spec-path2 ...]]",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "fail",
				Aliases:     []string{"f"},
				Destination: &fail,
				Usage:       "fail if any incompatibility detected",
			},
			&cli.BoolFlag{
				Name:        "quiet",
				Aliases:     []string{"q"},
				Destination: &quiet,
				Usage:       "do not print detailed results",
			},
			&cli.BoolFlag{
				Name:        "list",
				Aliases:     []string{"l"},
				Destination: &list,
				Usage:       "list executables that are incompatible",
			},
			&cli.BoolFlag{
				Name:        "only-incompat",
				Aliases:     []string{"o"},
				Destination: &onlyincompat,
				Usage:       "show detailed result only for incompatabilities",
			},
		},
		Action: func(c *cli.Context) error {
			if onlyincompat && quiet {
				return errors.New("--only-incompat and --quiet are mutually exclusive")
			}

			if c.Args().Len() == 0 {
				return nil
			}

			var incompats []string

			for _, path := range c.Args().Slice() {
				r, err := demand.DemandPath(path)
				if err != nil {
					return fmt.Errorf("%s: %w", path, err)
				}

				if !r.OK {
					incompats = append(incompats, r.Executable)
				}

				if !quiet && (!onlyincompat || !r.OK) {
					enc := json.NewEncoder(os.Stdout)
					enc.SetEscapeHTML(false)
					enc.SetIndent("", "  ")

					if err := enc.Encode(r); err != nil {
						return fmt.Errorf("json encode: %w", err)
					}
				}
			}

			if list {
				fmt.Println(strings.Join(incompats, " "))
			}

			if fail && len(incompats) > 0 {
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
