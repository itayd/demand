package demand

import "fmt"

type Test struct {
	Name string   `json:"name"`
	Args []string `json:"args"`
}

type Tester struct {
	Run      func(args []string, in string) (*TestResult, error)
	Validate func(args []string) error
}

var tests = make(map[string]Tester)

func RegisterTester(name string, fn Tester) { tests[name] = fn }

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

func RunTest(t *Test, in string) (*TestResult, error) {
	f, ok := tests[t.Name]
	if !ok {
		return nil, fmt.Errorf("test %q is not defined", t.Name)
	}

	r, err := f.Run(t.Args, in)
	if err != nil {
		return nil, fmt.Errorf("test %q: %w", t.Name, err)
	}

	return r, nil
}
