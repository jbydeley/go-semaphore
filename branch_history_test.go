package semaphore

import (
	"reflect"
	"testing"
)

func Test_GetBranchHistory_Success(t *testing.T) {
	expected := &BranchHistory{
		"gem_updates",
		"https://semaphoreapp.com/projects/44/branches/50",
		"https://semaphoreapp.com/api/v1/projects/649e584dc507ca4b73e1374d3125ef0b567a949c/89/status?auth_token=Yds3w6o26FLfJTnVK2y9",
		"https://semaphoreapp.com/api/v1/projects/649e584dc507ca4b73e1374d3125ef0b567a949c/89?auth_token=Yds3w6o26FLfJTnVK2y9",
		"base-app",
		[]Build{
			{
				"https://semaphoreapp.com/projects/57/branches/80/builds/46",
				"https://semaphoreapp.com/api/v1/projects/649e584dc507ca4b73e1374d3125ef0b567a949c/80/builds/1?auth_token=Yds3w6o26FLfJTnVK2y9",
				46,
				"failed",
				"2012-10-02T15:01:41Z",
				"2012-10-02T15:03:53Z",
				Commit{
					"a31d32d5de89613369f934eb7d30fbeb08883334",
					"https://github.com/renderedtext/base-app/commit/a31d32d5de89613369f934eb7d30fbeb08883334",
					"Vladimir Saric",
					"vladimir@renderedtext.com",
					"Update 'shoulda' gem.",
					"2012-10-02T07:00:14Z",
				},
			},
			{
				"https://semaphoreapp.com/projects/57/branches/80/builds/45",
				"https://semaphoreapp.com/api/v1/projects/649e584dc507ca4b73e1374d3125ef0b567a949c/80/builds/1?auth_token=Yds3w6o26FLfJTnVK2y9",
				45,
				"passed",
				"2012-10-02T14:47:06Z",
				"2012-10-02T14:51:35Z",
				Commit{
					"73fce130ad23f265add5d55ee1da1c23b38f85a4",
					"https://github.com/renderedtext/base-app/commit/73fce130ad23f265add5d55ee1da1c23b38f85a4",
					"Marko Anastasov",
					"marko@renderedtext.com",
					"Update 'factory_girl_rails' gem and use new short FactoryGirl syntax.",
					"2012-10-02T07:00:14Z",
				},
			},
		},
	}
	server, client := testAPICall(200, `
      {
         "branch_name": "gem_updates",
         "branch_URL": "https://semaphoreapp.com/projects/44/branches/50",
         "branch_status_URL": "https://semaphoreapp.com/api/v1/projects/649e584dc507ca4b73e1374d3125ef0b567a949c/89/status?auth_token=Yds3w6o26FLfJTnVK2y9",
         "branch_history_URL": "https://semaphoreapp.com/api/v1/projects/649e584dc507ca4b73e1374d3125ef0b567a949c/89?auth_token=Yds3w6o26FLfJTnVK2y9",
         "project_name": "base-app",
         "builds": [
            {
               "build_URL": "https://semaphoreapp.com/projects/57/branches/80/builds/46",
               "build_info_URL": "https://semaphoreapp.com/api/v1/projects/649e584dc507ca4b73e1374d3125ef0b567a949c/80/builds/1?auth_token=Yds3w6o26FLfJTnVK2y9",
               "build_number": 46,
               "result": "failed",
               "started_at": "2012-10-02T15:01:41Z",
               "finished_at": "2012-10-02T15:03:53Z",
               "commit": {
                  "id": "a31d32d5de89613369f934eb7d30fbeb08883334",
                  "URL": "https://github.com/renderedtext/base-app/commit/a31d32d5de89613369f934eb7d30fbeb08883334",
                  "author_name": "Vladimir Saric",
                  "author_email": "vladimir@renderedtext.com",
                  "message": "Update 'shoulda' gem.",
                  "timestamp": "2012-10-02T07:00:14Z"
               }
            },
            {
               "build_URL": "https://semaphoreapp.com/projects/57/branches/80/builds/45",
               "build_info_URL": "https://semaphoreapp.com/api/v1/projects/649e584dc507ca4b73e1374d3125ef0b567a949c/80/builds/1?auth_token=Yds3w6o26FLfJTnVK2y9",
               "build_number": 45,
               "result": "passed",
               "started_at": "2012-10-02T14:47:06Z",
               "finished_at": "2012-10-02T14:51:35Z",
               "commit": {
                  "id": "73fce130ad23f265add5d55ee1da1c23b38f85a4",
                  "URL": "https://github.com/renderedtext/base-app/commit/73fce130ad23f265add5d55ee1da1c23b38f85a4",
                  "author_name": "Marko Anastasov",
                  "author_email": "marko@renderedtext.com",
                  "message": "Update 'factory_girl_rails' gem and use new short FactoryGirl syntax.",
                  "timestamp": "2012-10-02T07:00:14Z"
               }
            }
         ]
      }`)
	defer server.Close()

	// Test the method!
	actual, _ := client.GetBranchHistory("123", 123)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Response did not match expected \nA: %v\n\n E: %v\n\n", actual, expected)
	}
}
