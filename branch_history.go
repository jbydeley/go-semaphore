package semaphore

// BranchHistory model
type BranchHistory struct {
	Name        string  `json:"branch_name"`
	URL         string  `json:"branch_url"`
	StatusURL   string  `json:"branch_status_url"`
	HistoryURL  string  `json:"branch_history_url"`
	ProjectName string  `json:"project_name"`
	Builds      []Build `json:"builds"`
}
