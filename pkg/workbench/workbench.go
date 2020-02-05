package workbench

import (
	"github.com/SYSC4005-Project/pkg/component"
	"github.com/SYSC4005-Project/pkg/inspector"
	"github.com/SYSC4005-Project/pkg/product"
)

type Workbench struct {
	Name           string
	Inspectors     []*inspector.Inspector
	Product        *product.Product
	ComponentArray map[string][2]*component.Component
}

func initComponentArray(product *product.Product) map[string][2]*component.Component {
	var ComponentArray map[string]interface{}
	for i, requirement := range product.RequiredComponents {
		ComponentArray[requirement.Name] = []*component.Component{}
	}
	return ComponentArray
}

func NewWorkbench(name string, inspectors *inspector.Inspector, product *product.Product) *Workbench {
	return &Workbench{
		Name:           name,
		Inspectors:     inspectors,
		Product:        product,
		ComponentArray: initComponentArray(product),
	}
}
