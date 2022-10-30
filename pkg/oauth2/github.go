package oauth2

import (
	"encoding/json"
	"fmt"
	"go-tagle/conf"
	"go-tagle/pkg/helper"
	"go-tagle/pkg/logger"
	"io/ioutil"
	"net/http"
	"strings"
)

type Token struct {
	AccessToken string
}

type GithubUserInfo struct {
	Username  string `json:"login"`
	AvatarUrl string `json:"avatar_url"`
}

func GetGithubUserInfo(code string) (*GithubUserInfo, bool) {
	resp, err := http.Post("https://github.com/login/oauth/access_token",
		"application/x-www-form-urlencoded",
		strings.NewReader(fmt.Sprintf("client_id=%s&client_secret=%s&code=%s",
			conf.GithubClientConf.ClientID,
			conf.GithubClientConf.ClientSecret,
			code)))
	if err != nil {
		logger.WarnString("oauth2", "获取github access_token失败", err.Error())
		return nil, false
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var token Token
	form := helper.ParseForm(string(body))
	for k, v := range form {
		if k == "access_token" {
			token.AccessToken = v
		}
	}
	req, _ := http.NewRequest("GET", "https://api.github.com/user", nil)
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		logger.WarnString("oauth2", "获取github用户信息失败", err.Error())
		return nil, false
	}
	defer resp.Body.Close()
	body, _ = ioutil.ReadAll(resp.Body)
	githubUserInfo := &GithubUserInfo{}
	if err = json.Unmarshal(body, githubUserInfo); err != nil {
		logger.WarnString("oauth2", "获取github用户信息失败", err.Error())
		return nil, false
	}
	return githubUserInfo, true
}
