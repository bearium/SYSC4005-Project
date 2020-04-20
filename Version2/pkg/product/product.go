package product

import "github.com/SYSC4005-Project/Version1/pkg/component"

type Product struct {
	Name               string
	RequiredComponents []*component.Component
}

func NewProduct(name string, requiredComponents []*component.Component) *Product {
	return &Product{
		Name:               name,
		RequiredComponents: requiredComponents,
	}

}
