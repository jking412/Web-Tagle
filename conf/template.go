package conf

import "go-tagle/pkg/config"

var TemplateConf = struct {
	TemplateDir string
}{}

func initTemplateConf() {
	TemplateConf.TemplateDir = config.LoadString("template.dir", "./views/front/*.html")
}
