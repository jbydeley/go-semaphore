package semaphore

import (
	"reflect"
	"testing"
)

func Test_GetBuildInformation_Success(t *testing.T) {
	expected := &BuildInformation{
		[]Commit{
			{
				"ce0d543b875884f09cf1e287fb303fb91a9e28a0",
				"https://github.com/renderedtext/base-app/commit/ce0d543b875884f09cf1e287fb303fb91a9e28a0",
				"Marko Anastasov",
				"marko@renderedtext.com",
				"Upgrade shoulda 1.1.6 -> 1.2.1",
				"2014-05-16T15:38:49+02:00",
			},
		},
		"base-app",
		"master",
		44,
		"passed",
		"2014-06-19T09:37:18+02:00",
		"2014-06-19T09:39:42+02:00",
		"2014-06-19T09:37:26+02:00",
		"2014-06-19T09:39:42+02:00",
		"https://semaphoreapp.com/renderedtext/base-app/branches/master/builds/44",
		"https://semaphoreapp.com/api/v1/projects/73c4b979-0a40-49db-b10e-571d41e10d9a/133529/builds/44/log?auth_token=:auth_token",
	}
	server, client := testAPICall(200, `
      {
        "commits": [{
          "ID": "ce0d543b875884f09cf1e287fb303fb91a9e28a0",
          "URL": "https://github.com/renderedtext/base-app/commit/ce0d543b875884f09cf1e287fb303fb91a9e28a0",
          "author_name": "Marko Anastasov",
          "author_email": "marko@renderedtext.com",
          "message": "Upgrade shoulda 1.1.6 -> 1.2.1",
          "timestamp": "2014-05-16T15:38:49+02:00"
        }],
        "project_name": "base-app",
        "branch_name": "master",
        "number": 44,
        "result": "passed",
        "created_at": "2014-06-19T09:37:18+02:00",
        "updated_at": "2014-06-19T09:39:42+02:00",
        "started_at": "2014-06-19T09:37:26+02:00",
        "finished_at": "2014-06-19T09:39:42+02:00",
        "html_URL": "https://semaphoreapp.com/renderedtext/base-app/branches/master/builds/44",
        "build_log_URL": "https://semaphoreapp.com/api/v1/projects/73c4b979-0a40-49db-b10e-571d41e10d9a/133529/builds/44/log?auth_token=:auth_token"
      }`)
	defer server.Close()

	// Test the method!
	actual, _ := client.GetBuildInformation("hash_ID", 1234, 4321)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Response dID not match expected \nA: %v\n\n E: %v\n\n", actual, expected)
	}
}
