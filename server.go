package semaphore

type Server struct {
	ProjectName  string `json:"project_name"`
	ServerName   string `json:"server_name"`
	Number       int    `json:"number"`
	Result       string `json:"result"`
	Created      string `json:"created_at"`
	Updated      string `json:"updated_at"`
	Started      string `json:"started_at"`
	Finished     string `json:"finished_at"`
	HtmlUrl      string `json:"html_url"`
	DeployUrl    string `json:"deploy_url"`
	DeployLogUrl string `json:"deploy_log_url"`
	BuildUrl     string `json:"build_url"`
	BuildHtmlUrl string `json:"build_html_url"`
	Commit       Commit `json:"commit"`
}
