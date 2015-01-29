package semaphore

// Server model
type Server struct {
	ProjectName  string `json:"project_name"`
	ServerName   string `json:"server_name"`
	Number       int    `json:"number"`
	Result       string `json:"result"`
	Created      string `json:"created_at"`
	Updated      string `json:"updated_at"`
	Started      string `json:"started_at"`
	Finished     string `json:"finished_at"`
	HTMLURL      string `json:"html_url"`
	DeployURL    string `json:"deploy_url"`
	DeployLogURL string `json:"deploy_log_url"`
	BuildURL     string `json:"build_url"`
	BuildHTMLURL string `json:"build_html_url"`
	Commit       Commit `json:"commit"`
}
