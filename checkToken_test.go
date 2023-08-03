package verifysphere_sdk_test

import (
	"encoding/json"
	"fmt"
	"github.com/guilinonline/verifysphere-sdk"
	"testing"
)

func TestLogin(t *testing.T) {
	baseURL := "http://121.36.148.107:9090/api/v1"
	client := verifysphere_sdk.NewClient(baseURL)
	ctx := map[string]string{
		"path":   "/api/v1/idc/get_vlan_detail",
		"method": "GET",
	}
	// 测试用例1：正确的用户名和密码
	ctxJSON, err := json.Marshal(ctx)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	vp := verifysphere_sdk.Request{
		Pattern: "midd",
		Server:  "a-626c10330c580",
		Context: string(ctxJSON),
	}
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVSUQiOiJ1LTYyNmU1NTI5YmY1ODAiLCJVc2VybmFtZSI6InRlc3QxIiwiVXNlclR5cGUiOiJub3JtYWwiLCJBdWRpZW5jZSI6InNlbGYiLCJpc3MiOiJWZXJpZnlTcGhlcmUiLCJzdWIiOiJ0ZXN0MSIsImV4cCI6MTY5MTA2MDI5NSwibmJmIjoxNjkxMDMxNDk1LCJpYXQiOjE2OTEwMzE0OTV9.pbA9JjH9djYmf-5gzXVhp3YtwM99z73YnnojWeGx6so"
	data, code, err := client.CheckToken(vp, token)
	if err != nil {
		t.Errorf("check failed for  %s", err.Error())
	}
	fmt.Println(string(data), code)

}
