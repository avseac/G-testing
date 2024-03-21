package router

import (
	"net/http"

	"GinGonic-Tmpl/api"

	"github.com/gin-gonic/gin"
)

func ServeProducts(c *gin.Context) {

	c.HTML(http.StatusOK, "products.html", gin.H{
		"title": "Products list",
		"data":  api.DataProductList,
	})
}

func ServeProductsForOrder(c *gin.Context) {
	c.HTML(http.StatusOK, "order.html", gin.H{
		"title": "Order",
		"data":  api.DataProductList,
	})
}
