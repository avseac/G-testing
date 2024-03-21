package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"GinGonic-Tmpl/api"
	"GinGonic-Tmpl/router"

	"github.com/gin-gonic/gin"
)

var url string = "127.0.0.1:8000"

func sum(a int) int {
	return a + 1
}

func main() {
	fmt.Println("Server online a: ", "http://"+url)
	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()

	app.LoadHTMLGlob("templates/*.html")
	app.SetFuncMap(template.FuncMap{
		"sum": sum,
	})
	// app.LoadHTMLGlob("components/*.html")

	app.Static("/static", "./static/")
	app.Static("/component", "./components/")
	app.GET("/menu", router.ServeMenu)
	app.GET("/products", router.ServeProducts)
	app.GET("/order", router.ServeProductsForOrder)

	app.GET("/", func(c *gin.Context) {
		// c.Redirect(http.StatusTemporaryRedirect, "/products")
		c.HTML(http.StatusOK, "tmp.html", gin.H{
			"a": 1,
			"b": 4,
		})
	})

	app.GET("/data/:page", func(c *gin.Context) {
		val, err := strconv.Atoi(c.Param("page"))
		if err != nil {
			panic(err)
		}

		c.HTML(http.StatusOK, "row.html", gin.H{
			"row": api.DataProductList[val],
		})

	})

	app.GET("/test", func(c *gin.Context) {
		c.HTML(http.StatusOK, "test.html", gin.H{
			"title": "Testing",
			"data":  api.DataProductList,
		})
	})

	app.Run(url)
}
