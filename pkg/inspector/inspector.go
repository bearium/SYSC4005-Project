package inspector

import (
	"bufio"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/SYSC4005-Project/pkg/component"
	"github.com/SYSC4005-Project/pkg/workbench"
)

const seed = 42

type Inspector struct {
	Name        string
	Components  []*component.Component
	Workbenches []*workbench.Workbench
	Mux         *sync.Mutex
	IdleTime    time.Duration
}

func NewInspector(name string, components []*component.Component, workbench []*workbench.Workbench, mux *sync.Mutex) *Inspector {
	return &Inspector{
		Name:        name,
		Components:  components,
		Workbenches: workbench,
		Mux:         mux,
	}
}

func (i *Inspector) ReadData() {
	rand.Seed(seed)
	for j := 0; j < len(i.Components); j++ {
		i.Components[j].AddScanner(bufio.NewScanner(i.Components[j].File))
	}

	for {
		randInt := 0
		if len(i.Components) > 1 {
			randInt = rand.Intn(len(i.Components))
		}
		currentComponent := i.Components[randInt]
		if currentComponent.Scanner.Scan() {
			scanText := strings.Trim(currentComponent.Scanner.Text(), " ")
			conv, _ := strconv.ParseFloat(scanText, 64)
			time.Sleep(time.Duration(conv) * time.Millisecond)
			fmt.Printf("Inspector %s completed component %s in %s seconds\n", i.Name, i.Components[randInt].Name, scanText)
			var start time.Time
			for {
				placeWorkBench := i.canPlace(currentComponent)
				if placeWorkBench != nil {
					if !start.IsZero() {
						t := time.Now()
						elapsed := t.Sub(start)
						i.IdleTime = i.IdleTime + elapsed
					}
					i.Mux.Lock()
					placeWorkBench.AddMaterials(currentComponent)
					i.Mux.Unlock()
					break
				}
				if start.IsZero() {
					start = time.Now()
				}
			}
		} else {
			i.Components = append(i.Components[:randInt], i.Components[randInt+1:]...)
			if len(i.Components) == 0 {
				totalIdleTime := i.IdleTime
				fmt.Printf("total idle time for %s: %v\n", i.Name, totalIdleTime)
				return
			}
		}
	}
}

func (i *Inspector) canPlace(currentComponent *component.Component) *workbench.Workbench {
	var currentBench *workbench.Workbench
	var currentMaxBenchComponents int
	for _, bench := range i.Workbenches {
		i.Mux.Lock()
		componentAmount := len(bench.ComponentArray[currentComponent.Name])
		if bench.ComponentArray[currentComponent.Name] != nil && componentAmount < 2 {
			if componentAmount < currentMaxBenchComponents || currentMaxBenchComponents == 0 {
				currentBench = bench
				currentMaxBenchComponents = componentAmount
			}
		}
		i.Mux.Unlock()
	}
	return currentBench
}
