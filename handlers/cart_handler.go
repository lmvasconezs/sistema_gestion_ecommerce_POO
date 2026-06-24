/*
Autor: Luis Miguel Vasconez
Fecha: 23/06/2026
Descripcion: Handler para manejo del carrito de compras
*/

package handlers

import (
	"encoding/json"
	"net/http"
	"sistema_gestion_ecommerce_POO/models"
	"sistema_gestion_ecommerce_POO/services"
)

// Estructuras locales para serialización JSON
type CartItemResponse struct {
	Product  ProductResponse `json:"product"` // Reutiliza ProductResponse de product_handler
	Quantity int             `json:"quantity"`
}

type CartResponse struct {
	Items []CartItemResponse `json:"items"`
	Total float64            `json:"total"`
}

type AddToCartRequest struct {
	ProductName string `json:"product_name"`
}

type CartHandler struct {
	CartService    *services.CartService
	ProductService *services.ProductService
}

func (h *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	var response CartResponse
	
	// Convertimos los items a DTOs
	for _, item := range h.CartService.Items {
		p := item.GetProduct()
		
		prodResp := ProductResponse{
			ID:       p.GetID(),
			Name:     p.GetName(),
			Price:    p.GetPrice(),
			Stock:    p.GetStock(),
			Category: p.GetCategory(),
		}
		
		response.Items = append(response.Items, CartItemResponse{
			Product:  prodResp,
			Quantity: item.GetQuantity(),
		})
	}

	// Aseguramos que Items sea un arreglo vacío y no null si no hay productos
	if response.Items == nil {
		response.Items = []CartItemResponse{}
	}
	
	response.Total = h.CartService.CalculateTotal()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *CartHandler) AddToCart(w http.ResponseWriter, r *http.Request) {
	var req AddToCartRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Error al leer el JSON", http.StatusBadRequest)
		return
	}

	// Buscar producto como en el flujo original de consola
	product, found := h.ProductService.FindProductByName(req.ProductName)
	if !found {
		http.Error(w, "Producto no encontrado", http.StatusNotFound)
		return
	}

	// Crear el item con cantidad fija 1
	item := models.NewCartItem(product, 1)
	h.CartService.AddToCart(item)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Producto añadido al carrito"}`))
}

func (h *CartHandler) ClearCart(w http.ResponseWriter, r *http.Request) {
	// Reutiliza funcionalidad de servicio para vaciar el carrito
	h.CartService.ClearCart()
	
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Carrito vaciado correctamente"}`))
}
