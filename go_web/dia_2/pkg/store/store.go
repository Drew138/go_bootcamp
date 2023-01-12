package store

import (
	"dia_2/internal/domain"
	"encoding/json"
	"io/ioutil"
	"os"
)

type Store interface {
	Read() ([]domain.Product, error)
	Save(*[]domain.Product) error
}

type ProductStore struct {
	file string
}

func NewProductStore(file string) Store {
	return &ProductStore{file}
}

func (p *ProductStore) Read() ([]domain.Product, error) {
	var products []domain.Product
	file, err := os.Open(p.file)
	if err != nil {
		return products, err
	}
	defer file.Close()
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return products, err
	}
	json.Unmarshal(byteValue, &products)
	return products, nil
}

func (p *ProductStore) Save(products *[]domain.Product) error {
	file, err := json.Marshal(*products)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(p.file, file, 0644)
	return err
}
