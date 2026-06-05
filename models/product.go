/*
Autor: Luis Miguel Vasconez
Fecha: 27/05/2026
Descripcion: Incio desarrollo de sistema de gestion de Ecommerce con Go
*/

package models

import "errors"

type Product struct {
	id       int
	name     string
	price    float64
	stock    int
	category string
}

func NewProduct(id int, name string, price float64, stock int, category string) Product {
	return Product{
		id:       id,
		name:     name,
		price:    price,
		stock:    stock,
		category: category,
	}
}

// Getters
func (p Product) GetID() int {
	return p.id
}

func (p Product) GetName() string {
	return p.name
}

func (p Product) GetPrice() float64 {
	return p.price
}

func (p Product) GetStock() int {
	return p.stock
}

func (p Product) GetCategory() string {
	return p.category
}

// Setter con manejo de errores
func (p *Product) SetPrice(price float64) error {

	if price <= 0 {
		return errors.New("el precio debe ser mayor que cero")
	}

	p.price = price
	return nil
}
