package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/SYSC4005-Project/pkg/component"
	"github.com/SYSC4005-Project/pkg/inspector"
	"github.com/SYSC4005-Project/pkg/product"
	"github.com/SYSC4005-Project/pkg/workbench"
)

func main() {
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
	mutex := &sync.Mutex{}
	w1 := workbench.NewWorkbench("workbench 1", p1, benchFile1, mutex)
	w2 := workbench.NewWorkbench("workbench 2", p2, benchFile2, mutex)
	w3 := workbench.NewWorkbench("workbench 3", p3, benchFile3, mutex)
	i1 := inspector.NewInspector("inspector 1", []*component.Component{component1}, []*workbench.Workbench{w1, w2, w3}, mutex)
	i2 := inspector.NewInspector("inspector 2", []*component.Component{component2, component3}, []*workbench.Workbench{w2, w3}, mutex)

	//starting threads
	go i1.ReadData()
	go i2.ReadData()
	go w1.ReadData()
	go w2.ReadData()
	go w3.ReadData()
	start := time.Now()

	for {
		if i1.Blocked && i2.Blocked && w1.Blocked && w2.Blocked && w3.Blocked {
			t := time.Now()
			elapsed := t.Local().Sub(start)
			fmt.Printf("total time: %v\n", elapsed)
			fmt.Printf("total idle time for %s: %v\n", i1.Name, i1.IdleTime)
			fmt.Printf("total idle time for %s: %v\n", i2.Name, i2.IdleTime)
			fmt.Printf("total products produced for %s: %v\n", w1.Name, w1.TotalProduced)
			fmt.Printf("total products produced for %s: %v\n", w2.Name, w2.TotalProduced)
			fmt.Printf("total products produced for %s: %v\n", w3.Name, w3.TotalProduced)
			return
		}
	}
}
