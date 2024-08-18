//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/Ingrid-Paulino/DI/product"
	"github.com/google/wire"
)

// passo uma conexão de banco de dados para o repositório
func NewUseCase(db *sql.DB) *product.ProductUseCase {
	wire.Build(
		product.NewProductUseCase,
		product.NewProductRepository,
	)
	return &product.ProductUseCase{}
}
