/*
Autor: Luis Miguel Vasconez
Fecha: 1/06/2026
Descripcion: Desarrollo de validaciones
*/

package utils

import "errors"

// Validar stock
func ValidateStock(stock int) error {

	if stock < 0 {
		return errors.New("stock inválido")
	}

	return nil
}

// Validar precio
func ValidatePrice(price float64) error {

	if price <= 0 {
		return errors.New("precio inválido")
	}

	return nil
}
