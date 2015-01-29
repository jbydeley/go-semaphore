package semaphore

type BuildLog struct {
	Threads      []Thread `json:"threads"`
	BuildInfoURL string   `json:"build_info_url"`
}
