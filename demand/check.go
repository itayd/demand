package demand

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

// All captured lines are stripped of leading and trailing spaces.
type Check struct {
	Args      []string `json:"args"` // will be supplied to command checked.
	FilterRE  string   `json:"filter_re"`
	CaptureRE string   `json:"capture_re"` // all lines will be fed joined by '\n's to checker.
	Test      Test     `json:"test"`

	filterRE, captureRE *regexp.Regexp
}

func (c *Check) compile() (err error) {
	if c.FilterRE != "" {
		if c.filterRE, err = regexp.Compile(c.FilterRE); err != nil {
			return fmt.Errorf("filter_re: %w", err)
		}
	}

	if c.CaptureRE != "" {
		if c.captureRE, err = regexp.Compile(c.CaptureRE); err != nil {
			return fmt.Errorf("capture_re: %w", err)
		}
	}

	if err := ValidateTest(&c.Test); err != nil {
		return fmt.Errorf("test: %w", err)
	}

	return nil
}

func check(path string, chk *Check) (*CheckResult, error) {
	r := CheckResult{Args: chk.Args}

	cmd := exec.Command(path, chk.Args...)

	out, err := cmd.CombinedOutput()
	if err != nil {
		r.ExitCode = cmd.ProcessState.ExitCode()
		return &r, nil
	}

	ls := strings.Split(string(out), "\n")

	var captured []string

	for _, l := range ls {
		if re := chk.filterRE; re != nil {
			if !re.MatchString(l) {
				continue
			}
		}

		if re := chk.captureRE; re != nil {
			m := re.FindString(l)

			if m == "" {
				continue
			}

			l = m
		}

		captured = append(captured, strings.TrimSpace(l))
	}

	r.Capture = strings.TrimSpace(strings.Join(captured, "\n"))

	r.Test, err = RunTest(&chk.Test, r.Capture)

	if r.Test != nil {
		r.OK = r.Test.OK
	}

	return &r, err
}
