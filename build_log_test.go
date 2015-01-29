package semaphore

import (
	"reflect"
	"testing"
)

func Test_GetBuildLog_Success(t *testing.T) {
	expected := &BuildLog{
		[]Thread{
			{
				1,
				[]Command{
					{
						"bundle install --deployment --path vendor/bundle",
						"0",
						"Fetching source index from http://rubygems.org/ Installing rake (10.0.3) Installing i18n (0.6.1) Installing multi_json (1.6.1) Installing activesupport (3.2.11) Installing builder (3.0.4) Installing activemodel (3.2.11) Installing erubis (2.7.0) Installing journey (1.0.4) Installing rack (1.4.5) Installing rack-cache (1.2) Installing rack-test (0.6.2) Installing hike (1.2.1) Installing tilt (1.3.3) Installing sprockets (2.2.2) Installing actionpack (3.2.11) Installing mime-types (1.21) Installing polyglot (0.3.3) Installing treetop (1.4.12) Installing mail (2.4.4) Installing actionmailer (3.2.11) Installing arel (3.0.2) Installing tzinfo (0.3.35) Installing activerecord (3.2.11) Installing activeresource (3.2.11) Installing addressable (2.3.3) Installing kaminari (0.14.1) Installing polyamorous (0.5.0) Installing meta_search (1.1.3) Using bundler (1.3.0) Installing rack-ssl (1.3.2) Installing json (1.7.7)",
						"2013-03-12T09:24:30Z",
						"2013-03-12T09:24:37Z",
						"00:45",
					},
					{
						"gem --version",
						"0",
						"1.8.23",
						"2013-03-12T09:25:30Z",
						"2013-03-12T09:25:37Z",
						"00:00",
					},
				},
			},
			{
				2,
				[]Command{
					{
						"bundle exec rake db:test:prepare",
						"0",
						"",
						"2013-03-12T09:24:37Z",
						"2013-03-12T09:24:44Z",
						"00:07",
					},
				},
			},
		},
		"https://semaphoreapp.com/api/v1/projects/3f1004b8343faabda63d441734526c854380ab89/85/builds/1?auth_token=Yds3w6o26FLfJTnVK2y9",
	}
	server, client := testAPICall(200, `
      {
        "threads": [
          {
            "number": 1,
            "commands": [
              {
                "name": "bundle install --deployment --path vendor/bundle",
                "result": "0",
                "output": "Fetching source index from http://rubygems.org/ Installing rake (10.0.3) Installing i18n (0.6.1) Installing multi_json (1.6.1) Installing activesupport (3.2.11) Installing builder (3.0.4) Installing activemodel (3.2.11) Installing erubis (2.7.0) Installing journey (1.0.4) Installing rack (1.4.5) Installing rack-cache (1.2) Installing rack-test (0.6.2) Installing hike (1.2.1) Installing tilt (1.3.3) Installing sprockets (2.2.2) Installing actionpack (3.2.11) Installing mime-types (1.21) Installing polyglot (0.3.3) Installing treetop (1.4.12) Installing mail (2.4.4) Installing actionmailer (3.2.11) Installing arel (3.0.2) Installing tzinfo (0.3.35) Installing activerecord (3.2.11) Installing activeresource (3.2.11) Installing addressable (2.3.3) Installing kaminari (0.14.1) Installing polyamorous (0.5.0) Installing meta_search (1.1.3) Using bundler (1.3.0) Installing rack-ssl (1.3.2) Installing json (1.7.7)",
                "start_time": "2013-03-12T09:24:30Z",
                "finish_time": "2013-03-12T09:24:37Z",
                "duration": "00:45"
              },
              {
                "name": "gem --version",
                "result": "0",
                "output": "1.8.23",
                "start_time": "2013-03-12T09:25:30Z",
                "finish_time": "2013-03-12T09:25:37Z",
                "duration": "00:00"
              }
            ]
          },
          {
            "number": 2,
            "commands": [
              {
                "name": "bundle exec rake db:test:prepare",
                "result": "0",
                "output": "",
                "start_time": "2013-03-12T09:24:37Z",
                "finish_time": "2013-03-12T09:24:44Z",
                "duration": "00:07"
              }
            ]
          }
        ],
        "build_info_URL": "https://semaphoreapp.com/api/v1/projects/3f1004b8343faabda63d441734526c854380ab89/85/builds/1?auth_token=Yds3w6o26FLfJTnVK2y9"
      }`)
	defer server.Close()

	// Test the method!
	actual, _ := client.GetBuildLog("hash_id", 1234, 4321)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Response did not match expected \nA: %v\n\n E: %v\n\n", actual, expected)
	}
}
