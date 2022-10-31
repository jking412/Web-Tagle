package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HomeController struct {
}

func (hc *HomeController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "home", gin.H{})
}
