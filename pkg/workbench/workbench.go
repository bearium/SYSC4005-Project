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

func NewWorkbench(name string, product *product.Product) *Workbench {
	return &Workbench{
		Name:           name,
		Product:        product,
		ComponentArray: initComponentArray(product),
	}
}

func (bench *Workbench)canMake() bool{
	for i, requirement := range bench.product.RequiredComponents {
		if bench.ComponentArray[requirement.Name] == nil {
			return false
		}
	}
	return true
}

func (bench *WorkBench)consumeMaterials() {
	for i, requirement := range bench.product.RequiredComponents {
		bench.ComponentArray[requirement.Name] = bench.ComponentArray[requirement.Name][:len(bench.ComponentArray[requirement.Name]) -1]
	}
}

func (bench *WorkBench)addMaterials() {

}
func (bench *Workbench) MakeProduct() {
	for (
		if bench.canMake() {
			bench.consumeMaterials()
			fmt.Printf("Made %s", bench.Product.Name)
		}
	)
}