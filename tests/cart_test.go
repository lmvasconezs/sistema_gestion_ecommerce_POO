/*
Autor: Luis Miguel Vasconez
Fecha: 23/06/2026
Descripcion: Pruebas unitarias para CartService
*/

package tests

import (
	"sistema_gestion_ecommerce_POO/models"
	"sistema_gestion_ecommerce_POO/services"
	"testing"
)

func TestAddToCart(t *testing.T) {
	cartService := services.CartService{}

	// Crear datos de ejemplo
	product := models.NewProduct(1, "Monitor", 250.0, 8, "Electronica")
	cartItem := models.NewCartItem(product, 2)

	// Ejecutar método real
	cartService.AddToCart(cartItem)

	// Validar resultados
	if len(cartService.Items) != 1 {
		t.Errorf("Se esperaba 1 item en el carrito, se encontraron %d", len(cartService.Items))
	}

	if cartService.Items[0].GetQuantity() != 2 {
		t.Errorf("Se esperaba cantidad 2, se encontró %d", cartService.Items[0].GetQuantity())
	}
}

func TestCalculateTotal(t *testing.T) {
	cartService := services.CartService{}

	// Crear datos de ejemplo
	product1 := models.NewProduct(1, "Mouse", 25.0, 20, "Accesorios")
	product2 := models.NewProduct(2, "Teclado", 45.0, 15, "Accesorios")

	item1 := models.NewCartItem(product1, 2) // Total parcial: 2 * 25 = 50
	item2 := models.NewCartItem(product2, 1) // Total parcial: 1 * 45 = 45

	cartService.AddToCart(item1)
	cartService.AddToCart(item2)

	// Ejecutar método real
	total := cartService.CalculateTotal()
	expectedTotal := 95.0

	// Validar resultados
	if total != expectedTotal {
		t.Errorf("Se esperaba un total de %.2f, se obtuvo %.2f", expectedTotal, total)
	}
}
