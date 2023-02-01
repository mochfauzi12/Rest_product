package product

import (
	"Go_RestApi_Product/domain/application"
	"Go_RestApi_Product/domain/repository"
	"Go_RestApi_Product/domain/service"
)

type productServiceInteractor struct {
	itemRepository repository.ItemRepository
	appRepository  application.ApplicationItem
}

func NewProductServiceInteractor(concreteItemRepo repository.ItemRepository,
	concreteApplicationRepo application.ApplicationItem) service.ProductService {
	return &productServiceInteractor{
		itemRepository: concreteItemRepo,
		appRepository:  concreteApplicationRepo,
	}
}
