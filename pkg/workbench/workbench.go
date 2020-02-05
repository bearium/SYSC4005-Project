package workbench

import (
	"github.com/SYSC4005-Project/pkg/inspector"
	"github.com/SYSC4005-Project/pkg/component"
)

type Workbench struct {
	Name string
	Inspectors []*inspector.Inspector
	//Product
	ComponentArray [][]*component.Component 
}