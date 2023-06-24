package demand

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

type Spec struct {
	Executable string            `json:"executable"`
	Checks     map[string]*Check `json:"checks"`
}

func (s *Spec) compile() error {
	for n, chk := range s.Checks {
		if err := chk.compile(); err != nil {
			return fmt.Errorf("check %q: %w", n, err)
		}
	}

	return nil
}

func ReadSpec(path string) (*Spec, error) {
	bs, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read: %w", err)
	}

	spec, err := ParseSpec(bs)
	if err != nil {
		return nil, fmt.Errorf("parse: %w", err)
	}

	return spec, nil
}

func ParseSpec(in []byte) (*Spec, error) {
	decoder := json.NewDecoder(bytes.NewReader(in))
	decoder.DisallowUnknownFields()

	var spec Spec

	if err := decoder.Decode(&spec); err != nil {
		return nil, err
	}

	return &spec, nil
}
