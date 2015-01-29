package semaphore

type BranchStatus struct {
	Name         string `json:"branch_name"`
	BranchURL    string `json:"branch_url"`
	StatusURL    string `json:"branch_status_url"`
	HistoryURL   string `json:"branch_history_url"`
	ProjectName  string `json:"project_name"`
	BuildURL     string `json:"build_url"`
	BuildInfoURL string `json:"build_info_url"`
	BuildNumber  int    `json:"build_number"`
	Result       string `json:"result"`
	Started      string `json:"started_at"`
	Finished     string `json:"finished_at"`
	Commit       Commit `json:"commit,omitempty"`
}
