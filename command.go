package semaphore

// Command model
type Command struct {
	Name     string `json:"name"`
	Result   string `json:"result"`
	Output   string `json:"output"`
	Start    string `json:"start_time"`
	Finish   string `json:"finish_time"`
	Duration string `json:"duration"`
}
