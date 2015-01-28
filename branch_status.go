package semaphore

type BranchStatus struct {
	Name         string `json:"branch_name"`
	BranchUrl    string `json:"branch_url"`
	StatusUrl    string `json:"branch_status_url"`
	HistoryUrl   string `json:"branch_history_url"`
	ProjectName  string `json:"project_name"`
	BuildUrl     string `json:"build_url"`
	BuildInfoUrl string `json:"build_info_url"`
	BuildNumber  int    `json:"build_number"`
	Result       string `json:"result"`
	Started      string `json:"started_at"`
	Finished     string `json:"finished_at"`
	Commit       Commit `json:"commit,omitempty"`
}
