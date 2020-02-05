package main

import (
	"fmt"
	"os"

	"github.com/SYSC4005-Project/pkg/component"
	"github.com/SYSC4005-Project/pkg/inspector"
)

func init() {
	fmt.Println("beginning simulation")
	file1, _ := os.Open("data/servinsp1.dat")
	file2, _ := os.Open("data/servinsp22.dat")
	file3, _ := os.Open("data/servinsp23.dat")
	component1 := component.NewComponent("Component 1", 1, file1)
	component2 := component.NewComponent("Component 2", 2, file2)
	component3 := component.NewComponent("Component 3", 3, file3)
	i := inspector.NewInspector("inspector 1", []*component.Component{component1})
	i2 := inspector.NewInspector("inspector 2", []*component.Component{component2, component3})
	go i.ReadData()
	go i2.ReadData()
}
func main() {
	for {

	}
}
