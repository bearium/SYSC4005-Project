package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/SYSC4005-Project/data"
	"github.com/SYSC4005-Project/pkg/component"
	"github.com/SYSC4005-Project/pkg/inspector"
	"github.com/SYSC4005-Project/pkg/product"
	"github.com/SYSC4005-Project/pkg/workbench"
)

func main() {
	fmt.Println("beginning simulation")
	for i := 0; i < 15; i++ {
		data.Generate("data/servinsp1Generate.dat", 0.0125, 0.9729, 300)
		data.Generate("data/servinsp22Generate.dat", 0.0125, 1.4594, 300)
		data.Generate("data/servinsp23Generate.dat", 0.0135, 1.5608, 300)
		data.Generate("data/ws1Generate.dat", 0.0127, 0.4044, 300)
		data.Generate("data/ws2Generate.dat", 0.0149, 0.6084, 300)
		data.Generate("data/ws3Generate.dat", 0.0125, 0.8026, 300)
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
	//	start := time.Now()

	for {
		if i1.Blocked && i2.Blocked && w1.Blocked && w2.Blocked && w3.Blocked {
			i1.Close = true
			i2.Close = true
			w1.Close = true
			w2.Close = true
			w3.Close = true
			// t := time.Now()
			// elapsed := t.Local().Sub(start)
			time.Sleep(1000)
			// fmt.Printf("total time: %v\n", elapsed)
			// fmt.Printf("total idle time for %s: %v\n", i1.Name, i1.IdleTime)
			// fmt.Printf("Running Time for %s: %v\n", i1.Name, i1.ClosedTime.Sub(start))
			// throughput := float64(i1.TotalComponentS) / float64(i1.ClosedTime.Sub(start).Seconds())
			// fmt.Printf("THROUGHPUT INSPECTOR 1: %f/min\n", (throughput / 1000))
			// throughput = float64(i2.TotalComponentS) / float64(i2.ClosedTime.Sub(start).Seconds())
			// fmt.Printf("THROUGHPUT INSPECTOR 2: %f/min\n", (throughput / 1000))
			fmt.Printf("%.5f\n", float64(w3.TotalProduced)/float64(w3.ClosedTime.Sub(w3.StartTime).Seconds()))
			// fmt.Printf("inspector 2 idle: %v%%\n", (float64(i2.IdleTime) / float64(i2.ClosedTime.Sub(start))))
			// fmt.Printf("total idle time for %s: %v\n", i2.Name, i2.IdleTime)
			// fmt.Printf("Running Time for %s: %v\n", i2.Name, i2.ClosedTime.Sub(start))
			// fmt.Printf("total products produced for %s: %v\n", w1.Name, w1.TotalProduced)
			// fmt.Printf("Running Time for %s: %v\n", w1.Name, w1.ClosedTime.Sub(w1.StartTime))
			// throughput := float64(w1.TotalProduced) / float64(w1.ClosedTime.Sub(w1.StartTime).Seconds())
			// fmt.Printf("Throughput: %v/s\n", throughput)
			// arivalRate := float64(w1.TotalComponents) / float64(w1.TotalArrivalTIme.Seconds())
			// fmt.Println("WORKBENCH 1:")
			// fmt.Printf("Arival Rate: %v/s\n", arivalRate)
			// totalTime := float64(w1.ClosedTime.Sub(w1.StartTime).Seconds()) - float64(w1.QueueData.ItemsInQueue[0].Seconds())
			// w := float64(totalTime) / float64(w1.TotalProduced)
			// fmt.Printf("w=%v\n", w)
			// fmt.Printf("lambda*w= %v\n", arivalRate*w)
			// // fmt.Printf("total idle time for %s: %v\n", w1.Name, w1.TotalIdle)

			// var TotalAverage float64
			// for i := 0; i < len(w1.QueueData.ItemsInQueue); i++ {
			// 	average := float64(w1.QueueData.ItemsInQueue[i].Microseconds()) / float64(w1.ClosedTime.Sub(w1.StartTime).Microseconds())
			// 	TotalAverage = TotalAverage + (float64(i) * average)
			// 	// fmt.Printf("total queue time for %s: %v:%v, average:%f\n", w1.Name, i, w1.QueueData.ItemsInQueue[i], average*100)
			// }
			// fmt.Printf("L= %v\n", TotalAverage)

			// // fmt.Printf("total products produced for %s: %v\n", w2.Name, w2.TotalProduced)
			// // fmt.Printf("Running Time for %s: %v\n", w2.Name, w2.ClosedTime.Sub(w2.StartTime))
			// // throughput = float64(w2.TotalProduced) / float64(w2.ClosedTime.Sub(w2.StartTime).Seconds())
			// // fmt.Printf("Throughput: %v/s\n", throughput)
			// fmt.Println("WORKBENCH 2:")
			// arivalRate = float64(w2.TotalComponents) / float64(w2.TotalArrivalTIme.Seconds())
			// fmt.Printf("Arival Rate: %v/s\n", arivalRate)
			// totalTime = float64(w2.ClosedTime.Sub(w2.StartTime).Seconds()) - float64(w2.QueueData.ItemsInQueue[0].Seconds())
			// w = float64(totalTime) / float64(w2.TotalProduced)
			// fmt.Printf("w=%v\n", w)
			// fmt.Printf("lambda*w= %v\n", arivalRate*w)
			// // fmt.Printf("total idle time for %s: %v\n", w2.Name, w2.TotalIdle)

			// TotalAverage = 0
			// for i := 0; i < len(w2.QueueData.ItemsInQueue); i++ {
			// 	average := float64(w2.QueueData.ItemsInQueue[i].Seconds()) / float64(w2.ClosedTime.Sub(w2.StartTime).Seconds())
			// 	TotalAverage = TotalAverage + (float64(i) * average)
			// 	// fmt.Printf("total queue time for %s: %v:%v, average:%f\n", w2.Name, i, w2.QueueData.ItemsInQueue[i], average*100)
			// }
			// fmt.Printf("L= %v\n", TotalAverage)

			// // fmt.Printf("total products produced for %s: %v\n", w3.Name, w3.TotalProduced)
			// // fmt.Printf("Running Time for %s: %v\n", w3.Name, w3.ClosedTime.Sub(w3.StartTime))
			// // throughput = float64(w3.TotalProduced) / float64(w3.ClosedTime.Sub(w3.StartTime).Seconds())
			// // fmt.Printf("Throughput: %v/s\n", throughput)
			// fmt.Println("WORKBENCH 3:")
			// arivalRate = float64(w3.TotalComponents) / float64(w3.TotalArrivalTIme.Seconds())
			// fmt.Printf("Arival Rate: %v/s\n", arivalRate)
			// totalTime = float64(w3.ClosedTime.Sub(w3.StartTime).Seconds()) - float64(w3.QueueData.ItemsInQueue[0].Seconds())
			// w = float64(totalTime) / float64(w3.TotalProduced)
			// fmt.Printf("w=%v\n", w)
			// fmt.Printf("lambda*w= %v\n", arivalRate*w)
			// // fmt.Printf("total idle time for %s: %v\n", w3.Name, w3.TotalIdle)
			// TotalAverage = 0
			// for i := 0; i < len(w3.QueueData.ItemsInQueue); i++ {
			// 	average := float64(w3.QueueData.ItemsInQueue[i].Microseconds()) / float64(w3.ClosedTime.Sub(w3.StartTime).Microseconds())
			// 	TotalAverage = TotalAverage + (float64(i) * average)
			// 	// fmt.Printf("total queue time for %s: %v:%v, average:%f\n", w3.Name, i, w3.QueueData.ItemsInQueue[i], average*100)
			// }
			// fmt.Printf("L= %v\n", TotalAverage)
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
