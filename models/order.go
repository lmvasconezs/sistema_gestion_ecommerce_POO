/*
Autor: Luis Miguel Vasconez
Fecha: 28/05/2026
Descripcion: Incio desarrollo clase order
*/

package models

type Order struct {
	id    int
	items []CartItem
	total float64
}

func NewOrder(id int, items []CartItem, total float64) Order {
	return Order{
		id:    id,
		items: items,
		total: total,
	}
}

func (o Order) GetID() int {
	return o.id
}

func (o Order) GetItems() []CartItem {
	return o.items
}

func (o Order) GetTotal() float64 {
	return o.total
}
