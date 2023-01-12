package product

import (
	"dia_2/internal/domain"
	"dia_2/pkg/response"
	"time"
)

// controller
type Service interface {
	Get(*float64, *float64) ([]domain.Product, error)
	GetByID(id int) (domain.Product, error)
	Create(domain.Product) (domain.Product, error)
	Delete(id int) error
	Update(domain.Product, int) (domain.Product, error)
}

type ProductService struct {
	rp Repository
}

func NewService(rp Repository) Service {
	return &ProductService{rp: rp}
}

func (sv *ProductService) Get(minPrice, maxPrice *float64) ([]domain.Product, error) {
	return sv.rp.Get(minPrice, maxPrice)
}

func (sv *ProductService) GetByID(id int) (domain.Product, error) {
	return sv.rp.GetByID(id)
}

func (sv *ProductService) Create(product domain.Product) (domain.Product, error) {
	exists, err := sv.rp.ExistsCodeValue(product.CodeValue)
	if err != nil {
		return domain.Product{}, err
	}
	if exists {
		return domain.Product{}, response.ErrAlreadyExist
	}
	_, err = time.Parse("01/02/2006", product.Expiration)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (sv *ProductService) Delete(id int) error {
	return sv.rp.Delete(id)
}

func (sv *ProductService) Update(updatedProduct domain.Product, id int) (domain.Product, error) {
	return sv.rp.Update(updatedProduct, id)
}
