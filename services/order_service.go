/*
Autor: Luis Miguel Vasconez
Fecha: 1/06/2026
Descripcion: Desarrollo de pedidos
*/

package services

import (
	"fmt"
	"sistema_gestion_ecommerce_POO/models"
)

type OrderService struct {
	Orders []models.Order
}

// Crear pedido
func (os *OrderService) CreateOrder(order models.Order) {
	os.Orders = append(os.Orders, order)
}

func (os OrderService) ShowOrders() {

	if len(os.Orders) == 0 {

		fmt.Println("No existen pedidos.")
		return
	}

	fmt.Println("\n===== PEDIDOS =====")

	for _, order := range os.Orders {

		fmt.Printf(
			"Pedido #%d - Total: $%.2f\n",
			order.GetID(),
			order.GetTotal(),
		)
	}
}
