package workbench

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type WorkBench1 struct {
	Scanner          *bufio.Scanner
	WaitTime         float64
	BlockedTime      float64
	BufferEmpty      bool
	TotalComponents  int
	Buffer           []int
	Started          bool
	ComponentMade    bool
	Closed           bool
	Blocked          bool
	TotalArrivalTime float64
	TotalTime0       float64
	TotalTime1       float64
	TotalTime2       float64
	TotalIn          int
	AverageTotalTime float64
	TimeSinceLast    float64
}

func NewWorkBench1(file *os.File) *WorkBench1 {
	return &WorkBench1{
		Scanner: bufio.NewScanner(file),
	}
}

func (i *WorkBench1) Updatewait() {
	if !i.Started {
		i.Started = true
	}
	i.TimeSinceLast = 0.000
	if i.Scanner.Scan() {
		scanText := strings.Trim(i.Scanner.Text(), " ")
		conv, _ := strconv.ParseFloat(scanText, 64)
		i.WaitTime = conv
		i.ComponentMade = true
		i.TotalArrivalTime = i.TotalArrivalTime + conv
	} else {
		i.Closed = true
	}
}

func (i *WorkBench1) Time() {
	if i.WaitTime > 0 {
		i.WaitTime = i.WaitTime - .001
	}
}

func (i *WorkBench1) UpdateBlockTime() {
	i.Blocked = true
	i.BlockedTime = i.BlockedTime + .001
}

func (i *WorkBench1) CanMake() bool {
	if len(i.Buffer) > 0 {
		i.Blocked = false
		return true
	}
	return false

}

func (i *WorkBench1) UpdateAverageTimes() {
	if len(i.Buffer) == 0 {
		i.TotalTime0 = i.TotalTime0 + 0.001
	}
	if len(i.Buffer) == 1 {
		i.TotalTime1 = i.TotalTime1 + 0.001
	}
	if len(i.Buffer) == 2 {
		i.TotalTime2 = i.TotalTime2 + 0.001
	}
}

func (i *WorkBench1) TotalAverage(totalTime float64) float64 {
	var returnAverage float64
	average1 := i.TotalTime1 / totalTime
	returnAverage = returnAverage + (1 * average1)
	average2 := i.TotalTime2 / totalTime
	returnAverage = returnAverage + (2 * average2)
	return returnAverage
}
