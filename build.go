package semaphore

// Build model
type Build struct {
	URL         string `json:"build_url"`
	InfoURL     string `json:"build_info_url"`
	BuildNumber int    `json:"build_number"`
	Result      string `json:"result"`
	Started     string `json:"started_at"`
	Finished    string `json:"finished_at"`
	Commit      Commit `json:"commit"`
}

// RebuildResponse model
type BuildResponse struct {
	Commits      []Commit `json:"commits"`
	ProjectName  string   `json:"project_name"`
	BranchName   string   `json:"branch_name"`
	Number       int      `json:"number"`
	Result       string   `json:"result,omitempty"`
	Created      string   `json:"created_at"`
	Updated      string   `json:"updated_at"`
	Started      string   `json:"started_at,omitempty"`
	Finished     string   `json:"finished_at,omitempty"`
	HTMLURL      string   `json:"html_url"`
	BuildLogURL  string   `json:"build_log_url"`
	BuildInfoURL string   `json:"build_info_url"`
}
