package verifysphere_sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Request struct {
	Pattern string `json:"pattern"`
	Server  string `json:"server"`
	Context string `json:"context"`
}

func (c *Client) CheckToken(data Request, token string) (Flag bool, err error) {
	// 构建登录请求的URL
	Flag = false
	loginURL := fmt.Sprintf("%s/oauth/", c.baseURL)

	// 序列化请求体
	requestBodyBytes, err := json.Marshal(data)
	if err != nil {
		return Flag, err
	}

	// 发送POST请求
	req, err := http.NewRequest(
		"POST",
		loginURL,
		bytes.NewBuffer(requestBodyBytes),
	)
	if err != nil {
		return Flag, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Flag, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		Flag = true
	} else if resp.StatusCode == 403 {
		Flag = false
	} else {
		return Flag, fmt.Errorf("鉴权服务异常，接口状态码为：%d", resp.StatusCode)
	}

	return Flag, nil
}
