package application

import "Go_RestApi_Product/domain/entity"

type ApplicationItem interface {
	GetListItemByCodeProduct(codeProduct string) ([]*entity.Item, error)
}