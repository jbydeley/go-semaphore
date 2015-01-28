package semaphore

type BuildLog struct {
	Threads      []Thread `json:"threads"`
	BuildInfoUrl string   `json:"build_info_url"`
}
