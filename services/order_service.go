/*
Autor: Luis Miguel Vasconez
Fecha: 1/06/2026
Descripcion: Desarrollo de pedidos
*/

package services

import (
	"fmt"
	"sistema_gestion_ecommerce_POO/models"
	"sync"
)

type OrderService struct {
	mu     sync.Mutex
	Orders []models.Order
}

// Crear pedido
func (os *OrderService) CreateOrder(order models.Order) {
	// Mutex para proteger el historial de pedidos
	os.mu.Lock()
	defer os.mu.Unlock()

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
