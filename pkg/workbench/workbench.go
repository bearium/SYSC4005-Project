package workbench

import (
	"bufio"
	"fmt"
	"os"

	"github.com/SYSC4005-Project/pkg/component"
	"github.com/SYSC4005-Project/pkg/product"
)

type Workbench struct {
	Name           string
	Product        *product.Product
	ComponentArray map[string][]*component.Component
	File           *os.File
	Scanner        *bufio.Scanner
}

func initComponentArray(product *product.Product) map[string][]*component.Component {
	var ComponentArray = make(map[string][]*component.Component)
	for _, requirement := range product.RequiredComponents {
		ComponentArray[requirement.Name] = []*component.Component{}
	}
	return ComponentArray
}

func NewWorkbench(name string, product *product.Product, file *os.File) *Workbench {
	return &Workbench{
		Name:           name,
		Product:        product,
		ComponentArray: initComponentArray(product),
		File:           file,
	}
}

func (bench *Workbench) canMake() bool {
	for _, requirement := range bench.Product.RequiredComponents {
		if bench.ComponentArray[requirement.Name] == nil {
			return false
		}
	}
	return true
}

func (bench *Workbench) consumeMaterials() {
	for _, requirement := range bench.Product.RequiredComponents {
		bench.ComponentArray[requirement.Name] = bench.ComponentArray[requirement.Name][:len(bench.ComponentArray[requirement.Name])-1]
	}
}

func (bench *Workbench) AddMaterials(component *component.Component) {
	for _, requirement := range bench.Product.RequiredComponents {
		if requirement.Name == component.Name {
			bench.ComponentArray[requirement.Name] = append(bench.ComponentArray[requirement.Name], component)
		}
	}
}

func (bench *Workbench) MakeProduct() {
	for {
		if bench.canMake() {
			bench.consumeMaterials()
			fmt.Printf("Made %s", bench.Product.Name)
		}
	}
}

func (bench *Workbench) addScanner(scanner *bufio.Scanner) {
	bench.Scanner = scanner
}
