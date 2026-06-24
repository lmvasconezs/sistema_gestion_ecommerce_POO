/*
Autor: Luis Miguel Vasconez
Fecha: 1/06/2026
Descripcion: Desarrollo de carrito
*/

package services

import (
	"fmt"
	"sistema_gestion_ecommerce_POO/models"
	"sync"
)

type CartService struct {
	mu    sync.Mutex
	Items []models.CartItem
}

// Añadir producto
func (cs *CartService) AddToCart(item models.CartItem) {
	// Mutex para evitar escrituras simultáneas en el carrito
	cs.mu.Lock()
	defer cs.mu.Unlock()

	cs.Items = append(cs.Items, item)
}

// Calcular total

// CalculateTotal recorre todos los productos
// del carrito y calcula el valor total
// de la compra realizada.
func (cs CartService) CalculateTotal() float64 {

	var total float64

	for _, item := range cs.Items {

		total += item.GetProduct().GetPrice() *
			float64(item.GetQuantity())
	}

	return total
}

func (cs CartService) ShowCart() {

	if len(cs.Items) == 0 {

		fmt.Println("El carrito está vacío.")
		return
	}

	fmt.Println("\n===== CARRITO =====")

	for _, item := range cs.Items {

		fmt.Printf(
			"%s x%d\n",
			item.GetProduct().GetName(),
			item.GetQuantity(),
		)
	}

	fmt.Printf(
		"\nTotal: $%.2f\n",
		cs.CalculateTotal(),
	)
}

func (cs *CartService) ClearCart() {
	// Mutex para evitar vaciar el carrito mientras se añaden items
	cs.mu.Lock()
	defer cs.mu.Unlock()

	cs.Items = []models.CartItem{}
}
