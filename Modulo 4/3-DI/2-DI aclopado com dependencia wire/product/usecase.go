package product

type ProductUseCase struct {
	repository *ProductRepository
}

func NewProductUseCase(repository *ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		repository: repository,
	}
}

// GetProduct returns a product by id
// This Product was not supposed to be returned. We should return a DTO instead.
// However, we will return it for now to keep the example simple.
func (u *ProductUseCase) GetProduct(id int) (Product, error) {
	return u.repository.GetProduct(id)
}