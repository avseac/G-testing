package router

import (
	"net/http"
	"time"

	"GinGonic-Tmpl/api"

	"github.com/gin-gonic/gin"
)

func ServeMenu(c *gin.Context) {
	t := time.Now()

	c.HTML(http.StatusOK, "menu.html", gin.H{
		"title": "Menu",
		"data":  api.DataProductGroupList,
		"today": t.Local().Format("02/01/2006"),
	})
}
