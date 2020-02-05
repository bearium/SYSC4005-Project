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
	ComponentArray map[string]*component.Component
}

func initComponentArray(product *product.Product) [][]*component.Component {
	var ComponentArray map[string]*component.Component
	for i, requirement := range product.RequiredComponents {
		ComponentArray()
	}
}

func NewWorkbench(name string, inspectors *inspector.Inspector, product *product.Product) *Workbench {
	return &Workbench{
		Name: name,
		Inspectors: inspectors,
		Product: product,
		ComponentArray
	}
}
