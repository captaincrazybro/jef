package functions

import (
	"fmt"
	"github.com/captaincrazybro/jef"
	"github.com/captaincrazybro/jef/domain"
)

// functionManager structure to store functions
type functionManager struct {
	functions []function
	jef domain.Jef
}

// New creates a new instance of functionManager
func New(j domain.Jef) domain.FunctionManager {
	fm := &functionManager{functions: []function{}, jef: j}
	fm.registerFunctions(j)
	return fm
}

// RegisterFunction registers a new function
func (fm *functionManager) RegisterFunction(name string, funcType string, params map[string]domain.Datatype, exec func(jef domain.Jef)) error {
	if fm.GetFunction(name) != nil {
		return fmt.Errorf("bad function declaration, function %q has already been declared", name)
	}

	function := function{
		name:     name,
		funcType: funcType,
		exec:     exec,
		params: params,
	}
	fm.functions = append(fm.functions, function)
	return nil
}

// GetFunction gets a function
func (fm *functionManager) GetFunction(name string) domain.Function {
	for i := 0; i < len(fm.functions); i++ {
		function := fm.functions[i]
		if function.name == name {
			return function
		}
	}
	return nil
}

// PrepJef adds function parameters to a n instance of jef
func (fm *functionManager) PrepJef(code string, varVals []string, f domain.Function) (domain.Jef, error) {
	if len(varVals) != len(f.GetParams()) {
		return nil, fmt.Errorf("incorrect number of parameters, expected %d got %d", len(f.GetParams()), len(varVals))
	}

	j := jef.NewFromExisting(fm.jef, code)

	i := 0
	for k, v := range f.GetParams() {
		err := j.GetVariableManager().RegisterVariable(k, v, varVals[i])
		if err != nil {
			return nil, err
		}
		i++
	}

	return j, nil
}