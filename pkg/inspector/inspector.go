package inspector

import (
	"bufio"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/SYSC4005-Project/pkg/component"
)

const seed = 42

type Inspector struct {
	Name       string
	Components []*component.Component
}

func NewInspector(name string, components []*component.Component) *Inspector {
	return &Inspector{
		Name:       name,
		Components: components,
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
		} else {
			i.Components = append(i.Components[:randInt], i.Components[randInt+1:]...)
			if len(i.Components) == 0 {
				return
			}
			continue
		}
	}
}
