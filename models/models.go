package models

import (
	"github.com/go-playground/validator/v10"
)

type UserModel struct {
	Uuid     string
	Name     string
	Email    string
	Password string
}

type Product struct {
	ID          string  `json:"id" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required,numeric"`
}

// Validar producto
func (p *Product) Validate() error {
	return validator.New().Struct(p)
}
