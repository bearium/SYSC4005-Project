package component

import (
	"bufio"
	"os"
	"time"
)

type Component struct {
	Scanner   *bufio.Scanner
	File      *os.File
	Name      string
	Id        int
	QueueData *Componentqueue
	TimeIn    time.Time
	TimeOut   time.Time
}

type Componentqueue struct {
	CurrentQueueSize    int
	ItemsInQueue        map[int]time.Duration
	TimeSinceLastUpdate time.Time
}

func NewComponent(name string, id int, file *os.File) *Component {
	return &Component{
		Name: name,
		Id:   id,
		File: file,
		QueueData: &Componentqueue{
			ItemsInQueue: make(map[int]time.Duration),
		},
	}
}

func (c *Component) AddScanner(scanner *bufio.Scanner) {
	c.Scanner = scanner
}
