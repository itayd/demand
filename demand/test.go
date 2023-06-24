package demand

import "fmt"

type Test struct {
	Name string   `json:"name"`
	Args []string `json:"args"`
}

type Tester struct {
	Run      func(args []string, in string) (pass bool, err error)
	Validate func(args []string) error
}

var tests = make(map[string]Tester)

func RegisterTest(name string, fn Tester) { tests[name] = fn }

func ValidateTest(t *Test) error {
	f, ok := tests[t.Name]
	if !ok {
		return fmt.Errorf("test %q is not defined", t.Name)
	}

	if err := f.Validate(t.Args); err != nil {
		return err
	}

	return nil
}

func RunTest(t *Test, in string) (bool, error) {
	f, ok := tests[t.Name]
	if !ok {
		return false, fmt.Errorf("test %q is not defined", t.Name)
	}

	pass, err := f.Run(t.Args, in)
	if err != nil {
		return false, fmt.Errorf("test %q: %w", t.Name, err)
	}

	return pass, nil
}
