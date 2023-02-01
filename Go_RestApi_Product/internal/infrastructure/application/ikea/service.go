package ikea

import (
	"Go_RestApi_Product/domain/application"
	"Go_RestApi_Product/domain/entity"
	"net/http"
)

type ikeaApplicationInteractor struct {
	httpClient *http.Client
}

func (i *ikeaApplicationInteractor) GetListItemByCodeProduct(codeProduct string) ([]*entity.Item, error) {
	panic("Implement")
}

func NewIkeaApplication(client *http.Client) application.ApplicationItem {
	return &ikeaApplicationInteractor{httpClient: client}
}
