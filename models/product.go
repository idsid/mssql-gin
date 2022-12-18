package models

import "fmt"

type Product struct {
	Id       int64
	Name     string
	Price    float64
	Quantity int
	Status   bool
}

type ProductInput struct {
	Name     string  `json:"name" binding:"required"`
	Price    float64 `json:"price" binding:"required"`
	Quantity int     `json:"quantity" binding:"required"`
	Status   bool    `json:"status" binding:"required"`
}

func (product Product) ToString() string {
	return fmt.Sprintf("Id: %d, Name: %s, Price: %f, Quantity: %d, Status: %t", product.Id, product.Name, product.Price, product.Quantity, product.Status)
}
