package main

import "fmt"

type Product struct {
	ID          int
	name        string
	price       float64
	description string
	category    string
}

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}
type Employee struct {
	Person
	DateOfBirth string
}

func (e *Employee) PrintEmployee() {
	fmt.Printf("%+v\n", e)
}

var products = []*Product{
	{
		1,
		"manzana",
		1.00,
		"una fruta",
		"fruta",
	},
	{
		2,
		"manzana",
		1.00,
		"una fruta",
		"fruta",
	},
}

func (p *Product) Save() {
	products = append(products, p)
}

func (p *Product) GetAll() {
	for _, product := range products {
		fmt.Printf("%+v\n", product)
	}
}

func getByID(id int) *Product {
	for _, product := range products {
		if product.ID == id {
			return product
		}
	}
	return nil
}

func appendSlice(p *[]*Product, pro *Product) {
	*(&(*(&(*p)))) = append(*p, pro)
}

func (p *Person) printPerson() {
	fmt.Printf("%+v\n", p)
}

// func main() {
// 	product := &Product{
// 		3,
// 		"mango",
// 		1.00,
// 		"una fruta",
// 		"fruta",
// 	}
// 	product.Save()
// 	product.GetAll()
// 	productByID := getByID(3)
// 	fmt.Printf("%+v", productByID)
// 	appendSlice(&products, product)
// 	// ============= ejercicio 2 =================
//
// 	person := Person{
// 		1,
// 		"Juan",
// 		"25/12/2022",
// 	}
// 	employee := Employee{
// 		person,
// 		"Desarrollador",
// 	}
// 	employee.PrintEmployee()
// 	person2 := &Person{
// 		1,
// 		"Juan",
// 		"25/12/2022",
// 	}
// 	person2.printPerson()
//
// }
