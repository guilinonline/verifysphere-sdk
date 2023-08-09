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
type UserInfo struct {
	Code    int        `json:"code"`
	Data    []UserData `json:"data"`
	Message string     `json:"message"`
}

type UserData struct {
	ID         string      `json:"id"`
	Username   string      `json:"username"`
	Nickname   string      `json:"nickname"`
	State      int         `json:"state"`
	Phone      string      `json:"phone"`
	UserType   string      `json:"user_type"`
	Email      string      `json:"email"`
	Sort       int         `json:"sort"`
	Department interface{} `json:"department"`
	CreateTime int64       `json:"create_time"`
	UpdateTime int64       `json:"update_time"`
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
	//fmt.Println(resp.StatusCode, token)
	if resp.StatusCode == 200 {
		Flag = true
	} else if resp.StatusCode == 403 {
		Flag = false
	} else {
		return Flag, fmt.Errorf("鉴权服务异常，接口状态码为：%d", resp.StatusCode)
	}

	return Flag, nil
}

func (c *Client) GetUserInfo(token string) (UserInfo, error) {
	var userinfo UserInfo
	url := fmt.Sprintf("%s/self/", c.baseURL)
	req, err := http.NewRequest(
		"GET",
		url,
		nil,
	)
	if err != nil {
		return userinfo, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return userinfo, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return userinfo, err
	}
	if resp.StatusCode != 200 && resp.StatusCode != 403 {
		return userinfo, errors.New(fmt.Sprintf("HTTP状态异常,状态码%s,返回信息：%s",
			resp.StatusCode, userinfo.Message,
		))
	}

	err = json.Unmarshal(body, &userinfo)
	if err != nil {
		return userinfo, err
	}
	if userinfo.Code != 0 {
		return userinfo, errors.New(fmt.Sprintf("数据异常,状态码%s,返回信息：%s,返回Code：%d",
			resp.StatusCode, userinfo.Message, userinfo.Code,
		))
	}
	return userinfo, nil
}

func (c *Client) Logout(token string) (bool, error) {
	url := fmt.Sprintf("%s/logout/", c.baseURL)
	req, err := http.NewRequest(
		"POST",
		url,
		nil,
	)
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return false, errors.New("HTTP状态异常")
	}

	return true, nil
}
