/*
Autor: Luis Miguel Vasconez
Fecha: 23/06/2026
Descripcion: Pruebas unitarias para OrderService
*/

package tests

import (
	"sistema_gestion_ecommerce_POO/models"
	"sistema_gestion_ecommerce_POO/services"
	"testing"
)

func TestCreateOrder(t *testing.T) {
	orderService := services.OrderService{}

	// Crear datos de ejemplo
	product := models.NewProduct(1, "Laptop", 1200.0, 10, "Electronica")
	item := models.NewCartItem(product, 1)

	items := []models.CartItem{item}
	order := models.NewOrder(1, items, 1200.0)

	// Ejecutar método real
	orderService.CreateOrder(order)

	// Validar resultados
	if len(orderService.Orders) != 1 {
		t.Errorf("Se esperaba 1 orden, se encontraron %d", len(orderService.Orders))
	}

	if orderService.Orders[0].GetID() != 1 {
		t.Errorf("Se esperaba el ID de orden 1, se encontró %d", orderService.Orders[0].GetID())
	}

	if orderService.Orders[0].GetTotal() != 1200.0 {
		t.Errorf("Se esperaba un total de 1200.0, se encontró %.2f", orderService.Orders[0].GetTotal())
	}
}
