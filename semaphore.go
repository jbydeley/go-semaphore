// Package semaphore API
package semaphore

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	// ErrHTTPNotFound occures when Semaphore returns a 404
	ErrHTTPNotFound = errors.New("Request returned 404")

	// ErrResponseNotRecognized occures when Semaphore returns json that can't be decoded
	ErrResponseNotRecognized = errors.New("Response could not be decoded")

	// ErrNotAuthorized occurs when Semaphore returns a 401 (usually due to auth token issues)
	ErrNotAuthorized = errors.New("Authentication failed. Check auth_token?")
)

const (
	getProjectsURL         = "%v/api/v1/projects?auth_token=%v"
	getBranchesURL         = "%v/api/v1/projects/%v/branches?auth_token=%v"
	getBranchStatusURL     = "%v/api/v1/projects/%v/%v/status?auth_token=%v"
	getBranchHistoryURL    = "%v/api/v1/projects/%v/%v?auth_token=%v"
	getBuildInformationURL = "%v/api/v1/projects/%v/%v/builds/%v?auth_token=%v"
	getBuildLogURL         = "%v/api/v1/projects/%v/%v/builds/%v/log?auth_token=%v"

	postRebuildBranch = "%v/api/v1/projects/%v/%v/build?auth_token=%v"
	postStopBuild     = "%v/api/v1/projects/%v/%v/builds/%v/stop?auth_token=%v"
)

func (a *Semaphore) request(URL string) ([]byte, error) {
	resp, err := a.httpClient.Get(URL)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 404:
		return nil, ErrHTTPNotFound
	case 401:
		return nil, ErrNotAuthorized
	default:
		return ioutil.ReadAll(resp.Body)
	}
}

func (a *Semaphore) post(URL string) ([]byte, error) {
	var empty []byte
	buf := bytes.NewReader(empty)
	resp, err := a.httpClient.Post(URL, "plain/text", buf)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 404:
		return nil, ErrHTTPNotFound
	case 401:
		return nil, ErrNotAuthorized
	default:
		return ioutil.ReadAll(resp.Body)
	}
}

// Semaphore API
type Semaphore struct {
	authToken, baseURL string
	httpClient         *http.Client
	log                *log.Logger
}

// NewSemaphore instantiates a new Semaphore API
func NewSemaphore(auth string) *Semaphore {
	logger := log.New(ioutil.Discard, "", log.LstdFlags)
	return NewSemaphoreWithLogger(auth, logger)
}

// NewSemaphoreWithLogger instantiates a new Semaphore API with a logger
func NewSemaphoreWithLogger(auth string, logger *log.Logger) *Semaphore {
	return &Semaphore{auth, "https://semaphoreapp.com", http.DefaultClient, logger}
}

// GetProjects returns an array of projects
func (a *Semaphore) GetProjects() ([]Project, error) {
	URL := fmt.Sprintf(getProjectsURL, a.baseURL, a.authToken)
	response, err := a.request(URL)
	if err != nil {
		a.log.Print(err)
		return nil, err
	}

	var projects []Project
	if err = json.Unmarshal(response, &projects); err != nil {
		return nil, ErrResponseNotRecognized
	}
	return projects, nil
}

// GetBranches returns an array of branches for a given project
func (a *Semaphore) GetBranches(project string) ([]Branch, error) {
	URL := fmt.Sprintf(getBranchesURL, a.baseURL, project, a.authToken)
	response, err := a.request(URL)

	if err != nil {
		a.log.Print(err)
		return nil, err
	}
	var branches []Branch
	if err = json.Unmarshal(response, &branches); err != nil {
		return nil, ErrResponseNotRecognized
	}
	return branches, nil
}

// GetBranchStatus returns a BranchStatus object for a given project + branch
func (a *Semaphore) GetBranchStatus(project string, branch int) (*BranchStatus, error) {
	URL := fmt.Sprintf(getBranchStatusURL, a.baseURL, project, branch, a.authToken)
	response, err := a.request(URL)
	if err != nil {
		a.log.Print(err)
		return nil, err
	}

	var status *BranchStatus
	if err = json.Unmarshal(response, &status); err != nil {
		return nil, ErrResponseNotRecognized
	}
	return status, nil
}

// GetBranchHistory returns a BranchHistory object for a given project + branch
func (a *Semaphore) GetBranchHistory(project string, branch int) (*BranchHistory, error) {
	URL := fmt.Sprintf(getBranchHistoryURL, a.baseURL, project, branch, a.authToken)
	response, err := a.request(URL)
	if err != nil {
		a.log.Print(err)
		return nil, err
	}

	var history *BranchHistory
	if err = json.Unmarshal(response, &history); err != nil {
		return nil, ErrResponseNotRecognized
	}
	return history, nil
}

// GetBuildInformation returns a BuildInformation object for a given project + branch + build
func (a *Semaphore) GetBuildInformation(project string, branch, build int) (*BuildInformation, error) {
	URL := fmt.Sprintf(getBuildInformationURL, a.baseURL, project, branch, build, a.authToken)
	response, err := a.request(URL)
	if err != nil {
		a.log.Print(err)
		return nil, err
	}

	var info *BuildInformation
	if err = json.Unmarshal(response, &info); err != nil {
		return nil, ErrResponseNotRecognized
	}
	return info, nil
}

// GetBuildLog returns a BuildLog object for a given project + branch + build
func (a *Semaphore) GetBuildLog(project string, branch, build int) (*BuildLog, error) {
	URL := fmt.Sprintf(getBuildLogURL, a.baseURL, project, branch, build, a.authToken)
	response, err := a.request(URL)
	if err != nil {
		a.log.Print(err)
		return nil, err
	}

	var log *BuildLog
	if err = json.Unmarshal(response, &log); err != nil {
		return nil, ErrResponseNotRecognized
	}
	return log, nil
}

// RebuildBranch returns a BuildResponse object for a given project + branch
func (a *Semaphore) RebuildBranch(project string, branch int) (*BuildResponse, error) {
	URL := fmt.Sprintf(postRebuildBranch, a.baseURL, project, branch, a.authToken)
	response, err := a.post(URL)
	if err != nil {
		a.log.Print(err)
		return nil, err
	}

	var buildResponse *BuildResponse
	if err = json.Unmarshal(response, &buildResponse); err != nil {
		return nil, ErrResponseNotRecognized
	}
	return buildResponse, nil
}

func (a *Semaphore) StopBuild(project string, branch, build int) (*BuildResponse, error) {
	URL := fmt.Sprintf(postStopBuild, a.baseURL, project, branch, build, a.authToken)
	response, err := a.post(URL)
	if err != nil {
		a.log.Print(err)
		return nil, err
	}

	var buildResponse *BuildResponse
	if err = json.Unmarshal(response, &buildResponse); err != nil {
		return nil, ErrResponseNotRecognized
	}
	return buildResponse, nil
}

// This is not yet implemented (API unclear and need to contact semaphore for clarification)
// LaunchBuild returns a BuildResponse object for a given project + branch + commit
// func (a *Semaphore) LaunchBuild(project string, branch int, commit string) (*BuildResponse, error) {
// 	URL := fmt.Sprintf(postLaunchBuild, a.baseURL, project, branch, commit, a.authToken)
// 	response, err := a.post(URL)
// 	if err != nil {
// 		a.log.Print(err)
// 		return nil, err
// 	}

// 	var buildResponse *BuildResponse
// 	if err = json.Unmarshal(response, &buildResponse); err != nil {
// 		return nil, ErrResponseNotRecognized
// 	}
// 	return buildResponse, nil
// }
