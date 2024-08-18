//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/Ingrid-Paulino/DI/product"
	"github.com/google/wire"
)

// toda vez que eu tiver um NewProductRepository sig que ele vai trocar meu productRepositoryInterface, o wire permite criar set de dependencias para agrupar, de forma que possamos reutilizar
var setRepositoryDependency = wire.NewSet(
	product.NewProductRepository,
	wire.Bind(new(product.ProductRepositoryInterface), new(*product.ProductRepository)),
)

func NewUseCase(db *sql.DB) *product.ProductUseCase {
	wire.Build(
		setRepositoryDependency,
		product.NewProductUseCase,
	)
	return &product.ProductUseCase{}
}

// passo uma conexão de banco de dados para o repositório
// trabalhando com interfaces para tirar aclopamento
//func NewUseCase(db *sql.DB) *product.ProductUseCase {
//	wire.Build(
//		product.NewProductUseCase,
//		product.NewProductRepository,
//	)
//	return &product.ProductUseCase{}
//}
