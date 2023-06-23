package demand

import (
	"fmt"
	"log"
	"os/exec"
)

type Config struct{}

var DefaultConfig Config

var (
	DemandPath = DefaultConfig.DemandPath
	Demand     = DefaultConfig.Demand
)

func (c Config) DemandPath(path string) error {
	spec, err := ReadSpec(path)
	if err != nil {
		return fmt.Errorf("spec: %w", err)
	}

	return c.Demand(spec)
}

func (Config) Demand(spec *Spec) error {
	s := *spec
	spec = &s

	if err := spec.compile(); err != nil {
		return err
	}

	path, err := exec.LookPath(spec.Name)
	if err != nil {
		return err
	}

	log.Printf("%q: found at %q", spec.Name, path)

	for n, chk := range spec.Checks {
		if err := check(path, chk); err != nil {
			return fmt.Errorf("check %q: %w", n, err)
		}
	}

	return nil
}
