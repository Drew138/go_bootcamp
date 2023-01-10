package main

import (
	"dia_2/cmd/handlers"

	"github.com/gin-gonic/gin"
)

func registerHandlers(r *gin.Engine) {
	r.GET("/ping", handlers.Ping)
	r.GET("/product/:id", handlers.RetrieveProduct)
	r.GET("/product", handlers.GetProducts)
	r.GET("/product/search", handlers.FilterProductsByPrice)
	r.POST("/product", handlers.CreateProduct)
}

func main() {
	r := gin.Default()
	registerHandlers(r)
	r.Run()
}
