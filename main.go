package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/SYSC4005-Project/pkg/component"
	"github.com/SYSC4005-Project/pkg/inspector"
	"github.com/SYSC4005-Project/pkg/product"
	"github.com/SYSC4005-Project/pkg/workbench"
)

func init() {
	fmt.Println("beginning simulation")

	//instantiating files to be read
	//component files
	componentFile1, _ := os.Open("data/servinsp1.dat")
	componentFile2, _ := os.Open("data/servinsp22.dat")
	componentFile3, _ := os.Open("data/servinsp23.dat")
	//workbench files
	benchFile1, _ := os.Open("data/ws1.dat")
	benchFile2, _ := os.Open("data/ws2.dat")
	benchFile3, _ := os.Open("data/ws3.dat")

	//initializing objects
	component1 := component.NewComponent("Component 1", 1, componentFile1)
	component2 := component.NewComponent("Component 2", 2, componentFile2)
	component3 := component.NewComponent("Component 3", 3, componentFile3)
	p1 := product.NewProduct("product 1", []*component.Component{component1})
	p2 := product.NewProduct("product 2", []*component.Component{component1, component2})
	p3 := product.NewProduct("product 3", []*component.Component{component1, component3})
	w1 := workbench.NewWorkbench("workbench 1", p1, benchFile1)
	w2 := workbench.NewWorkbench("workbench 2", p2, benchFile2)
	w3 := workbench.NewWorkbench("workbench 3", p3, benchFile3)
	mutex := &sync.Mutex{}
	i1 := inspector.NewInspector("inspector 1", []*component.Component{component1}, []*workbench.Workbench{w1, w2, w3}, mutex)
	i2 := inspector.NewInspector("inspector 2", []*component.Component{component2, component3}, []*workbench.Workbench{w2, w3}, mutex)

	//starting threads
	go i1.ReadData()
	go i2.ReadData()
	go w1.ReadData()
	go w2.ReadData()
	go w3.ReadData()
}
func main() {
	for {

	}
}
