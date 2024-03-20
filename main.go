package main

import (
	"fmt"
	"net/http"

	"GinGonic-Tmpl/router"

	"github.com/gin-gonic/gin"
)

var url string = "127.0.0.1:8000"

func main() {
	fmt.Println("Server online a: ", "http://"+url)
	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()

	app.LoadHTMLGlob("templates/*.html")

	app.Static("/static", "./static/")
	app.GET("/menu", router.ServeMenu)
	app.GET("/products", router.ServeProducts)

	app.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/products")
	})

	// app.GET("/", router.ServeIndex)
	// app.GET("/products", router.ServeProducts)
	// app.GET("/order", router.ServeOrder)
	// app.GET("/orders", router.ServeOrders)

	app.Run(url)
}
