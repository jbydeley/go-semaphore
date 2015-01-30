package semaphore

import (
	"reflect"
	"testing"
)

func Test_GetBranches_Success(t *testing.T) {
	expected := []Branch{
		{1324, "new-build-page", "https://semaphoreapp.com/api/v1/projects/:hash_ID/1324/status?auth_token=:auth_token"},
		{1120, "development", "https://semaphoreapp.com/api/v1/projects/:hash_ID/1120/status?auth_token=:auth_token"},
		{987, "branches_api", "https://semaphoreapp.com/api/v1/projects/:hash_ID/987/status?auth_token=:auth_token"},
	}
	server, client := testAPICall(200, `
      [
         {
            "ID": 1324,
            "name": "new-build-page",
            "branch_URL": "https://semaphoreapp.com/api/v1/projects/:hash_ID/1324/status?auth_token=:auth_token"
         },
         {
            "ID": 1120,
            "name": "development",
            "branch_URL": "https://semaphoreapp.com/api/v1/projects/:hash_ID/1120/status?auth_token=:auth_token"
         },
         {
            "ID": 987,
            "name": "branches_api",
            "branch_URL": "https://semaphoreapp.com/api/v1/projects/:hash_ID/987/status?auth_token=:auth_token"
         }
      ]`)
	defer server.Close()

	// Test the method!
	actual, _ := client.GetBranches("123")

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Response dID not match expected \nA: %v\n\n E: %v\n\n", actual, expected)
	}
}

func Test_GetBranches_NotFound(t *testing.T) {
	server, client := testAPICall(404, "")
	defer server.Close()

	results, err := client.GetBranches("123")

	if err != ErrHTTPNotFound {
		t.Errorf("Should have recieved 404 but got '%v'", err)
	}

	if results != nil {
		t.Errorf("Results should be nil but got '%+v'", results)
	}
}

func Test_GetBranches_NotRecognized(t *testing.T) {
	server, client := testAPICall(200, "???")
	defer server.Close()

	results, err := client.GetBranches("123")

	if err != ErrResponseNotRecognized {
		t.Errorf("Should have received an ErrResponseNotRecognized but got '%v'", err)
	}

	if results != nil {
		t.Errorf("Results should be nil but got '%+v'", results)
	}
}

func Test_GetBranches_FailedAuthentication(t *testing.T) {
	server, client := testAPICall(401, "")
	defer server.Close()

	results, err := client.GetBranches("123")

	if err != ErrNotAuthorized {
		t.Errorf("Should have recieved 401 but got '%v'", err)
	}

	if results != nil {
		t.Errorf("Results should be nil but got '%+v'", results)
	}
}
