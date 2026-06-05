/*
Autor: Luis Miguel Vasconez
Fecha: 30/05/2026
Descripcion: Desarrollo de servicio de productos
*/

package services

import (
	"fmt"
	"sistema_gestion_ecommerce_POO/models"
)

type ProductService struct {
	Products []models.Product
}

// Agregar producto
func (ps *ProductService) AddProduct(product models.Product) {
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
