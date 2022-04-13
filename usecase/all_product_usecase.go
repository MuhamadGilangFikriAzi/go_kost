package usecase

import (
	"gokost.com/m/model"
	"gokost.com/m/repository"
)

type AllProductUseCase interface {
	GetAllProduct() ([]model.Product, error)
}

type allProductUseCase struct {
	repo repository.ProductRepo
}

func (a *allProductUseCase) GetAllProduct() ([]model.Product, error) {
	return a.repo.GetAll()
}

func NewAllProductUseCase(repo repository.ProductRepo) AllProductUseCase {
	return &allProductUseCase{
		repo,
	}
}
