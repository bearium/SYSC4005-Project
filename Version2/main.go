package main

import (
	"fmt"
	"os"

	"github.com/SYSC4005-Project/Version2/data"
	"github.com/SYSC4005-Project/Version2/pkg/inspectors"
	"github.com/SYSC4005-Project/Version2/pkg/workbench"
)

func main() {

	data.Generate("data/servinsp1Generate.dat", 0.0125, 0.9729, 1000)
	data.Generate("data/servinsp22Generate.dat", 0.0125, 1.4594, 1000)
	data.Generate("data/servinsp23Generate.dat", 0.0135, 1.5608, 1000)
	data.Generate("data/ws1Generate.dat", 0.0127, 0.4044, 1000)
	data.Generate("data/ws2Generate.dat", 0.0149, 0.6084, 1000)
	data.Generate("data/ws3Generate.dat", 0.0125, 0.8026, 1000)

	file, _ := os.Open("data/servinsp1Generate.dat")
	i1 := inspectors.NewInspector1(file)
	file, _ = os.Open("data/servinsp22Generate.dat")
	file2, _ := os.Open("data/servinsp23Generate.dat")
	i2 := inspectors.NewInspector2(file, file2)
	file, _ = os.Open("data/ws1Generate.dat")
	w1 := workbench.NewWorkBench1(file)
	file, _ = os.Open("data/ws2Generate.dat")
	w2 := workbench.NewWorkBench2(file)
	file, _ = os.Open("data/ws3Generate.dat")
	w3 := workbench.NewWorkBench3(file)
	TotalTime := 0.0
	for {
		TotalTime = TotalTime + 0.001
		if !i1.Closed {
			i1.TotalTime = i1.TotalTime + 0.001
			if i1.BufferFull == true {
				i1.UpdateBlockTime()
			} else if i1.WaitTime <= 0 {
				i1.Updatewait()
			} else {
				i1.Time()
			}
			if i1.WaitTime <= 0 && !i1.Closed {
				i1.BufferFull = true
			}
			if i1.BufferFull == true {
				i1.TryAndPlace(w1, w2, w3)
			}
		} else {
			i1.Blocked = true
		}

		i2.TotalTime = i2.TotalTime + 0.001
		if i2.BufferFull == true {
			i2.UpdateBlockTime()
		} else if i2.WaitTime <= 0 {
			i2.Updatewait()
		} else {
			i2.Time()
		}
		if i2.WaitTime <= 0 {
			i2.BufferFull = true
		}
		if i2.BufferFull == true {
			i2.TryAndPlace(w2, w3)
		}

		if !w1.Closed {
			if w1.WaitTime <= 0 && w1.CanMake() {
				w1.Updatewait()
			} else if w1.WaitTime > 0 {
				w1.Time()
			} else if !w1.CanMake() {
				if w1.Started {
					w1.UpdateBlockTime()
				}
			}

			if w1.WaitTime <= 0 && w1.Started && w1.ComponentMade {
				w1.AverageTotalTime = w1.AverageTotalTime + w1.TimeSinceLast
				_, w1.Buffer = w1.Buffer[len(w1.Buffer)-1], w1.Buffer[:len(w1.Buffer)-1]
				w1.ComponentMade = false
				w1.TotalComponents++
			}
		}

		if !w2.Closed {
			if w2.WaitTime <= 0 && w2.CanMake() {
				w2.Updatewait()
			} else if w2.WaitTime > 0 {
				w2.Time()
			} else if !w2.CanMake() {

				if w2.Started {
					w2.UpdateBlockTime()
				}
			}

			if w2.WaitTime <= 0 && w2.Started && w2.ComponentMade {
				w2.AverageTotalTime = w2.AverageTotalTime + w2.TimeSinceLast
				_, w2.Buffer = w2.Buffer[len(w2.Buffer)-1], w2.Buffer[:len(w2.Buffer)-1]
				_, w2.Buffer2 = w2.Buffer2[len(w2.Buffer2)-1], w2.Buffer2[:len(w2.Buffer2)-1]
				w2.ComponentMade = false
				w2.TotalComponents++
			}
		}

		if !w3.Closed {
			if w3.WaitTime <= 0 && w3.CanMake() {
				w3.Updatewait()
			} else if w3.WaitTime > 0 {
				w3.Time()
			} else if !w3.CanMake() {
				if w3.Started {
					w3.UpdateBlockTime()
				}
			}

			if w3.WaitTime <= 0 && w3.Started && w3.ComponentMade {
				_, w3.Buffer = w3.Buffer[len(w3.Buffer)-1], w3.Buffer[:len(w3.Buffer)-1]
				_, w3.Buffer2 = w3.Buffer2[len(w3.Buffer2)-1], w3.Buffer2[:len(w3.Buffer2)-1]
				w3.ComponentMade = false
				w3.TotalComponents++
			}
		}

		w1.UpdateAverageTimes()
		w2.UpdateAverageTimes()
		w3.UpdateAverageTimes()
		if w2.Started {
			w2.TotalTimer = w2.TotalTimer + 0.001
		}

		if len(w1.Buffer) > 0 {
			w1.TimeSinceLast = w1.TimeSinceLast + 0.001
		}
		if len(w2.Buffer) > 0 {
			w2.TimeSinceLast = w2.TimeSinceLast + 0.001
		}

		//return
		if i1.Blocked && i2.Blocked && w1.Blocked && w2.Blocked && w3.Blocked {
			fmt.Println("LITTLES LAW:")
			fmt.Println("Inspector 1:")
			arivalRate := float64(i1.TotalProduced) / float64(i1.TotalArrivalTime)
			fmt.Printf("Lambda= %v/s\n", arivalRate)
			w := float64(i1.TotalTime) / float64(i1.TotalProduced)
			fmt.Printf("w=%v\n", w)
			fmt.Printf("lambda*w= %v\n", arivalRate*w)
			TotalAverage := float64(1)
			fmt.Printf("L= %v\n", TotalAverage)
			fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

			fmt.Println("Inspector 2:")
			arivalRate = float64(i2.TotalProduced) / float64(i2.TotalArrivalTime)
			fmt.Printf("Lambda= %v/s\n", arivalRate)
			w = (i2.TotalTime - i2.BlockedTime) / float64(i2.TotalProduced)
			fmt.Printf("w=%v\n", w)
			fmt.Printf("lambda*w= %v\n", arivalRate*w)
			TotalAverage = float64(1)
			fmt.Printf("L= %v\n", TotalAverage)
			fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

			fmt.Println("WorkBench 1:")
			arivalRate = float64(w1.TotalComponents) / TotalTime
			fmt.Printf("Lambda= %v/s\n", arivalRate)
			fmt.Println(w1.TotalComponents)
			w = float64(w1.AverageTotalTime) / float64(w1.TotalComponents)
			fmt.Printf("w=%v\n", w)
			fmt.Printf("lambda*w= %v\n", arivalRate*w)

			fmt.Printf("L= %v\n", w1.TotalAverage(TotalTime))
			fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
			fmt.Println("WorkBench 2:")
			arivalRate = float64(w2.TotalComponents) / w2.TotalTimer
			fmt.Printf("Lambda= %v/s\n", arivalRate)
			w = w2.AverageTotalTime / float64(w2.TotalComponents)
			fmt.Printf("w=%v\n", w)
			fmt.Printf("lambda*w= %v\n", arivalRate*w)

			fmt.Printf("L= %v\n", w2.TotalAverage(w2.TotalTimer))
			fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

			fmt.Println("WorkBench 3:")
			arivalRate = float64(w3.TotalComponents) / float64(w1.TotalArrivalTime)
			fmt.Printf("Lambda= %v/s\n", arivalRate)
			w = float64(TotalTime) / float64(w3.TotalComponents)
			fmt.Printf("w=%v\n", w)
			fmt.Printf("lambda*w= %v\n", arivalRate*w)

			fmt.Printf("L= %v\n", w3.TotalAverage(TotalTime))

			return
		}
	}
}
