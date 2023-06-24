package demand

type TestResult struct {
	OK       bool     `json:"ok"`
	Messages []string `json:"messages,omitempty"`
}

type CheckResult struct {
	OK       bool        `json:"ok"`
	Args     []string    `json:"args"`
	ExitCode int         `json:"exit_code"`
	Capture  string      `json:"capture"`
	Test     *TestResult `json:"test,omitempty"`
}

type Result struct {
	OK       bool           `json:"ok"`
	FullPath string         `json:"full_path"`
	Checks   []*CheckResult `json:"checks"`
}
