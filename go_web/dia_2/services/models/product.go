package models

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Quantity    int     `json:"quantity" validate:"required"`
	CodeValue   string  `json:"code_value" validate:"required"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
}

var Products []Product

func init() {
	file, err := os.Open("../products.json")
	if err != nil {
		log.Fatal(err)
		return
	}
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()
	json.Unmarshal(byteValue, &Products)
}

func SearchProductById(id int) (Product, error) {
	for _, product := range Products {
		if product.Id == id {
			return product, nil
		}
	}
	return Product{}, errors.New("product does not exist")
}

func FilterProductByPrice(price float64) []Product {
	var products = []Product{}
	for _, product := range Products {
		if product.Price == price {
			products = append(products, product)
		}
	}
	return products
}

func SaveProduct(newProduct *Product) error {
	tmp := 0
	for _, product := range Products {
		if product.CodeValue == newProduct.CodeValue {
			return errors.New("invalid code value")
		}
		if product.Id > tmp {
			tmp = product.Id
		}
	}
	newProduct.Id = tmp + 1
	_, err := time.Parse("01/02/2006", newProduct.Expiration)
	if err != nil {
		return err
	}
	Products = append(Products, *newProduct)
	return nil
}
