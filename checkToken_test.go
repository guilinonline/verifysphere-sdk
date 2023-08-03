package verifysphere_sdk_test

import (
	"testing"
)

func TestLogin(t *testing.T) {
	baseURL := "https://api.example.com"
	client := mysdk.NewClient(baseURL)

	// 测试用例1：正确的用户名和密码
	username := "correct_username"
	password := "correct_password"

	token, err := client.Login(username, password)
	if err != nil {
		t.Errorf("Login failed for username '%s': %s", username, err.Error())
	}

	// 测试用例2：错误的用户名和密码
	username = "incorrect_username"
	password = "incorrect_password"

	_, err = client.Login(username, password)
	if err == nil {
		t.Errorf("Login succeeded for invalid username '%s' and password '%s'", username, password)
	}
}
