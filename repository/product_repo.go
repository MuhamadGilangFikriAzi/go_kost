package repository

import (
	"gokost.com/m/model"
	"gorm.io/gorm"
)

type ProductRepo interface {
	GetAll() ([]model.Product, error)
}

type productRepo struct {
	mysqlConn *gorm.DB
}

func (p *productRepo) GetAll() ([]model.Product, error) {
	var products []model.Product
	result := p.mysqlConn.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func NewProductRepo(dbConn *gorm.DB) ProductRepo {
	return &productRepo{
		mysqlConn: dbConn,
	}
}
