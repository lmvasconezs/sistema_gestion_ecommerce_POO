/*
Autor: Luis Miguel Vasconez
Fecha: 23/06/2026
Descripcion: Handler para manejo de peticiones de productos
*/

package handlers

import (
	"encoding/json"
	"net/http"
	"sistema_gestion_ecommerce_POO/models"
	"sistema_gestion_ecommerce_POO/services"
)

// Estructura local para serializar Product a JSON (DTO)
type ProductResponse struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Stock    int     `json:"stock"`
	Category string  `json:"category"`
}

// Estructura local para leer JSON y crear un Product
type ProductRequest struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Stock    int     `json:"stock"`
	Category string  `json:"category"`
}

type ProductHandler struct {
	ProductService *services.ProductService
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	var response []ProductResponse
	
	// Convertimos los modelos a DTOs usando sus getters públicos
	for _, p := range h.ProductService.Products {
		response = append(response, ProductResponse{
			ID:       p.GetID(),
			Name:     p.GetName(),
			Price:    p.GetPrice(),
			Stock:    p.GetStock(),
			Category: p.GetCategory(),
		})
	}

	// Si no hay productos, inicializar como array vacío en vez de null
	if response == nil {
		response = []ProductResponse{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req ProductRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Error al leer el JSON", http.StatusBadRequest)
		return
	}

	// Usamos el constructor original del modelo
	newProduct := models.NewProduct(req.ID, req.Name, req.Price, req.Stock, req.Category)
	h.ProductService.AddProduct(newProduct)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Producto creado con éxito"}`))
}

func (h *ProductHandler) SearchProduct(w http.ResponseWriter, r *http.Request) {
	// Extraer el parámetro de query ?name=...
	name := r.URL.Query().Get("name")
	
	product, found := h.ProductService.FindProductByName(name)
	if !found {
		http.Error(w, `{"error": "Producto no encontrado"}`, http.StatusNotFound)
		return
	}

	response := ProductResponse{
		ID:       product.GetID(),
		Name:     product.GetName(),
		Price:    product.GetPrice(),
		Stock:    product.GetStock(),
		Category: product.GetCategory(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
