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
	server, client := testApiCall(200, `
      {
         "branch_name": "gem_updates",
         "branch_url": "https://semaphoreapp.com/projects/44/branches/50",
         "branch_status_url": "https://semaphoreapp.com/api/v1/projects/649e584dc507ca4b73e1374d3125ef0b567a949c/89/status?auth_token=Yds3w6o26FLfJTnVK2y9",
         "branch_history_url": "https://semaphoreapp.com/api/v1/projects/649e584dc507ca4b73e1374d3125ef0b567a949c/89?auth_token=Yds3w6o26FLfJTnVK2y9",
         "project_name": "base-app",
         "build_url": "https://semaphoreapp.com/projects/44/branches/50/builds/15",
         "build_info_url": "https://semaphoreapp.com/api/v1/projects/649e584dc507ca4b73e1374d3125ef0b567a949c/89/builds/1?auth_token=Yds3w6o26FLfJTnVK2y9",
         "build_number": 15,
         "result": "passed",
         "started_at": "2012-07-09T15:23:53Z",
         "finished_at": "2012-07-09T15:30:16Z",
         "commit": {
            "id": "dc395381e650f3bac18457909880829fc20e34ba",
            "url": "https://github.com/renderedtext/base-app/commit/dc395381e650f3bac18457909880829fc20e34ba",
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
		t.Errorf("Response did not match expected \nA: %v\n\n E: %v\n\n", actual, expected)
	}
}
