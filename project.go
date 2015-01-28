package semaphore

const (
	PROJECT_URL = "https://semaphoreapp.com/api/v1/projects"
)

type Project struct {
	Id       int            `json:"id"`
	HashId   string         `json:"hash_id"`
	Name     string         `json:"name"`
	Owner    string         `json:"owner"`
	HtmlUrl  string         `json:"html_url"`
	Created  string         `json:"created_at"`
	Updated  string         `json:"updated_at"`
	Branches []BranchStatus `json:"branches"`
	Servers  []Server       `json:"servers"`
}
