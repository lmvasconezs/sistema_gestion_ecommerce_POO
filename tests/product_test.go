/*
Autor: Luis Miguel Vasconez
Fecha: 23/06/2026
Descripcion: Pruebas unitarias para ProductService
*/

package tests

import (
	"sistema_gestion_ecommerce_POO/models"
	"sistema_gestion_ecommerce_POO/services"
	"testing"
)

func TestAddProduct(t *testing.T) {
	productService := services.ProductService{}

	// Crear datos de ejemplo
	product := models.NewProduct(1, "Laptop", 1200.0, 10, "Electronica")
	
	// Ejecutar método real
	productService.AddProduct(product)

	// Validar resultados
	if len(productService.Products) != 1 {
		t.Errorf("Se esperaba 1 producto, se encontraron %d", len(productService.Products))
	}

	if productService.Products[0].GetName() != "Laptop" {
		t.Errorf("Se esperaba el producto Laptop, se encontró %s", productService.Products[0].GetName())
	}
}

func TestFindProductByName(t *testing.T) {
	productService := services.ProductService{}

	// Crear datos de ejemplo
	product := models.NewProduct(1, "Mouse", 25.0, 20, "Accesorios")
	productService.AddProduct(product)

	// Ejecutar método real para caso exitoso
	foundProduct, found := productService.FindProductByName("Mouse")

	// Validar resultados exitosos
	if !found {
		t.Errorf("Se esperaba encontrar el producto Mouse")
	}

	if foundProduct.GetID() != 1 {
		t.Errorf("Se esperaba ID 1, se encontró %d", foundProduct.GetID())
	}

	// Ejecutar método real para caso fallido
	_, foundMissing := productService.FindProductByName("Teclado")
	
	// Validar resultados fallidos
	if foundMissing {
		t.Errorf("No se esperaba encontrar el producto Teclado")
	}
}
