package usecase

import "Go_RestApi_Product/domain/entity"

type ProductUsecase interface {
	GetDetailProduct(slug string) (*entity.Product,error)
}