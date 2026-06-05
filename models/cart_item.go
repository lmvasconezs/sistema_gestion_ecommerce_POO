/*
Autor: Luis Miguel Vasconez
Fecha: 28/05/2026
Descripcion: Incio desarrollo clase cart item
*/

package models

type CartItem struct {
	product  Product
	quantity int
}

func NewCartItem(product Product, quantity int) CartItem {
	return CartItem{
		product:  product,
		quantity: quantity,
	}
}

func (c CartItem) GetProduct() Product {
	return c.product
}

func (c CartItem) GetQuantity() int {
	return c.quantity
}
