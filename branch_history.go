package semaphore

type BranchHistory struct {
	Name        string  `json:"branch_name"`
	Url         string  `json:"branch_url"`
	StatusUrl   string  `json:"branch_status_url"`
	HistoryUrl  string  `json:"branch_history_url"`
	ProjectName string  `json:"project_name"`
	Builds      []Build `json:"builds"`
}
