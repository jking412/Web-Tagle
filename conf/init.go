package conf

func InitConf() {
	initLogConf()
	initTemplateConf()
	initDBConf()
	initSessionConf()
	initRedisConf()
	initGithubClientConf()
	initServerConf()
	initEmailConf()
	initCaptchaConf()
}
