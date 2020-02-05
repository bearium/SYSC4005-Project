package component

import (
	"bufio"
	"os"
)

type Component struct {
	Scanner *bufio.Scanner
	File    *os.File
	Name    string
	Id      int
}

func NewComponent(name string, id int, file *os.File) *Component {
	return &Component{
		Name: name,
		Id:   id,
		File: file,
	}
}

func (c *Component) AddScanner(scanner *bufio.Scanner) {
	c.Scanner = scanner
}
