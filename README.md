# Go-Semaphore API SDK
[![Build Status](https://semaphoreapp.com/api/v1/projects/3316409d-5584-457e-b1ab-227ff98210e0/337417/badge.png)](https://semaphoreapp.com/jbydeley/go-semaphore)
[![Coverage Status](https://coveralls.io/repos/jbydeley/go-semaphore/badge.svg?branch=master)](https://coveralls.io/r/jbydeley/go-semaphore?branch=master)
[![GoDoc](https://godoc.org/github.com/jbydeley/go-semaphore?status.svg)](https://godoc.org/github.com/jbydeley/go-semaphore)

Go-Semaphore is an SDK to map to the SemaphoreApp API. This lets you check on build status, branch status and branch history. Check the [GoDoc link](https://godoc.org/github.com/jbydeley/go-semaphore#Semaphore) for a list of functions you can call. Check the [semaphore api](https://semaphoreapp.com/docs/api.html) for more information.

**TODO:** Implement POST methods.

# Installation

Get the package
`go get -u github.com/jbydeley/go-semaphore`

Import it
`import "github.com/jbydeley/go-semaphore"`

# Example Usage

## Get a list of projects

```
api := semaphore.NewSemaphore("auth_token")
projects, err := api.GetProjects()
```

[GetProjects](https://godoc.org/github.com/jbydeley/go-semaphore#Semaphore.GetProjects) returns a [project](https://godoc.org/github.com/jbydeley/go-semaphore#Project).

## Get a list of branches for a project

```
api := semaphore.NewSemaphore("auth_token")
branches, err := api.GetBranches("project_hash")
```

[GetBranches](https://godoc.org/github.com/jbydeley/go-semaphore#Semaphore.GetBranches) returns an array of [branches](https://godoc.org/github.com/jbydeley/go-semaphore#Branch).

## Get the status of a branch

```
api := semaphore.NewSemaphore("auth_token")
branchStatus, err := api.GetBranchStatus("project_hash", branch_id)
```

[GetBranchStatus](https://godoc.org/github.com/jbydeley/go-semaphore#Semaphore.GetBranchStatus) returns a [branch status](https://godoc.org/github.com/jbydeley/go-semaphore#BranchStatus).
