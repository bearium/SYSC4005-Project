package main

import (
	"fmt"
	"os"

	"github.com/SYSC4005-Project/pkg/component"
	"github.com/SYSC4005-Project/pkg/inspector"
	"github.com/SYSC4005-Project/pkg/product"
	"github.com/SYSC4005-Project/pkg/workbench"
)

func init() {
	fmt.Println("beginning simulation")
	file1, _ := os.Open("data/servinsp1.dat")
	file2, _ := os.Open("data/servinsp22.dat")
	file3, _ := os.Open("data/servinsp23.dat")
	component1 := component.NewComponent("Component 1", 1, file1)
	component2 := component.NewComponent("Component 2", 2, file2)
	component3 := component.NewComponent("Component 3", 3, file3)
	i1 := inspector.NewInspector("inspector 1", []*component.Component{component1})
	i2 := inspector.NewInspector("inspector 2", []*component.Component{component2, component3})
	p1 := product.NewProduct("product 1", []*component.Component{component1})
	p2 := product.NewProduct("product 2", []*component.Component{component1, component2})
	p3 := product.NewProduct("product 3", []*component.Component{component1, component3})
	w1 := workbench.NewWorkbench("workbench 1", []*inspector.Inspector{i1}, p1)
	w2 := workbench.NewWorkbench("workbench 2", []*inspector.Inspector{i1, i2}, p2)
	w3 := workbench.NewWorkbench("workbench 3", []*inspector.Inspector{i1, i2}, p3)
	go i1.ReadData()
	go i2.ReadData()
}
func main() {
	for {

	}
}
