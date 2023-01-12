package handlers

import (
	"dia_2/internal/domain"
	"dia_2/internal/product"
	"dia_2/pkg/response"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var ()

type ProductHandlerManager struct {
	sv product.Service
}

func NewProductHandlerManager(sv product.Service) *ProductHandlerManager {
	return &ProductHandlerManager{sv: sv}
}

func (pr *ProductHandlerManager) Ping() func(*gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}
}

func (pr *ProductHandlerManager) GetById() func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusBadRequest, response.Err(response.ErrInvalidId))
		}
		product, err := pr.sv.GetByID(id)
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusNotFound, response.Err(response.ErrNotFound))
			return
		}
		c.JSON(http.StatusOK, response.Ok("product retrieved successfully", &product))
	}
}

func parseFloat64Param(key string, c *gin.Context) (*float64, error) {
	if val := c.Param(key); val != "" {
		parsedValue, err := strconv.ParseFloat(val, 64)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		return &parsedValue, nil
	}
	return nil, nil
}

func (pr *ProductHandlerManager) Get() func(*gin.Context) {
	return func(c *gin.Context) {
		minPrice, err := parseFloat64Param("priceGt", c)
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusBadRequest, response.Err(errors.New("could not parse priceGt")))
			return
		}
		maxPrice, err := parseFloat64Param("priceLt", c)
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusBadRequest, response.Err(errors.New("could not parse priceLt")))
		}
		products, err := pr.sv.Get(minPrice, maxPrice)
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, response.Err(response.ErrInternalServerError))
			return
		}
		c.JSON(http.StatusOK, response.Ok("products retrieved successfully", products))
	}
}

func (pr *ProductHandlerManager) Create() func(*gin.Context) {
	return func(c *gin.Context) {
		var product domain.Product
		if err := c.ShouldBind(&product); err != nil {
			log.Fatal(err)
			c.JSON(http.StatusBadRequest, response.Err(errors.New("could not create a new product")))
			return
		}
		product, err := pr.sv.Create(product)
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusBadRequest, response.Err(err))
			return
		}
		c.JSON(http.StatusCreated, response.Ok("product created successfully", product))
	}
}

func (pr *ProductHandlerManager) Delete() func(*gin.Context) {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusBadRequest, response.Err(response.ErrInvalidId))
		}
		err = pr.sv.Delete(id)
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusBadRequest, response.Err(err))
			return
		}
		c.JSON(http.StatusNoContent, response.Ok("product deleted successfully", nil))
	}
}

func (pr *ProductHandlerManager) Put() func(*gin.Context) {
	return func(c *gin.Context) {
		var product domain.Product
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusBadRequest, response.Err(response.ErrInvalidId))
		}
		if err = c.ShouldBind(&product); err != nil {
			log.Fatal(err)
			c.JSON(http.StatusBadRequest, response.Err(errors.New("could not update product")))
			return
		}
		product, err = pr.sv.Update(product, id)
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, response.Err(response.ErrInternalServerError))
			return
		}
		c.JSON(http.StatusCreated, response.Ok("product updated successfully", product))
	}
}

func (pr *ProductHandlerManager) Patch() func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusBadRequest, response.Err(response.ErrInvalidId))
		}
		product, err := pr.sv.GetByID(id)
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusBadRequest, response.Err(errors.New("could not find product")))
			return
		}
		if err = c.ShouldBind(&product); err != nil {
			log.Fatal(err)
			c.JSON(http.StatusBadRequest, response.Err(errors.New("could not update product")))
			return
		}
		product, err = pr.sv.Update(product, id)
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, response.Err(response.ErrInternalServerError))
			return
		}
		c.JSON(http.StatusCreated, response.Ok("product updated successfully", product))
	}
}
