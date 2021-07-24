package model

import "github.com/satori/go.uuid"

type Product struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

type Products struct {
	Product []Product
}

func (receiver Products) Add(product Product)  {
	receiver.Product = append(receiver.Product, product)
}

func NewProduct() *Product  {
	return &Product{
		ID: uuid.NewV4().String(),
	}
}