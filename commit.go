package semaphore

type Commit struct {
	Id        string `json:"id"`
	Url       string `json:"url"`
	Author    string `json:"author_name"`
	Email     string `json:"author_email"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}
