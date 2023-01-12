package product

import (
	"dia_2/internal/domain"
	"dia_2/pkg/store"
	"errors"
)

var (
	ErrProductNotFound = errors.New("product not found")
)

type Repository interface {
	Get(*float64, *float64) ([]domain.Product, error)
	GetByID(int) (domain.Product, error)
	ExistsCodeValue(codeValue string) (bool, error)
	Delete(id int) error
	Update(domain.Product, int) (domain.Product, error)
	Create(domain.Product) (domain.Product, error)
}

func NewRepository(st store.Store) Repository {
	return &ProductRepository{st}
}

type ProductRepository struct {
	st store.Store
}

func (r *ProductRepository) GetByID(id int) (domain.Product, error) {
	products, err := r.st.Read()
	if err != nil {
		return domain.Product{}, err
	}
	for _, product := range products {
		if product.Id == id {
			return product, nil
		}
	}
	return domain.Product{}, ErrProductNotFound
}

func (r *ProductRepository) Get(minPrice, maxPrice *float64) ([]domain.Product, error) {
	var products = []domain.Product{}
	dbProducts, err := r.st.Read()
	if err != nil {
		return products, err
	}
	for _, product := range dbProducts {
		if minPrice != nil && *minPrice > product.Price {
			continue
		}
		if maxPrice != nil && *maxPrice < product.Price {
			continue
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *ProductRepository) Create(product domain.Product) (domain.Product, error) {
	products, err := r.st.Read()
	id := r.nextId(&products)
	if err != nil {
		return domain.Product{}, err
	}
	product.Id = id
	products = append(products, product)
	return product, nil
}

func (r *ProductRepository) ExistsCodeValue(codeValue string) (bool, error) {

	products, err := r.st.Read()
	if err != nil {
		return false, err
	}
	for _, product := range products {
		if product.CodeValue == codeValue {
			return true, nil
		}
	}
	return false, nil
}

func (r *ProductRepository) nextId(products *[]domain.Product) int {
	tmp := 0
	for _, product := range *products {
		if product.Id > tmp {
			tmp = product.Id
		}
	}
	return tmp + 1
}

func (r *ProductRepository) Delete(id int) error {
	products, err := r.st.Read()
	if err != nil {
		return err
	}
	var newProducts = []domain.Product{}
	var found = false
	for _, product := range products {
		if product.Id == id {
			found = true
			continue
		}
		newProducts = append(newProducts, product)
	}
	if !found {
		return ErrProductNotFound
	}
	if err = r.st.Save(&newProducts); err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) Update(updatedProduct domain.Product, id int) (domain.Product, error) {
	products, err := r.st.Read()
	if err != nil {
		return domain.Product{}, err
	}
	var newProducts = []domain.Product{}
	for _, product := range products {
		if product.Id == id {
			product = updatedProduct
			break
		}
	}
	if err = r.st.Save(&newProducts); err != nil {
		return domain.Product{}, err
	}
	return updatedProduct, nil
}
