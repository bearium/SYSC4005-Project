package workbench

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/SYSC4005-Project/pkg/component"
	"github.com/SYSC4005-Project/pkg/product"
)

type Workbench struct {
	Name                   string
	Product                *product.Product
	ComponentArray         map[string][]*component.Component
	File                   *os.File
	Scanner                *bufio.Scanner
	Blocked                bool
	TotalProduced          int
	Mux                    *sync.Mutex
	Close                  bool
	StartTime              time.Time
	ClosedTime             time.Time
	TotalIdle              time.Duration
	timeSinceComponent     map[int]time.Duration
	QueueData              *Componentqueue
	TotalComponents        int
	TotalArrivalTIme       time.Duration
	TimeSinceLastComponent time.Time
	TimeInSystem           time.Duration
}

type Componentqueue struct {
	CurrentQueueSize    int
	ItemsInQueue        map[int]time.Duration
	TimeSinceLastUpdate time.Time
}

func initComponentArray(product *product.Product) map[string][]*component.Component {
	var ComponentArray = make(map[string][]*component.Component)
	for _, requirement := range product.RequiredComponents {
		ComponentArray[requirement.Name] = []*component.Component{}
	}
	return ComponentArray
}

func NewWorkbench(name string, product *product.Product, file *os.File, mux *sync.Mutex) *Workbench {
	return &Workbench{
		Name:           name,
		Product:        product,
		ComponentArray: initComponentArray(product),
		File:           file,
		Mux:            mux,
		QueueData: &Componentqueue{
			ItemsInQueue: make(map[int]time.Duration),
		},
	}
}

func (bench *Workbench) canMake() bool {
	bench.Mux.Lock()
	for _, requirement := range bench.Product.RequiredComponents {
		if len(bench.ComponentArray[requirement.Name]) == 0 {
			bench.Mux.Unlock()
			return false
		}
	}
	bench.Mux.Unlock()
	return true
}

func (bench *Workbench) consumeMaterials() {
	for _, requirement := range bench.Product.RequiredComponents {
		bench.Mux.Lock()
		component := bench.ComponentArray[requirement.Name][len(bench.ComponentArray[requirement.Name])-1]
		component.TimeOut = time.Now()

		bench.TimeInSystem = bench.TimeInSystem + component.TimeOut.Sub(component.TimeIn)
		bench.ComponentArray[requirement.Name] = bench.ComponentArray[requirement.Name][:len(bench.ComponentArray[requirement.Name])-1]
		elapsed := time.Now()
		requirement.QueueData.ItemsInQueue[requirement.QueueData.CurrentQueueSize] = requirement.QueueData.ItemsInQueue[requirement.QueueData.CurrentQueueSize] + elapsed.Sub(requirement.QueueData.TimeSinceLastUpdate)
		requirement.QueueData.TimeSinceLastUpdate = elapsed
		requirement.QueueData.CurrentQueueSize--
		bench.QueueData.ItemsInQueue[bench.QueueData.CurrentQueueSize] = bench.QueueData.ItemsInQueue[bench.QueueData.CurrentQueueSize] + elapsed.Sub(bench.QueueData.TimeSinceLastUpdate)
		bench.QueueData.TimeSinceLastUpdate = elapsed
		bench.QueueData.CurrentQueueSize--
		bench.Mux.Unlock()
	}
}

func (bench *Workbench) AddMaterials(component *component.Component) {
	for _, requirement := range bench.Product.RequiredComponents {
		if requirement.Name == component.Name {
			if bench.TimeSinceLastComponent.IsZero() {
				bench.TimeSinceLastComponent = time.Now()
				bench.TotalComponents++
			} else {
				elapsed := time.Now()
				bench.TotalArrivalTIme = bench.TotalArrivalTIme + elapsed.Sub(bench.TimeSinceLastComponent)
				bench.TimeSinceLastComponent = elapsed
				bench.TotalComponents++
			}
			component.TimeIn = time.Now()
			bench.ComponentArray[requirement.Name] = append(bench.ComponentArray[requirement.Name], component)
			if bench.QueueData.TimeSinceLastUpdate.IsZero() {
				bench.QueueData.TimeSinceLastUpdate = time.Now()
				bench.QueueData.CurrentQueueSize = 1
			} else {
				elapsed := time.Now()
				bench.QueueData.ItemsInQueue[bench.QueueData.CurrentQueueSize] = bench.QueueData.ItemsInQueue[bench.QueueData.CurrentQueueSize] + elapsed.Sub(bench.QueueData.TimeSinceLastUpdate)
				bench.QueueData.TimeSinceLastUpdate = elapsed
				bench.QueueData.CurrentQueueSize++
			}
		}
	}
}

func (bench *Workbench) AddScanner(scanner *bufio.Scanner) {
	bench.Scanner = scanner
}

func (bench *Workbench) ReadData() {
	bench.AddScanner(bufio.NewScanner(bench.File))
	var startIdle time.Time
	for {
		if bench.Close {
			bench.ClosedTime = time.Now()

			elapsed := time.Now()
			bench.QueueData.ItemsInQueue[bench.QueueData.CurrentQueueSize] = bench.QueueData.ItemsInQueue[bench.QueueData.CurrentQueueSize] + elapsed.Sub(bench.QueueData.TimeSinceLastUpdate)
			bench.QueueData.TimeSinceLastUpdate = elapsed

			return
		}
		if bench.canMake() {
			if !startIdle.IsZero() {
				stopIdle := time.Now()
				bench.TotalIdle = bench.TotalIdle + stopIdle.Sub(startIdle)
				startIdle = time.Time{}
			}
			if bench.StartTime.IsZero() {
				bench.StartTime = time.Now()
			}
			bench.Blocked = false
			if bench.Scanner.Scan() {
				scanText := strings.Trim(bench.Scanner.Text(), " ")
				conv, _ := strconv.ParseFloat(scanText, 64)
				time.Sleep(time.Duration(conv) * time.Millisecond)
				bench.consumeMaterials()
				bench.TotalProduced++
			} else {
				bench.ClosedTime = time.Now()
				elapsed := time.Now()
				bench.QueueData.ItemsInQueue[bench.QueueData.CurrentQueueSize] = bench.QueueData.ItemsInQueue[bench.QueueData.CurrentQueueSize] + elapsed.Sub(bench.QueueData.TimeSinceLastUpdate)
				bench.QueueData.TimeSinceLastUpdate = elapsed

				return
			}
		}
		if startIdle.IsZero() {
			startIdle = time.Now()
		}
		bench.Blocked = true
	}
}
