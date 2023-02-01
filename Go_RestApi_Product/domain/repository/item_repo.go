package repository

import (
	"Go_RestApi_Product/domain/entity"
)

type ItemRepository interface {
	GetItemsByProductSlug(slug string) ([]*entity.Item, error)
}
