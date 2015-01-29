package semaphore

import (
	"reflect"
	"testing"
)

func Test_GetBranches_Success(t *testing.T) {
	expected := []Branch{
		{1324, "new-build-page", "https://semaphoreapp.com/api/v1/projects/:hash_id/1324/status?auth_token=:auth_token"},
		{1120, "development", "https://semaphoreapp.com/api/v1/projects/:hash_id/1120/status?auth_token=:auth_token"},
		{987, "branches_api", "https://semaphoreapp.com/api/v1/projects/:hash_id/987/status?auth_token=:auth_token"},
	}
	server, client := testApiCall(200, `
      [
         {
            "id": 1324,
            "name": "new-build-page",
            "branch_url": "https://semaphoreapp.com/api/v1/projects/:hash_id/1324/status?auth_token=:auth_token"
         },
         {
            "id": 1120,
            "name": "development",
            "branch_url": "https://semaphoreapp.com/api/v1/projects/:hash_id/1120/status?auth_token=:auth_token"
         },
         {
            "id": 987,
            "name": "branches_api",
            "branch_url": "https://semaphoreapp.com/api/v1/projects/:hash_id/987/status?auth_token=:auth_token"
         }
      ]`)
	defer server.Close()

	// Test the method!
	actual, _ := client.GetBranches("123")

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Response did not match expected \nA: %v\n\n E: %v\n\n", actual, expected)
	}
}
