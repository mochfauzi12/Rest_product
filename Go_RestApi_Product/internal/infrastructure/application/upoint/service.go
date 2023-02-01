package upoint

import (
	"Go_RestApi_Product/domain/application"
	"Go_RestApi_Product/domain/entity"
	"net/http"
)

type upointApplication struct {
	httpClient *http.Client
}

func (u *upointApplication) GetListItemByCodeProduct(codeProduct string) ([]*entity.Item, error) {
	//TODO implement me
	panic("implement me")
}

func NewUpointApplication(client *http.Client) application.ApplicationItem {
	return &upointApplication{httpClient: client}
}
