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
