/*
Autor: Luis Miguel Vasconez
Fecha: 30/05/2026
Descripcion: Desarrollo de servicio de productos
*/

package services

import (
	"fmt"
	"sistema_gestion_ecommerce_POO/models"
	"sync"
)

type ProductService struct {
	mu       sync.Mutex
	Products []models.Product
}

// Agregar producto
func (ps *ProductService) AddProduct(product models.Product) {
	// Mutex para proteger modificaciones concurrentes
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ps.Products = append(ps.Products, product)
}

// Mostrar productos
func (ps ProductService) ShowProducts() {

	for _, product := range ps.Products {

		fmt.Println(
			product.GetID(),
			product.GetName(),
			product.GetPrice(),
			product.GetStock(),
		)
	}
}

func (ps ProductService) Search(name string) {

	for _, product := range ps.Products {

		if product.GetName() == name {

			fmt.Println("Producto encontrado:")
			fmt.Println(product.GetName())

			return
		}
	}

	fmt.Println("Producto no encontrado")
}

func (ps ProductService) FindProductByName(name string) (models.Product, bool) {

	for _, product := range ps.Products {

		if product.GetName() == name {
			return product, true
		}
	}

	return models.Product{}, false
}
