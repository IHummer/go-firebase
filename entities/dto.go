package entities

import (
	"github.com/go-playground/validator/v10"
)

type UserDto struct {
	Uuid     string `json:"uuid"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password"`
}

// Validar producto
func (p *UserDto) Validate() error {
	return validator.New().Struct(p)
}
