package demand

import (
	"errors"
	"fmt"
	"strings"

	"github.com/hashicorp/go-version"
)

func init() {
	RegisterTester("semver", Tester{
		Validate: func(args []string) error {
			if len(args) == 0 {
				return errors.New("expecting at least a single constraint")
			}

			_, err := version.NewConstraint(strings.Join(args, ","))

			return err
		},
		Run: func(args []string, in string) (*TestResult, error) {
			c, err := version.NewConstraint(args[0])
			if err != nil {
				return nil, fmt.Errorf("invalid constraint: %w", err)
			}

			v1, err := version.NewVersion(in)
			if err != nil {
				return &TestResult{
					Messages: []string{err.Error()},
				}, nil
			}

			return &TestResult{OK: c.Check(v1)}, nil
		},
	})
}
