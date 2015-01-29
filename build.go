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
