// Go-Semaphore

// // For Example
// //  api := NewApi("authorization_token")
// //  branch, err := api.GetBranchStatus("hash_id", branch_id)

package semaphore

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	GET_PROJECTS_URL          = "%v/api/v1/projects?auth_token=%v"
	GET_BRANCHES_URL          = "%v/api/v1/projects/%v/branches?auth_token=%v"
	GET_BRANCH_STATUS_URL     = "%v/api/v1/projects/%v/%v/status?auth_token=%v"
	GET_BRANCH_HISTORY_URL    = "%v/api/v1/projects/%v/%v?auth_token=%v"
	GET_BUILD_INFORMATION_URL = "%v/api/v1/projects/%v/%v/builds/%v?auth_token=%v"
	GET_BUILD_LOG_URL         = "%v/api/v1/projects/%v/%v/builds/%v/log?auth_token=%v"
)

func (a *Api) request(url string) ([]byte, error) {
	resp, err := a.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

type Api struct {
	authToken, baseUrl string
	httpClient         *http.Client
}

func NewApi(auth string) *Api {
	return &Api{auth, "https://semaphoreapp.com", http.DefaultClient}
}

func (a *Api) GetProjects() ([]Project, error) {
	url := fmt.Sprintf(GET_PROJECTS_URL, a.baseUrl, a.authToken)
	response, err := a.request(url)
	if err != nil {
		return nil, err
	}

	var projects []Project
	if err = json.Unmarshal(response, &projects); err != nil {
		return nil, err
	}
	return projects, nil
}

func (a *Api) GetBranches(project string) ([]Branch, error) {
	url := fmt.Sprintf(GET_BRANCHES_URL, a.baseUrl, project, a.authToken)
	response, err := a.request(url)
	if err != nil {
		return nil, err
	}

	var branches []Branch
	if err = json.Unmarshal(response, &branches); err != nil {
		return nil, err
	}
	return branches, nil
}

func (a *Api) GetBranchStatus(project string, branch int) (*BranchStatus, error) {
	url := fmt.Sprintf(GET_BRANCH_STATUS_URL, a.baseUrl, project, branch, a.authToken)
	response, err := a.request(url)
	if err != nil {
		return nil, err
	}

	var status *BranchStatus
	if err = json.Unmarshal(response, &status); err != nil {
		return nil, err
	}
	return status, nil
}

func (a *Api) GetBranchHistory(project string, branch int) (*BranchHistory, error) {
	url := fmt.Sprintf(GET_BRANCH_HISTORY_URL, a.baseUrl, project, branch, a.authToken)
	response, err := a.request(url)
	if err != nil {
		return nil, err
	}

	var history *BranchHistory
	if err = json.Unmarshal(response, &history); err != nil {
		return nil, err
	}
	return history, nil
}

func (a *Api) GetBuildInformation(project string, branch, build int) (*BuildInformation, error) {
	url := fmt.Sprintf(GET_BUILD_INFORMATION_URL, a.baseUrl, project, branch, build, a.authToken)
	response, err := a.request(url)
	if err != nil {
		return nil, err
	}

	var info *BuildInformation
	if err = json.Unmarshal(response, &info); err != nil {
		return nil, err
	}
	return info, nil
}

func (a *Api) GetBuildLog(project string, branch, build int) (*BuildLog, error) {
	url := fmt.Sprintf(GET_BUILD_LOG_URL, a.baseUrl, project, branch, build, a.authToken)
	response, err := a.request(url)
	if err != nil {
		return nil, err
	}

	var log *BuildLog
	if err = json.Unmarshal(response, &log); err != nil {
		return nil, err
	}
	return log, nil
}
