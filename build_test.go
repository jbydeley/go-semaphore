package semaphore

import (
	"reflect"
	"testing"
)

func Test_RebuildLastRevision_Success(t *testing.T) {
	expected := &BuildResponse{
		[]Commit{
			Commit{
				"ee89ebaaaeasdaasdasdqwewlweqlqwleqlwe",
				"https://github.com/renderedtext/semaphore/commit/dasadsdasadsdasadsdsaasdasdasd",
				"Rastko Jokic",
				"rastko@renderedtext.com",
				"Add cucumber_in_groups",
				"2013-12-25T11:57:44+01:00",
			},
			Commit{
				"2a6e8df4llllll11427e1asdasl97506ffac15",
				"https://github.com/renderedtext/semaphore/commit/2a6e8dasddasdsasadaf69753d5d06ffac15",
				"Marko Anastasov",
				"marko@renderedtext.com",
				"Merge pull request #410 from renderedtext/rj/cucumber-groups\n\nAdd cucumber_in_groups",
				"2013-12-25T12:31:07+01:00",
			},
		},
		"semaphore",
		"development",
		1,
		"",
		"2013-12-25T14:56:43+01:00",
		"2013-12-25T14:56:44+01:00",
		"",
		"",
		"https://semaphoreapp.com/darkofabijan/semaphore/branches/development/builds/1",
		"https://semaphoreapp.com/api/v1/projects/73c4b979-0a40-49db-b10e-571d41e10d9a/133529/builds/1/log?auth_token=:auth_token",
		"https://semaphoreapp.com/api/v1/projects/73c4b979-0a40-49db-b10e-571d41e10d9a/133529/builds/1?auth_token=:auth_token",
	}

	server, client := testAPICall(200, `{
  "commits": [
    {
      "id": "ee89ebaaaeasdaasdasdqwewlweqlqwleqlwe",
      "url": "https://github.com/renderedtext/semaphore/commit/dasadsdasadsdasadsdsaasdasdasd",
      "author_name": "Rastko Jokic",
      "author_email": "rastko@renderedtext.com",
      "message": "Add cucumber_in_groups",
      "timestamp": "2013-12-25T11:57:44+01:00"
    },
    {
      "id": "2a6e8df4llllll11427e1asdasl97506ffac15",
      "url": "https://github.com/renderedtext/semaphore/commit/2a6e8dasddasdsasadaf69753d5d06ffac15",
      "author_name": "Marko Anastasov",
      "author_email": "marko@renderedtext.com",
      "message": "Merge pull request #410 from renderedtext/rj/cucumber-groups\n\nAdd cucumber_in_groups",
      "timestamp": "2013-12-25T12:31:07+01:00"
    }
  ],
  "project_name": "semaphore",
  "branch_name": "development",
  "number": 1,
  "result": null,
  "created_at": "2013-12-25T14:56:43+01:00",
  "updated_at": "2013-12-25T14:56:44+01:00",
  "started_at": null,
  "finished_at": null,
  "html_url": "https://semaphoreapp.com/darkofabijan/semaphore/branches/development/builds/1",
  "build_log_url": "https://semaphoreapp.com/api/v1/projects/73c4b979-0a40-49db-b10e-571d41e10d9a/133529/builds/1/log?auth_token=:auth_token",
  "build_info_url": "https://semaphoreapp.com/api/v1/projects/73c4b979-0a40-49db-b10e-571d41e10d9a/133529/builds/1?auth_token=:auth_token"
}`)
	defer server.Close()

	actual, err := client.RebuildBranch("123", 456)

	if err != nil {
		t.Errorf("Error should be nil but was %v", err.Error())
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Response did not match expected \nA: %v\n\n E: %v\n\n", actual, expected)
	}
}

func Test_StopBuild_Success(t *testing.T) {
	expected := &BuildResponse{
		[]Commit{
			Commit{
				"ee89ebaaaeasdaasdasdqwewlweqlqwleqlwe",
				"https://github.com/renderedtext/semaphore/commit/dasadsdasadsdasadsdsaasdasdasd",
				"Rastko Jokic",
				"rastko@renderedtext.com",
				"Add cucumber_in_groups",
				"2013-12-25T11:57:44+01:00",
			},
			Commit{
				"2a6e8df4llllll11427e1asdasl97506ffac15",
				"https://github.com/renderedtext/semaphore/commit/2a6e8dasddasdsasadaf69753d5d06ffac15",
				"Marko Anastasov",
				"marko@renderedtext.com",
				"Merge pull request #410 from renderedtext/rj/cucumber-groups\n\nAdd cucumber_in_groups",
				"2013-12-25T12:31:07+01:00",
			},
		},
		"semaphore",
		"development",
		1,
		"",
		"2013-12-25T14:56:43+01:00",
		"2013-12-25T14:56:44+01:00",
		"",
		"",
		"https://semaphoreapp.com/darkofabijan/semaphore/branches/development/builds/1",
		"https://semaphoreapp.com/api/v1/projects/73c4b979-0a40-49db-b10e-571d41e10d9a/133529/builds/1/log?auth_token=:auth_token",
		"https://semaphoreapp.com/api/v1/projects/73c4b979-0a40-49db-b10e-571d41e10d9a/133529/builds/1?auth_token=:auth_token",
	}

	server, client := testAPICall(200, `{
  "commits": [
    {
      "id": "ee89ebaaaeasdaasdasdqwewlweqlqwleqlwe",
      "url": "https://github.com/renderedtext/semaphore/commit/dasadsdasadsdasadsdsaasdasdasd",
      "author_name": "Rastko Jokic",
      "author_email": "rastko@renderedtext.com",
      "message": "Add cucumber_in_groups",
      "timestamp": "2013-12-25T11:57:44+01:00"
    },
    {
      "id": "2a6e8df4llllll11427e1asdasl97506ffac15",
      "url": "https://github.com/renderedtext/semaphore/commit/2a6e8dasddasdsasadaf69753d5d06ffac15",
      "author_name": "Marko Anastasov",
      "author_email": "marko@renderedtext.com",
      "message": "Merge pull request #410 from renderedtext/rj/cucumber-groups\n\nAdd cucumber_in_groups",
      "timestamp": "2013-12-25T12:31:07+01:00"
    }
  ],
  "project_name": "semaphore",
  "branch_name": "development",
  "number": 1,
  "result": null,
  "created_at": "2013-12-25T14:56:43+01:00",
  "updated_at": "2013-12-25T14:56:44+01:00",
  "started_at": null,
  "finished_at": null,
  "html_url": "https://semaphoreapp.com/darkofabijan/semaphore/branches/development/builds/1",
  "build_log_url": "https://semaphoreapp.com/api/v1/projects/73c4b979-0a40-49db-b10e-571d41e10d9a/133529/builds/1/log?auth_token=:auth_token",
  "build_info_url": "https://semaphoreapp.com/api/v1/projects/73c4b979-0a40-49db-b10e-571d41e10d9a/133529/builds/1?auth_token=:auth_token"
}`)
	defer server.Close()

	actual, err := client.StopBuild("123", 456, 789)

	if err != nil {
		t.Errorf("Error should be nil but was %v", err.Error())
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Response did not match expected \nA: %v\n\n E: %v\n\n", actual, expected)
	}
}
