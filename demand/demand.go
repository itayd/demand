package demand

import (
	"fmt"
	"os/exec"
)

type Config struct{}

var DefaultConfig Config

var (
	DemandPath = DefaultConfig.DemandPath
	Demand     = DefaultConfig.Demand
)

func (c Config) DemandPath(path string) (*Result, error) {
	spec, err := ReadSpec(path)
	if err != nil {
		return nil, fmt.Errorf("spec: %w", err)
	}

	return c.Demand(spec)
}

func (Config) Demand(spec *Spec) (*Result, error) {
	s := *spec
	spec = &s

	if err := spec.compile(); err != nil {
		return nil, err
	}

	r := Result{Executable: spec.Executable}

	path, err := exec.LookPath(spec.Executable)
	if err != nil {
		return &r, nil
	}

	r.FullPath = path
	r.OK = true

	for n, chk := range spec.Checks {
		chkr, err := check(path, chk)
		if err != nil {
			return nil, fmt.Errorf("check %q: %w", n, err)
		}

		if !chkr.OK {
			r.OK = false
		}

		r.Checks = append(r.Checks, chkr)
	}

	return &r, nil
}
