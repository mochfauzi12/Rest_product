package repository

import (
	"Go_RestApi_Product/domain/entity"
)

type IkeaProductRepository interface {
	GetAllProduct() ([]*entity.Product, error)
	GetProductBySlug(slug string) (*entity.Product, error)
	Create(*entity.Product) (*entity.Product, error)
	Update(*entity.Product) (*entity.Product, error)
	Delete(*entity.Product) (*entity.Product, error)
}
