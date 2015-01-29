package semaphore

// Project model
type Project struct {
	ID       int            `json:"ID"`
	HashID   string         `json:"hash_ID"`
	Name     string         `json:"name"`
	Owner    string         `json:"owner"`
	HTML     string         `json:"html_url"`
	Created  string         `json:"created_at"`
	Updated  string         `json:"updated_at"`
	Branches []BranchStatus `json:"branches"`
	Servers  []Server       `json:"servers"`
}
