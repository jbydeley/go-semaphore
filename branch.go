package semaphore

// Branch model
type Branch struct {
	ID   int    `json:"ID"`
	Name string `json:"name"`
	URL  string `json:"branch_url"`
}
