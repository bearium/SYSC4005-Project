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
	for i := 1; i < 2; i++ {
		runBenchMark(i)
	}
}

func runBenchMark(i int) {

	//instantiating files to be read
	//component files
	componentFile1, _ := os.Open("data/servinsp1Generate.dat")
	componentFile2, _ := os.Open("data/servinsp22Generate.dat")
	componentFile3, _ := os.Open("data/servinsp23Generate.dat")
	//workbench files
	benchFile1, _ := os.Open("data/ws1Generate.dat")
	benchFile2, _ := os.Open("data/ws2Generate.dat")
	benchFile3, _ := os.Open("data/ws3Generate.dat")

	//initializing objects
	component1 := component.NewComponent("Component 1", 1, componentFile1)
	wb2Component1 := component.NewComponent("Component 1", 1, componentFile1)
	wb3Component1 := component.NewComponent("Component 1", 1, componentFile1)
	component2 := component.NewComponent("Component 2", 2, componentFile2)
	component3 := component.NewComponent("Component 3", 3, componentFile3)
	p1 := product.NewProduct("product 1", []*component.Component{component1})
	p2 := product.NewProduct("product 2", []*component.Component{wb2Component1, component2})
	p3 := product.NewProduct("product 3", []*component.Component{wb3Component1, component3})
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
			i1.Close = true
			i2.Close = true
			w1.Close = true
			w2.Close = true
			w3.Close = true
			t := time.Now()
			elapsed := t.Local().Sub(start)
			time.Sleep(1000)
			fmt.Printf("total time: %v\n", elapsed)
			fmt.Printf("total idle time for %s: %v\n", i1.Name, i1.IdleTime)
			fmt.Printf("Running Time for %s: %v\n", i1.Name, i1.ClosedTime.Sub(start))
			fmt.Printf("total idle time for %s: %v\n", i2.Name, i2.IdleTime)
			fmt.Printf("Running Time for %s: %v\n", i2.Name, i2.ClosedTime.Sub(start))
			fmt.Printf("total products produced for %s: %v\n", w1.Name, w1.TotalProduced)
			fmt.Printf("Running Time for %s: %v\n", w1.Name, w1.ClosedTime.Sub(w1.StartTime))
			fmt.Printf("total idle time for %s: %v\n", w1.Name, w1.TotalIdle)
			for _, component := range w1.Product.RequiredComponents {
				for i := 0; i < len(component.QueueData.ItemsInQueue); i++ {
					fmt.Printf("total queue time for %s: %v:%v\n", component.Name, i, component.QueueData.ItemsInQueue[i])
				}
			}
			fmt.Printf("total products produced for %s: %v\n", w2.Name, w2.TotalProduced)
			fmt.Printf("Running Time for %s: %v\n", w2.Name, w2.ClosedTime.Sub(w2.StartTime))
			fmt.Printf("total idle time for %s: %v\n", w2.Name, w2.TotalIdle)
			for _, component := range w2.Product.RequiredComponents {
				for i := 0; i < len(component.QueueData.ItemsInQueue); i++ {
					fmt.Printf("total queue time for %s: %v:%v\n", component.Name, i, component.QueueData.ItemsInQueue[i])
				}
			}
			fmt.Printf("total products produced for %s: %v\n", w3.Name, w3.TotalProduced)
			fmt.Printf("Running Time for %s: %v\n", w3.Name, w3.ClosedTime.Sub(w3.StartTime))
			fmt.Printf("total idle time for %s: %v\n", w3.Name, w3.TotalIdle)
			for _, component := range w3.Product.RequiredComponents {
				for i := 0; i < len(component.QueueData.ItemsInQueue); i++ {
					fmt.Printf("total queue time for %s: %v:%v\n", component.Name, i, component.QueueData.ItemsInQueue[i])
				}
			}
			time.Sleep(1000)
			componentFile1.Close()
			componentFile2.Close()
			componentFile3.Close()
			//workbench files
			benchFile1.Close()
			benchFile2.Close()
			benchFile3.Close()
			return
		}
	}

}
