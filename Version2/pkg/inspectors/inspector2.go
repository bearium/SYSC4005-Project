package inspectors

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/SYSC4005-Project/Version2/pkg/workbench"
)

type Inspector2 struct {
	Scanner          *bufio.Scanner
	Scanner2         *bufio.Scanner
	WaitTime         float64
	BlockedTime      float64
	BufferFull       bool
	MadeComonent     int
	TotalProduced    int
	Closed           bool
	Blocked          bool
	TotalTime        float64
	TotalArrivalTime float64
	TotalCounter     float64
}

func NewInspector2(file *os.File, file2 *os.File) *Inspector2 {
	return &Inspector2{
		Scanner:  bufio.NewScanner(file),
		Scanner2: bufio.NewScanner(file2),
	}
}

func (i *Inspector2) Updatewait() {
	randInt := rand.Intn(2)
	if randInt == 0 {
		i.MadeComonent = 0
		i.Scanner.Scan()
		scanText := strings.Trim(i.Scanner.Text(), " ")
		conv, _ := strconv.ParseFloat(scanText, 64)
		i.WaitTime = conv
		i.TotalArrivalTime = i.TotalArrivalTime + conv
	} else {
		i.MadeComonent = 1
		i.Scanner2.Scan()
		scanText := strings.Trim(i.Scanner2.Text(), " ")
		conv, _ := strconv.ParseFloat(scanText, 64)
		i.WaitTime = conv
		i.TotalArrivalTime = i.TotalArrivalTime + conv
	}
}

func (i *Inspector2) Time() {
	if i.WaitTime > 0 {
		i.WaitTime = i.WaitTime - .001
	}
}

func (i *Inspector2) UpdateBlockTime() {
	i.Blocked = true
	i.BlockedTime = i.BlockedTime + .001
}

func (i *Inspector2) TryAndPlace(w2 *workbench.Workbench2, w3 *workbench.Workbench3) {
	if i.MadeComonent == 0 {
		if len(w2.Buffer2) < 2 {
			w2.Buffer2 = append(w2.Buffer2, 1)
			w2.TotalIn++
			i.BufferFull = false
			i.TotalProduced++
		}
	} else {
		if len(w3.Buffer2) < 2 {
			w3.Buffer2 = append(w3.Buffer2, 1)
			i.BufferFull = false
			w3.TotalIn++
			i.TotalProduced++
		}
	}
}
