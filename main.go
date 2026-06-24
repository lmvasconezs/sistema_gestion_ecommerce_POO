/*
Autor: Luis Miguel Vasconez
Fecha: 02/06/2026 (Actualizado 23/06/2026)
Descripcion: Entrypoint de la API REST del sistema
*/

package main

import (
	"fmt"
	"net/http"

	"sistema_gestion_ecommerce_POO/handlers"
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

	// Productos iniciales (carga quemada igual que en consola)
	productService.AddProduct(models.NewProduct(1, "Laptop", 1200, 10, "Electronica"))
	productService.AddProduct(models.NewProduct(2, "Mouse", 25, 20, "Accesorios"))
	productService.AddProduct(models.NewProduct(3, "Teclado", 45, 15, "Accesorios"))
	productService.AddProduct(models.NewProduct(4, "Monitor", 250, 8, "Electronica"))

	// Inicialización de Handlers
	productHandler := handlers.ProductHandler{
		ProductService: &productService,
	}
	
	cartHandler := handlers.CartHandler{
		CartService:    &cartService,
		ProductService: &productService,
	}

	orderHandler := handlers.OrderHandler{
		OrderService: &orderService,
		CartService:  &cartService,
		OrderID:      &orderID, // Pasamos por referencia para incrementar en cada compra
	}

	// Configuración de Rutas (Endpoints HTTP)
	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			productHandler.GetProducts(w, r)
		} else if r.Method == http.MethodPost {
			productHandler.CreateProduct(w, r)
		} else {
			http.Error(w, "Método no soportado", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/products/search", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			productHandler.SearchProduct(w, r)
		} else {
			http.Error(w, "Método no soportado", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/cart", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			cartHandler.GetCart(w, r)
		} else if r.Method == http.MethodPost {
			cartHandler.AddToCart(w, r)
		} else if r.Method == http.MethodDelete {
			cartHandler.ClearCart(w, r)
		} else {
			http.Error(w, "Método no soportado", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/checkout", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			orderHandler.Checkout(w, r)
		} else {
			http.Error(w, "Método no soportado", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			orderHandler.GetOrders(w, r)
		} else {
			http.Error(w, "Método no soportado", http.StatusMethodNotAllowed)
		}
	})

	// Iniciar Servidor
	fmt.Println("=================================================")
	fmt.Println("Servidor API REST iniciado en http://localhost:8080")
	fmt.Println("Rutas disponibles:")
	fmt.Println("- GET    /products")
	fmt.Println("- POST   /products")
	fmt.Println("- GET    /products/search?name=Producto")
	fmt.Println("- GET    /cart")
	fmt.Println("- POST   /cart")
	fmt.Println("- DELETE /cart")
	fmt.Println("- GET    /orders")
	fmt.Println("- POST   /checkout")
	fmt.Println("=================================================")
	
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
