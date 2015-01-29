// Package semaphore API
package semaphore

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	getProjectsURL         = "%v/api/v1/projects?auth_token=%v"
	getBranchesURL         = "%v/api/v1/projects/%v/branches?auth_token=%v"
	getBranchStatusURL     = "%v/api/v1/projects/%v/%v/status?auth_token=%v"
	getBranchHistoryURL    = "%v/api/v1/projects/%v/%v?auth_token=%v"
	getBuildInformationURL = "%v/api/v1/projects/%v/%v/builds/%v?auth_token=%v"
	getBuildLogURL         = "%v/api/v1/projects/%v/%v/builds/%v/log?auth_token=%v"
)

func (a *Semaphore) request(URL string) ([]byte, error) {
	resp, err := a.httpClient.Get(URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// Semaphore API
type Semaphore struct {
	authToken, baseURL string
	httpClient         *http.Client
}

// NewSemaphore instantiates a new Semaphore API
func NewSemaphore(auth string) *Semaphore {
	return &Semaphore{auth, "https://semaphoreapp.com", http.DefaultClient}
}

// GetProjects returns an array of projects
func (a *Semaphore) GetProjects() ([]Project, error) {
	URL := fmt.Sprintf(getProjectsURL, a.baseURL, a.authToken)
	response, err := a.request(URL)
	if err != nil {
		return nil, err
	}

	var projects []Project
	if err = json.Unmarshal(response, &projects); err != nil {
		return nil, err
	}
	return projects, nil
}

// GetBranches returns an array of branches for a given project
func (a *Semaphore) GetBranches(project string) ([]Branch, error) {
	URL := fmt.Sprintf(getBranchesURL, a.baseURL, project, a.authToken)
	response, err := a.request(URL)
	if err != nil {
		return nil, err
	}

	var branches []Branch
	if err = json.Unmarshal(response, &branches); err != nil {
		return nil, err
	}
	return branches, nil
}

// GetBranchStatus returns a BranchStatus object for a given project + branch
func (a *Semaphore) GetBranchStatus(project string, branch int) (*BranchStatus, error) {
	URL := fmt.Sprintf(getBranchStatusURL, a.baseURL, project, branch, a.authToken)
	response, err := a.request(URL)
	if err != nil {
		return nil, err
	}

	var status *BranchStatus
	if err = json.Unmarshal(response, &status); err != nil {
		return nil, err
	}
	return status, nil
}

// GetBranchHistory returns a BranchHistory object for a given project + branch
func (a *Semaphore) GetBranchHistory(project string, branch int) (*BranchHistory, error) {
	URL := fmt.Sprintf(getBranchHistoryURL, a.baseURL, project, branch, a.authToken)
	response, err := a.request(URL)
	if err != nil {
		return nil, err
	}

	var history *BranchHistory
	if err = json.Unmarshal(response, &history); err != nil {
		return nil, err
	}
	return history, nil
}

// GetBuildInformation returns a BuildInformation object for a given project + branch + build
func (a *Semaphore) GetBuildInformation(project string, branch, build int) (*BuildInformation, error) {
	URL := fmt.Sprintf(getBuildInformationURL, a.baseURL, project, branch, build, a.authToken)
	response, err := a.request(URL)
	if err != nil {
		return nil, err
	}

	var info *BuildInformation
	if err = json.Unmarshal(response, &info); err != nil {
		return nil, err
	}
	return info, nil
}

// GetBuildLog returns a BuildLog object for a given project + branch + build
func (a *Semaphore) GetBuildLog(project string, branch, build int) (*BuildLog, error) {
	URL := fmt.Sprintf(getBuildLogURL, a.baseURL, project, branch, build, a.authToken)
	response, err := a.request(URL)
	if err != nil {
		return nil, err
	}

	var log *BuildLog
	if err = json.Unmarshal(response, &log); err != nil {
		return nil, err
	}
	return log, nil
}
