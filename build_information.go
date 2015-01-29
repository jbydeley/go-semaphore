package semaphore

// BuildInformation model
type BuildInformation struct {
	Commits     []Commit `json:"commits"`
	ProjectName string   `json:"project_name"`
	BranchName  string   `json:"branch_name"`
	Number      int      `json:"number"`
	Result      string   `json:"result"`
	Created     string   `json:"created_at"`
	Updated     string   `json:"updated_at"`
	Started     string   `json:"started_at"`
	Finished    string   `json:"finished_at"`
	HTMLURL     string   `json:"html_url"`
	BuildLogURL string   `json:"build_log_url"`
}
