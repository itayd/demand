package demand

type CheckResult struct {
	Args       []string `json:"args"`
	ExitCode   int      `json:"exit_code"`
	CmdCapture string   `json:"cmd_capture"`
	OK         bool     `json:"ok"`
}

type Result struct {
	OK       bool           `json:"ok"`
	FullPath string         `json:"full_path"`
	Checks   []*CheckResult `json:"checks"`
}
