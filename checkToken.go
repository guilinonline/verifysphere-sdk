package verifysphere_sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Request struct {
	Pattern string `json:"pattern"`
	Server  string `json:"server"`
	Context string `json:"context"`
}

func (c *Client) CheckToken(data Request, token string) (resp_data []byte, http_code int, err error) {
	// 构建登录请求的URL
	loginURL := fmt.Sprintf("%s/oauth/", c.baseURL)

	// 序列化请求体
	requestBodyBytes, err := json.Marshal(data)
	if err != nil {
		return nil, 0, err
	}

	// 发送POST请求
	req, err := http.NewRequest(
		"POST",
		loginURL,
		bytes.NewBuffer(requestBodyBytes),
	)
	if err != nil {
		return nil, 0, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 && resp.StatusCode != 403 {
		return nil, resp.StatusCode, errors.New(fmt.Sprintf("鉴权服务异常，接口状态码为：%d", resp.StatusCode))
	}
	// 解析响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	return body, resp.StatusCode, nil
}
