package conf

import "go-tagle/pkg/config"

var GithubClientConf = struct {
	ClientID     string
	ClientSecret string
}{}

func initGithubClientConf() {
	GithubClientConf.ClientID = config.LoadString("github.clientId", "")
	GithubClientConf.ClientSecret = config.LoadString("github.clientSecret", "")
}
