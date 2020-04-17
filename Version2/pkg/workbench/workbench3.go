package workbench

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Workbench3 struct {
	Scanner         *bufio.Scanner
	WaitTime        float64
	BlockedTime     float64
	BufferFull      bool
	Buffer          []int
	Buffer2         []int
	Started         bool
	ComponentMade   bool
	TotalComponents int
	Closed          bool
	Blocked         bool
	TotalTime0      float64
	TotalTime1      float64
	TotalTime2      float64
	TotalTime3      float64
	TotalTime4      float64
	TotalIn         int
}

func NewWorkBench3(file *os.File) *Workbench3 {
	return &Workbench3{
		Scanner: bufio.NewScanner(file),
	}
}

func (i *Workbench3) Updatewait() {
	if !i.Started {
		i.Started = true
	}
	if i.Scanner.Scan() {
		scanText := strings.Trim(i.Scanner.Text(), " ")
		conv, _ := strconv.ParseFloat(scanText, 64)
		i.WaitTime = conv
		i.ComponentMade = true
	} else {
		i.Closed = true
	}
}

func (i *Workbench3) Time() {
	if i.WaitTime > 0 {
		i.WaitTime = i.WaitTime - .001
	}
}

func (i *Workbench3) UpdateBlockTime() {
	i.Blocked = true
	i.BlockedTime = i.BlockedTime + .001
}

func (i *Workbench3) CanMake() bool {
	if len(i.Buffer) > 0 && len(i.Buffer2) > 0 {
		i.Blocked = false
		return true
	}
	return false

}

func (i *Workbench3) UpdateAverageTimes() {
	total := len(i.Buffer) + len(i.Buffer2)

	if total == 0 {
		i.TotalTime0 = i.TotalTime0 + 0.001
	}
	if total == 1 {
		i.TotalTime1 = i.TotalTime1 + 0.001
	}
	if total == 2 {
		i.TotalTime2 = i.TotalTime2 + 0.001
	}
	if total == 3 {
		i.TotalTime3 = i.TotalTime3 + 0.001
	}
	if total == 4 {
		i.TotalTime4 = i.TotalTime4 + 0.001
	}
}

func (i *Workbench3) TotalAverage(totalTime float64) float64 {
	var returnAverage float64
	average1 := i.TotalTime1 / totalTime
	returnAverage = returnAverage + (1 * average1)
	average2 := i.TotalTime0 / totalTime
	returnAverage = returnAverage + (2 * average2)
	average3 := i.TotalTime3 / totalTime
	returnAverage = returnAverage + (3 * average3)
	average4 := i.TotalTime4 / totalTime
	returnAverage = returnAverage + (4 * average4)
	return returnAverage
}
