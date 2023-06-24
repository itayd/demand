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

	path, err := exec.LookPath(spec.Name)
	if err != nil {
		return &Result{}, nil
	}

	result := Result{
		FullPath: path,
		OK:       true,
	}

	for n, chk := range spec.Checks {
		chkr, err := check(path, chk)
		if err != nil {
			return nil, fmt.Errorf("check %q: %w", n, err)
		}

		if !chkr.OK {
			result.OK = false
		}

		result.Checks = append(result.Checks, chkr)
	}

	return &result, nil
}
