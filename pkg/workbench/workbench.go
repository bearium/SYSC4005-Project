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
	Name           string
	Product        *product.Product
	ComponentArray map[string][]*component.Component
	File           *os.File
	Scanner        *bufio.Scanner
	Blocked        bool
	TotalProduced  int
	Mux            *sync.Mutex
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
		bench.ComponentArray[requirement.Name] = bench.ComponentArray[requirement.Name][:len(bench.ComponentArray[requirement.Name])-1]
		bench.Mux.Unlock()
	}
}

func (bench *Workbench) AddMaterials(component *component.Component) {
	for _, requirement := range bench.Product.RequiredComponents {
		if requirement.Name == component.Name {
			bench.ComponentArray[requirement.Name] = append(bench.ComponentArray[requirement.Name], component)
		}
	}
}

func (bench *Workbench) AddScanner(scanner *bufio.Scanner) {
	bench.Scanner = scanner
}

func (bench *Workbench) ReadData() {
	bench.AddScanner(bufio.NewScanner(bench.File))
	for {
		if bench.canMake() {
			bench.Blocked = false
			if bench.Scanner.Scan() {
				scanText := strings.Trim(bench.Scanner.Text(), " ")
				conv, _ := strconv.ParseFloat(scanText, 64)
				time.Sleep(time.Duration(conv) * time.Millisecond)
				// fmt.Printf("Workbench %s completed %s in %s seconds\n", bench.Name, bench.Product.Name, scanText)
				bench.consumeMaterials()
				bench.TotalProduced++
			} else {
				return
			}
		}
		bench.Blocked = true
	}
}
