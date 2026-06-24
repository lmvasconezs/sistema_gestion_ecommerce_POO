/*
Autor: Luis Miguel Vasconez
Fecha: 23/06/2026
Descripcion: Handler para manejo y creación de pedidos
*/

package handlers

import (
	"encoding/json"
	"net/http"
	"sistema_gestion_ecommerce_POO/models"
	"sistema_gestion_ecommerce_POO/services"
)

type OrderResponse struct {
	ID    int                `json:"id"`
	Items []CartItemResponse `json:"items"` // Usamos el DTO de cart_handler
	Total float64            `json:"total"`
}

type OrderHandler struct {
	OrderService *services.OrderService
	CartService  *services.CartService
	OrderID      *int // Puntero para mantener el contador global
}

func (h *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	var response []OrderResponse

	for _, order := range h.OrderService.Orders {
		var itemsResp []CartItemResponse
		
		for _, item := range order.GetItems() {
			p := item.GetProduct()
			itemsResp = append(itemsResp, CartItemResponse{
				Product: ProductResponse{
					ID:       p.GetID(),
					Name:     p.GetName(),
					Price:    p.GetPrice(),
					Stock:    p.GetStock(),
					Category: p.GetCategory(),
				},
				Quantity: item.GetQuantity(),
			})
		}

		response = append(response, OrderResponse{
			ID:    order.GetID(),
			Items: itemsResp,
			Total: order.GetTotal(),
		})
	}

	// Si no hay pedidos, devolver arreglo vacío en vez de null
	if response == nil {
		response = []OrderResponse{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *OrderHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	if len(h.CartService.Items) == 0 {
		http.Error(w, "El carrito está vacío", http.StatusBadRequest)
		return
	}

	total := h.CartService.CalculateTotal()

	// Crear pedido
	order := models.NewOrder(
		*h.OrderID,
		h.CartService.Items,
		total,
	)

	h.OrderService.CreateOrder(order)
	
	// Incrementar contador
	*h.OrderID++
	
	// Vaciar carrito
	h.CartService.ClearCart()

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Pedido generado correctamente"}`))
}
