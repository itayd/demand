package demand

type TestResult struct {
	OK bool `json:"ok"`
}

type CheckResult struct {
	OK       bool        `json:"ok"`
	Args     []string    `json:"args"`
	ExitCode int         `json:"exit_code"`
	Capture  string      `json:"capture"`
	Test     *TestResult `json:"test,omitempty"`
}

type Result struct {
	OK         bool           `json:"ok"`
	Executable string         `json:"executable"`
	FullPath   string         `json:"full_path,omitempty"`
	Checks     []*CheckResult `json:"checks,omitempty"`
}
