package verifysphere_sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginResponse struct {
	Token string `json:"token"`
}

func (c *Client) Login(username, password string) (string, error) {
	// 构建登录请求的URL
	loginURL := fmt.Sprintf("%s/login", c.baseURL)

	// 构建请求体
	requestBody := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{
		Username: username,
		Password: password,
	}

	// 序列化请求体
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	// 发送POST请求
	resp, err := http.Post(loginURL, "application/json", bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 解析响应
	var loginResponse LoginResponse
	if err := json.NewDecoder(resp.Body).Decode(&loginResponse); err != nil {
		return "", err
	}

	return loginResponse.Token, nil
}
