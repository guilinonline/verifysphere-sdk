package verifysphere_sdk

import (
	"fmt"

	"testing"
)

//func TestLogin(t *testing.T) {
//	baseURL := "http://172.25.162.56:9090/api/v1"
//	client := NewClient(baseURL)
//
//	ctx := map[string]string{
//		"path":   "/api/v1/user/search/",
//		"method": "POST",
//	}
//	ctxJSON, err := json.Marshal(ctx)
//	if err != nil {
//		fmt.Println("Error:", err)
//		return
//	}
//
//	vp := Request{
//		Pattern: "midd",
//		Server:  "VerifySphere",
//		Context: string(ctxJSON),
//	}
//	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVSUQiOiJ1LTYyNzYyNjZiYzAxODAiLCJVc2VybmFtZSI6Imt0eHgwMDI3IiwiVXNlclR5cGUiOiJsZGFwIiwiQXVkaWVuY2UiOiJWZXJpZnlTcGhlcmUiLCJpc3MiOiJWZXJpZnlTcGhlcmUiLCJzdWIiOiJrdHh4MDAyNyIsImV4cCI6MTY5MTU5NzMzMiwibmJmIjoxNjkxNTY4NTMyLCJpYXQiOjE2OTE1Njg1MzJ9.s-UPfiYAD6dijhks_pLEjbx8nN5X4V5bpPbyHlpBn30"
//	data, err := client.CheckToken(vp, token)
//	if err != nil {
//		t.Errorf("check failed for  %s", err.Error())
//	}
//	fmt.Println(data)
//
//}
//
//func TestUserinfo(t *testing.T) {
//	baseURL := "http://172.25.162.56:9090/api/v1"
//	client := NewClient(baseURL)
//	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVSUQiOiJ1LTYyNzYyNjZiYzAxODAiLCJVc2VybmFtZSI6Imt0eHgwMDI3IiwiVXNlclR5cGUiOiJsZGFwIiwiQXVkaWVuY2UiOiJWZXJpZnlTcGhlcmUiLCJpc3MiOiJWZXJpZnlTcGhlcmUiLCJzdWIiOiJrdHh4MDAyNyIsImV4cCI6MTY5MTU5NzMzMiwibmJmIjoxNjkxNTY4NTMyLCJpYXQiOjE2OTE1Njg1MzJ9.s-UPfiYAD6dijhks_pLEjbx8nN5X4V5bpPbyHlpBn30"
//
//	data, err := client.GetUserInfo(token)
//	if err != nil {
//		t.Errorf("check failed for  %s", err.Error())
//	}
//	fmt.Printf("结果：%#v", data)
//}
//
//func TestLogout(t *testing.T) {
//	baseURL := "http://172.25.162.56:9090/api/v1"
//	client := NewClient(baseURL)
//	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVSUQiOiJ1LTYyNzYyNjZiYzAxODAiLCJVc2VybmFtZSI6Imt0eHgwMDI3IiwiVXNlclR5cGUiOiJsZGFwIiwiQXVkaWVuY2UiOiJWZXJpZnlTcGhlcmUiLCJpc3MiOiJWZXJpZnlTcGhlcmUiLCJzdWIiOiJrdHh4MDAyNyIsImV4cCI6MTY5MTU5NzI0NiwibmJmIjoxNjkxNTY4NDQ2LCJpYXQiOjE2OTE1Njg0NDZ9.vpspHW_zWg3T197_0a1GFPlcSf4YMzSacczEFs4Yaxs"
//
//	data, err := client.Logout(token)
//	if err != nil {
//		t.Errorf("check failed for  %s", err.Error())
//	}
//	fmt.Printf("结果：%#v", data)
//}

func TestPing(t *testing.T) {
	baseURL := "http://172.25.162.56:9090/api/v1"
	client := NewClient(baseURL)
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVSUQiOiJ1LTYyNzYyNjZiYzAxODAiLCJVc2VybmFtZSI6Imt0eHgwMDI3IiwiVXNlclR5cGUiOiJsZGFwIiwiQXVkaWVuY2UiOiJWZXJpZnlTcGhlcmUiLCJpc3MiOiJWZXJpZnlTcGhlcmUiLCJzdWIiOiJrdHh4MDAyNyIsImV4cCI6MTY5MTg0NjgyNCwibmJmIjoxNjkxODE4MDI0LCJpYXQiOjE2OTE4MTgwMjR9.cgOGbyMdz8RfxZmidtneltftxODenLgo-9W7mDCUdss"

	data, err := client.Ping(token)
	if err != nil {
		t.Errorf("check failed for  %s", err.Error())
	}
	fmt.Printf("结果：%#v", data)
}
