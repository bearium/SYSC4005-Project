package inspectors

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/SYSC4005-Project/Version2/pkg/workbench"
)

type Inspector1 struct {
	Scanner          *bufio.Scanner
	WaitTime         float64
	BlockedTime      float64
	BufferFull       bool
	TotalProduced    int
	Closed           bool
	Blocked          bool
	TotalTime        float64
	TotalArrivalTime float64
}

func NewInspector1(file *os.File) *Inspector1 {
	return &Inspector1{
		Scanner: bufio.NewScanner(file),
	}
}

func (i *Inspector1) Updatewait() {
	if i.Scanner.Scan() {
		scanText := strings.Trim(i.Scanner.Text(), " ")
		conv, _ := strconv.ParseFloat(scanText, 64)
		i.WaitTime = conv
		i.TotalArrivalTime = i.TotalArrivalTime + conv
	} else {
		i.Closed = true
	}
}

func (i *Inspector1) Time() {
	if i.WaitTime > 0 {
		i.WaitTime = i.WaitTime - .001
	}
}

func (i *Inspector1) UpdateBlockTime() {
	i.Blocked = true
	i.BlockedTime = i.BlockedTime + .001
}

func (i *Inspector1) TryAndPlace(w1 *workbench.WorkBench1, w2 *workbench.Workbench2, w3 *workbench.Workbench3) {
	var currentBench = 4
	var currentMaxBenchComponents = 4
	if len(w1.Buffer) < 2 {
		currentBench = 1
		currentMaxBenchComponents = len(w1.Buffer)
	}
	if currentMaxBenchComponents > len(w2.Buffer) && len(w2.Buffer) < 2 {
		currentBench = 2
		currentMaxBenchComponents = len(w2.Buffer)
	}
	if currentMaxBenchComponents > len(w3.Buffer) && len(w3.Buffer) < 2 {
		currentBench = 3
		currentMaxBenchComponents = len(w3.Buffer)
	}
	if currentBench == 4 {
		return
	}
	if currentBench == 1 {
		w1.Buffer = append(w1.Buffer, 1)
		i.BufferFull = false
		w1.TotalIn++
	}
	if currentBench == 2 {
		w2.Buffer = append(w2.Buffer, 1)
		i.BufferFull = false
		w2.TotalIn++
	}
	if currentBench == 3 {
		w3.Buffer = append(w3.Buffer, 1)
		i.BufferFull = false
		w3.TotalIn++
	}
	i.TotalProduced++
}
