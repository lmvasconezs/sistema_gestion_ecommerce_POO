/*
Autor: Luis Miguel Vasconez
Fecha: 02/06/2026
Descripcion: Menu principal del sistema
*/

package main

import (
	"fmt"

	"sistema_gestion_ecommerce_POO/models"
	"sistema_gestion_ecommerce_POO/services"
)

func main() {

	// Servicios principales
	productService := services.ProductService{}
	cartService := services.CartService{}
	orderService := services.OrderService{}

	// Contador de pedidos
	orderID := 1

	// Productos iniciales
	productService.AddProduct(
		models.NewProduct(
			1,
			"Laptop",
			1200,
			10,
			"Electronica",
		),
	)

	productService.AddProduct(
		models.NewProduct(
			2,
			"Mouse",
			25,
			20,
			"Accesorios",
		),
	)

	productService.AddProduct(
		models.NewProduct(
			3,
			"Teclado",
			45,
			15,
			"Accesorios",
		),
	)

	productService.AddProduct(
		models.NewProduct(
			4,
			"Monitor",
			250,
			8,
			"Electronica",
		),
	)

	for {

		fmt.Println("\n===== SISTEMA DE E-COMMERCE =====")
		fmt.Println("1. Ver productos")
		fmt.Println("2. Buscar producto")
		fmt.Println("3. Añadir al carrito")
		fmt.Println("4. Ver carrito")
		fmt.Println("5. Comprar")
		fmt.Println("6. Ver pedidos")
		fmt.Println("7. Salir")

		fmt.Print("\nSeleccione una opcion: ")

		var option int
		fmt.Scanln(&option)

		switch option {

		case 1:

			fmt.Println("\n===== LISTA DE PRODUCTOS =====")
			productService.ShowProducts()

		case 2:

			var productName string

			fmt.Print("Ingrese el nombre del producto: ")
			fmt.Scanln(&productName)

			productService.Search(productName)

		case 3:

			var productName string

			fmt.Print("Ingrese el nombre del producto: ")
			fmt.Scanln(&productName)

			product, found := productService.FindProductByName(productName)

			if !found {

				fmt.Println("Producto no encontrado.")
				break
			}

			item := models.NewCartItem(product, 1)

			cartService.AddToCart(item)

			fmt.Println("Producto añadido al carrito.")

		case 4:

			cartService.ShowCart()

		case 5:

			if len(cartService.Items) == 0 {

				fmt.Println("El carrito está vacío.")
				break
			}

			total := cartService.CalculateTotal()

			order := models.NewOrder(
				orderID,
				cartService.Items,
				total,
			)

			orderService.CreateOrder(order)

			orderID++

			cartService.ClearCart()

			fmt.Println("Pedido generado correctamente.")

		case 6:

			orderService.ShowOrders()

		case 7:

			fmt.Println("Gracias por utilizar el sistema.")
			return

		default:

			fmt.Println("Opción inválida.")
		}
	}
}
