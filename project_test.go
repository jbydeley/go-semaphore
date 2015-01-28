package semaphore

import (
	"reflect"
	"testing"
)

func Test_GetProjects_Success(t *testing.T) {
	expected := []Project{
		Project{
			61,
			"3f1004b8343faabda63d441734526c854380ab89",
			"testapp-sphinx",
			"renderedtext",
			"https://semaphoreapp.com/darkofabijan/testapp-sphinx",
			"2012-09-04T11:53:22Z",
			"2012-09-04T12:01:17Z",
			[]BranchStatus{
				BranchStatus{
					"master",
					"https://semaphoreapp.com/projects/61/branches/85",
					"https://semaphoreapp.com/api/v1/projects/3f1004b8343faabda63d441734526c854380ab89/85/status?auth_token=Yds3w6o26FLfJTnVK2y9",
					"https://semaphoreapp.com/api/v1/projects/3f1004b8343faabda63d441734526c854380ab89/85?auth_token=Yds3w6o26FLfJTnVK2y9",
					"testapp-sphinx",
					"https://semaphoreapp.com/projects/61/branches/85/builds/1",
					"https://semaphoreapp.com/api/v1/projects/3f1004b8343faabda63d441734526c854380ab89/85/builds/1?auth_token=Yds3w6o26FLfJTnVK2y9",
					1,
					"passed",
					"2012-09-04T11:55:07Z",
					"2012-09-04T12:01:16Z",
					Commit{},
				},
			},
			[]Server{
				Server{
					"heroku-deploy-test",
					"server-heroku-master-automatic-1",
					1,
					"passed",
					"2013-07-19T14:57:18+02:00",
					"2013-07-19T14:58:49+02:00",
					"2013-07-19T14:57:24+02:00",
					"2013-07-19T14:58:49+02:00",
					"https://semaphoreapp.com/projects/2420/servers/80/deploys/1",
					"https://semaphoreapp.com/api/v1/projects/dccc4ab03f2b47167cac51553d58e468c6750c2c/servers/80/deploys/1?auth_token=qH36J7zzMfAGG8xmF72f",
					"https://semaphoreapp.com/api/v1/projects/dccc4ab03f2b47167cac51553d58e468c6750c2c/servers/80/deploys/1/log?auth_token=qH36J7zzMfAGG8xmF72f",
					"https://semaphoreapp.com/api/v1/projects/dccc4ab03f2b47167cac51553d58e468c6750c2c/58394/builds/5?auth_token=qH36J7zzMfAGG8xmF72f",
					"https://semaphoreapp.com/projects/2420/branches/58394/builds/5",
					Commit{
						"43ddb7516ecc743f0563abd7418f0bd3617348c4",
						"https://github.com/rastasheep/heroku-deploy-test/commit/43ddb7516ecc743f0563abd7418f0bd3617348c4",
						"Aleksandar Diklic",
						"rastasheep3@gmail.com",
						"One more time",
						"2013-07-19T14:56:25+02:00",
					},
				},
			},
		},
		Project{
			63,
			"649e584dc507ca4b73e1374d3125ef0b567a949c",
			"testapp-mongodb",
			"renderedtext",
			"https://semaphoreapp.com/darkofabijan/testapp-mongodb",
			"2012-09-14T10:53:38Z",
			"2012-09-14T11:16:51Z",
			[]BranchStatus{
				BranchStatus{
					"mongoid3",
					"https://semaphoreapp.com/projects/63/branches/89",
					"https://semaphoreapp.com/api/v1/projects/3f1004b8343faabda63d441734526c854380ab89/85/status?auth_token=Yds3w6o26FLfJTnVK2y9",
					"https://semaphoreapp.com/api/v1/projects/3f1004b8343faabda63d441734526c854380ab89/85?auth_token=Yds3w6o26FLfJTnVK2y9",
					"testapp-mongodb",
					"https://semaphoreapp.com/projects/63/branches/89/builds/3",
					"https://semaphoreapp.com/api/v1/projects/3f1004b8343faabda63d441734526c854380ab89/85/builds/1?auth_token=Yds3w6o26FLfJTnVK2y9",
					3,
					"passed",
					"2012-09-14T11:11:39Z",
					"2012-09-14T11:16:51Z",
					Commit{},
				},
			},
			[]Server{},
		},
	}
	server, client := testApiCall(200, `
      [
         {
            "id": 61,
            "hash_id": "3f1004b8343faabda63d441734526c854380ab89",
            "name": "testapp-sphinx",
            "owner": "renderedtext",
            "html_url": "https://semaphoreapp.com/darkofabijan/testapp-sphinx",
            "created_at": "2012-09-04T11:53:22Z",
            "updated_at": "2012-09-04T12:01:17Z",
            "branches": [
               {
                  "branch_name": "master",
                  "branch_url": "https://semaphoreapp.com/projects/61/branches/85",
                  "branch_status_url": "https://semaphoreapp.com/api/v1/projects/3f1004b8343faabda63d441734526c854380ab89/85/status?auth_token=Yds3w6o26FLfJTnVK2y9",
                  "branch_history_url": "https://semaphoreapp.com/api/v1/projects/3f1004b8343faabda63d441734526c854380ab89/85?auth_token=Yds3w6o26FLfJTnVK2y9",
                  "project_name": "testapp-sphinx",
                  "build_url": "https://semaphoreapp.com/projects/61/branches/85/builds/1",
                  "build_info_url": "https://semaphoreapp.com/api/v1/projects/3f1004b8343faabda63d441734526c854380ab89/85/builds/1?auth_token=Yds3w6o26FLfJTnVK2y9",
                  "build_number": 1,
                  "result": "passed",
                  "started_at": "2012-09-04T11:55:07Z",
                  "finished_at": "2012-09-04T12:01:16Z"
               }
            ],
            "servers": [
               {
                  "project_name": "heroku-deploy-test",
                  "server_name": "server-heroku-master-automatic-1",
                  "number": 1,
                  "result": "passed",
                  "created_at": "2013-07-19T14:57:18+02:00",
                  "updated_at": "2013-07-19T14:58:49+02:00",
                  "started_at": "2013-07-19T14:57:24+02:00",
                  "finished_at": "2013-07-19T14:58:49+02:00",
                  "html_url": "https://semaphoreapp.com/projects/2420/servers/80/deploys/1",
                  "deploy_url": "https://semaphoreapp.com/api/v1/projects/dccc4ab03f2b47167cac51553d58e468c6750c2c/servers/80/deploys/1?auth_token=qH36J7zzMfAGG8xmF72f",
                  "deploy_log_url": "https://semaphoreapp.com/api/v1/projects/dccc4ab03f2b47167cac51553d58e468c6750c2c/servers/80/deploys/1/log?auth_token=qH36J7zzMfAGG8xmF72f",
                  "build_url": "https://semaphoreapp.com/api/v1/projects/dccc4ab03f2b47167cac51553d58e468c6750c2c/58394/builds/5?auth_token=qH36J7zzMfAGG8xmF72f",
                  "build_html_url": "https://semaphoreapp.com/projects/2420/branches/58394/builds/5",
                  "commit": {
                     "id": "43ddb7516ecc743f0563abd7418f0bd3617348c4",
                     "url": "https://github.com/rastasheep/heroku-deploy-test/commit/43ddb7516ecc743f0563abd7418f0bd3617348c4",
                     "author_name": "Aleksandar Diklic",
                     "author_email": "rastasheep3@gmail.com",
                     "message": "One more time",
                     "timestamp": "2013-07-19T14:56:25+02:00"
                  }
               }
            ]
         },
         {
            "id": 63,
            "hash_id": "649e584dc507ca4b73e1374d3125ef0b567a949c",
            "name": "testapp-mongodb",
            "owner": "renderedtext",
            "html_url": "https://semaphoreapp.com/darkofabijan/testapp-mongodb",
            "created_at": "2012-09-14T10:53:38Z",
            "updated_at": "2012-09-14T11:16:51Z",
            "branches": [
               {
                  "branch_name": "mongoid3",
                  "branch_url": "https://semaphoreapp.com/projects/63/branches/89",
                  "branch_status_url": "https://semaphoreapp.com/api/v1/projects/3f1004b8343faabda63d441734526c854380ab89/85/status?auth_token=Yds3w6o26FLfJTnVK2y9",
                  "branch_history_url": "https://semaphoreapp.com/api/v1/projects/3f1004b8343faabda63d441734526c854380ab89/85?auth_token=Yds3w6o26FLfJTnVK2y9",
                  "project_name": "testapp-mongodb",
                  "build_url": "https://semaphoreapp.com/projects/63/branches/89/builds/3",
                  "build_info_url": "https://semaphoreapp.com/api/v1/projects/3f1004b8343faabda63d441734526c854380ab89/85/builds/1?auth_token=Yds3w6o26FLfJTnVK2y9",
                  "build_number": 3,
                  "result": "passed",
                  "started_at": "2012-09-14T11:11:39Z",
                  "finished_at": "2012-09-14T11:16:51Z"
               }
            ],
            "servers": []
         }
      ]`)
	defer server.Close()

	// Test the method!
	actual, _ := client.GetProjects()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Response did not match expected \nA: %v\n\n E: %v\n\n", actual, expected)
	}
}
