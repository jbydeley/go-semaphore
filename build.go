package semaphore

type Build struct {
	Url         string `json:"build_url"`
	InfoUrl     string `json:"build_info_url"`
	BuildNumber int    `json:"build_number"`
	Result      string `json:"result"`
	Started     string `json:"started_at"`
	Finished    string `json:"finished_at"`
	Commit      Commit `json:"commit"`
}
