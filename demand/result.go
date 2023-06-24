package demand

type TestResult struct {
	OK       bool     `json:"ok"`
	Messages []string `json:"messages,omitempty"`
	Name     string   `json:"name"`
	Args     []string `json:"args,omitempty"`
}

type CheckResult struct {
	OK       bool        `json:"ok"`
	ExitCode int         `json:"exit_code,omitempty"`
	Args     []string    `json:"args"`
	Capture  string      `json:"capture"`
	Test     *TestResult `json:"test,omitempty"`
}

type Result struct {
	OK         bool           `json:"ok"`
	Executable string         `json:"executable"`
	FullPath   string         `json:"full_path,omitempty"`
	Checks     []*CheckResult `json:"checks,omitempty"`
}
