package template

import (
	"github.com/gin-gonic/gin"
	"go-tagle/conf"
	"go-tagle/pkg/logger"
	"html/template"
)

func InitTemplate(r *gin.Engine) {
	r.LoadHTMLGlob(conf.TemplateConf.TemplateDir)
	r.SetFuncMap(template.FuncMap{})

	logger.InfoString("template", "模板初始化成功", "")
	//r.Static("/static", viperlib.GetString("static.path"))
}
