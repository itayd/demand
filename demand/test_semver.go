package demand

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/mod/semver"
)

func init() {
	RegisterTest("semver", Tester{
		Validate: func(args []string) error {
			if len(args) != 2 {
				return errors.New("expecting operator and version")
			}

			switch args[0] {
			case "<", "=", "==", ">", "<=", ">=":
				// nop
			default:
				return errors.New("invalid operator")
			}

			return nil
		},
		Run: func(args []string, in string) (bool, string, error) {
			if !strings.HasPrefix(in, "v") {
				in = "v" + in
			}

			v := args[1]
			if !strings.HasPrefix(v, "v") {
				v = "v" + v
			}

			if !semver.IsValid(in) {
				return false, "", fmt.Errorf("invalid checked version %q", in)
			}

			if !semver.IsValid(v) {
				return false, "", fmt.Errorf("invalid arg version %q", v)
			}

			cmp := semver.Compare(semver.Canonical(in), semver.Canonical(v))

			ok := false

			switch args[0] {
			case "=", "==":
				ok = cmp == 0
			case "<":
				ok = cmp < 0
			case ">":
				ok = cmp > 0
			case "<=":
				ok = cmp <= 0
			case ">=":
				ok = cmp >= 0
			default:
				panic("invalid operator")
			}

			return ok, "", nil
		},
	})
}
