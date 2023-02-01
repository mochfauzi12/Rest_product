package product

import (
	"Go_RestApi_Product/domain/entity"
	"errors"
)

func (pro *productServiceInteractor) GetListItemOfProduct(product *entity.Product) (*entity.Product, error) {
	//IKEA REPO (APPLICATION)
	if product.IsExternalProduk() {
		listItem, err := pro.appRepository.GetListItemByCodeProduct(product.GetCodeProduct())
		if err != nil {
			return nil, err
		}

		product.AddItem(listItem)
	}

	if product.IsInternalProduk() {
		listItem, err := pro.itemRepository.GetItemsByProductSlug(product.GetSlug())
		if err != nil {
			return nil, err
		}

		product.AddItem(listItem)
	}

	return nil, errors.New("ITEMS NOT FOUND")
}
