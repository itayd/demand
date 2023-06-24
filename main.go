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
	var (
		list, quiet, fail, onlyincompat bool
		specs                           cli.StringSlice
	)

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
			&cli.StringSliceFlag{
				Name:        "spec",
				Aliases:     []string{"s"},
				Destination: &specs,
				Usage:       "one liner spec",
			},
		},
		Action: func(c *cli.Context) error {
			if onlyincompat && quiet {
				return errors.New("--only-incompat and --quiet are mutually exclusive")
			}

			if c.Args().Len()+len(specs.Value()) == 0 {
				return nil
			}

			var incompats []string

			emit := func(r *demand.Result) error {
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

				return nil
			}

			for _, txt := range specs.Value() {
				var spec *demand.Spec
				if txt[0] == '{' {
					var err error
					if spec, err = demand.ParseJSONSpec([]byte(txt)); err != nil {
						return fmt.Errorf("json one liner: %w", err)
					}
				} else {
					fs := strings.Fields(txt)

					if len(fs) < 3 {
						return fmt.Errorf("expecting <command> <arg> <test-name> [test-arg1 [test-arg-2 ...]]")
					}

					spec = &demand.Spec{
						Executable: fs[0],
						Checks: map[string]*demand.Check{
							"check": {
								Args: []string{fs[1]},
								Test: demand.Test{
									Name: fs[2],
									Args: fs[3:],
								},
							},
						},
					}
				}

				r, err := demand.Demand(spec)
				if err != nil {
					return fmt.Errorf("one liner: %w", err)
				}

				if err := emit(r); err != nil {
					return err
				}
			}

			for _, path := range c.Args().Slice() {
				r, err := demand.DemandPath(path)
				if err != nil {
					return fmt.Errorf("%s: %w", path, err)
				}

				if err := emit(r); err != nil {
					return err
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
