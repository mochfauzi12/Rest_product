package service

import "Go_RestApi_Product/domain/entity"


type ProductService interface {
	GetListItemOfProduct(product *entity.Product) (*entity.Product, error) 
}
