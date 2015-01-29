package semaphore

// Commit model
type Commit struct {
	ID        string `json:"ID"`
	URL       string `json:"url"`
	Author    string `json:"author_name"`
	Email     string `json:"author_email"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}
