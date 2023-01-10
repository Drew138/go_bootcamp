package handlers

import (
	"dia_2/pkg/response"
	"dia_2/services/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func RetrieveProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = fmt.Errorf("invalid id, %s", err)
		c.JSON(http.StatusBadRequest, response.Err(err))
	}
	product, err := models.SearchProductById(id)
	if err != nil {
		err = fmt.Errorf("%s, %s", response.ErrNotFound, err)
		c.JSON(http.StatusNotFound, response.Err(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok("Product retrieved successfully", &product))
}

func GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, response.Ok("products retrieved successfully", models.Products))
}

func FilterProductsByPrice(c *gin.Context) {
	priceGt, err := strconv.Atoi(c.Param("priceGt"))
	if err != nil {
		err = fmt.Errorf("invalid id, %s", err)
		c.JSON(http.StatusBadRequest, response.Err(err))
	}
	var products = models.FilterProductByPrice(float64(priceGt))
	c.JSON(http.StatusOK, response.Ok("products retrieved successfully", products))
}

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBind(&product); err != nil {
		c.JSON(http.StatusBadRequest, response.Err(err))
		return
	}
	if err := models.SaveProduct(&product); err != nil {
		c.JSON(http.StatusBadRequest, response.Err(err))
		return
	}
	c.JSON(http.StatusCreated, response.Ok("product created successfully", product))
}
