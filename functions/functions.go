package functions

import (
	"fmt"
	"github.com/captaincrazybro/jef/domain"
)

// functionManager structure to store functions
type functionManager struct {
	functions []function
	jef domain.Jef
}

// New creates a new instance of functionManager
func New(j domain.Jef) domain.FunctionManager {
	return &functionManager{functions: []function{}, jef: j}
}

// RegisterFunction registers a new function
func (fm *functionManager) RegisterFunction(name string, funcType string, exec func(jef domain.Jef)) error {
	if fm.GetFunction(name) != nil {
		return fmt.Errorf("bad function declaration, function %q has already been declared", name)
	}

	function := function{
		name:     name,
		funcType: funcType,
		exec:     exec,
	}
	fm.functions = append(fm.functions, function)
	return nil
}

// GetFunction gets a function
func (fm functionManager) GetFunction(name string) domain.Function {
	for i := 0; i < len(fm.functions); i++ {
		function := fm.functions[i]
		if function.name == name {
			return function
		}
	}
	return nil
}