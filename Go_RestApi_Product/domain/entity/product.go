package entity

type Product struct {
	name        string
	slug        string
	codeProduct string
	version     int
	items       []*Item
}

type DTOAddItemIkea struct {
	Name     string
	CodeItem string
}

func NewIkeaProduct(Name string, Slug string, Code string, Version int) *Product {
	return &Product{
		name:        Name,
		slug:        Slug,
		codeProduct: Code,
		version:     Version,
	}

}

func (pro *Product) GetName() string {
	return pro.name
}

func (pro *Product) GetSlug() string {
	return pro.slug
}

func (prod *Product) GetCodeProduct() string {
	return prod.slug
}

func (prod *Product) GetVersion() int {
	return prod.version
}

func (prod *Product) AddItem(listItems []*Item) *Product {
	prod.items = listItems
	return prod

}

func (prod *Product) IsExternalProduk() bool {
	return prod.version == 1

}

func (prod *Product) IsInternalProduk() bool {
	return prod.version == 2
}
