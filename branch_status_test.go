package semaphore

import (
	"reflect"
	"testing"
)

func Test_GetBranchStatus_Success(t *testing.T) {
	expected := &BranchStatus{
		"gem_updates",
		"https://semaphoreapp.com/projects/44/branches/50",
		"https://semaphoreapp.com/api/v1/projects/649e584dc507ca4b73e1374d3125ef0b567a949c/89/status?auth_token=Yds3w6o26FLfJTnVK2y9",
		"https://semaphoreapp.com/api/v1/projects/649e584dc507ca4b73e1374d3125ef0b567a949c/89?auth_token=Yds3w6o26FLfJTnVK2y9",
		"base-app",
		"https://semaphoreapp.com/projects/44/branches/50/builds/15",
		"https://semaphoreapp.com/api/v1/projects/649e584dc507ca4b73e1374d3125ef0b567a949c/89/builds/1?auth_token=Yds3w6o26FLfJTnVK2y9",
		15,
		"passed",
		"2012-07-09T15:23:53Z",
		"2012-07-09T15:30:16Z",
		Commit{
			"dc395381e650f3bac18457909880829fc20e34ba",
			"https://github.com/renderedtext/base-app/commit/dc395381e650f3bac18457909880829fc20e34ba",
			"Vladimir Saric",
			"vladimir@renderedtext.com",
			"Update 'shoulda' gem.",
			"2012-07-04T18:14:08Z",
		},
	}
	server, client := testAPICall(200, `
      {
         "branch_name": "gem_updates",
         "branch_URL": "https://semaphoreapp.com/projects/44/branches/50",
         "branch_status_URL": "https://semaphoreapp.com/api/v1/projects/649e584dc507ca4b73e1374d3125ef0b567a949c/89/status?auth_token=Yds3w6o26FLfJTnVK2y9",
         "branch_history_URL": "https://semaphoreapp.com/api/v1/projects/649e584dc507ca4b73e1374d3125ef0b567a949c/89?auth_token=Yds3w6o26FLfJTnVK2y9",
         "project_name": "base-app",
         "build_URL": "https://semaphoreapp.com/projects/44/branches/50/builds/15",
         "build_info_URL": "https://semaphoreapp.com/api/v1/projects/649e584dc507ca4b73e1374d3125ef0b567a949c/89/builds/1?auth_token=Yds3w6o26FLfJTnVK2y9",
         "build_number": 15,
         "result": "passed",
         "started_at": "2012-07-09T15:23:53Z",
         "finished_at": "2012-07-09T15:30:16Z",
         "commit": {
            "ID": "dc395381e650f3bac18457909880829fc20e34ba",
            "URL": "https://github.com/renderedtext/base-app/commit/dc395381e650f3bac18457909880829fc20e34ba",
            "author_name": "Vladimir Saric",
            "author_email": "vladimir@renderedtext.com",
            "message": "Update 'shoulda' gem.",
            "timestamp": "2012-07-04T18:14:08Z"
         }
      }`)
	defer server.Close()

	// Test the method!
	actual, _ := client.GetBranchStatus("123", 123)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Response dID not match expected \nA: %v\n\n E: %v\n\n", actual, expected)
	}
}

func Test_GetBranchStatus_NotFound(t *testing.T) {
	server, client := testAPICall(404, "")
	defer server.Close()

	results, err := client.GetBranchStatus("123", 456)

	if err != ErrHTTPNotFound {
		t.Errorf("Should have recieved 404 but got '%v'", err)
	}

	if results != nil {
		t.Errorf("Results should be nil but got '%+v'", results)
	}
}

func Test_GetBranchStatus_NotRecognized(t *testing.T) {
	server, client := testAPICall(200, "???")
	defer server.Close()

	results, err := client.GetBranchStatus("123", 456)

	if err != ErrResponseNotRecognized {
		t.Errorf("Should have received an ErrResponseNotRecognized but got '%v'", err)
	}

	if results != nil {
		t.Errorf("Results should be nil but got '%+v'", results)
	}
}

func Test_GetBranchStatus_FailedAuthentication(t *testing.T) {
	server, client := testAPICall(401, "")
	defer server.Close()

	results, err := client.GetBranchStatus("123", 456)

	if err != ErrNotAuthorized {
		t.Errorf("Should have recieved 401 but got '%v'", err)
	}

	if results != nil {
		t.Errorf("Results should be nil but got '%+v'", results)
	}
}
