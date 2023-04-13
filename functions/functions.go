package functions

import (
	"fmt"
	"github.com/captaincrazybro/jef/domain"
)

// functionManager structure to store functions
type functionManager struct {
	functions []function
	jef       domain.Jef
}

// New creates a new instance of functionManager
func New(j domain.Jef) domain.FunctionManager {
	fm := &functionManager{functions: []function{}, jef: j}
	fm.registerFunctions(j)
	return fm
}

// RegisterFunction registers a new function
func (fm *functionManager) RegisterFunction(name string, funcType domain.Datatype, params []domain.Parameter, exec func(jef domain.Jef)) error {
	if fm.GetFunction(name) != nil {
		return fmt.Errorf("bad function declaration, function %q has already been declared", name)
	}

	function := function{
		name:       name,
		returnType: funcType,
		exec:       exec,
		params:     params,
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

// validateParameters validates the parameters with given parameter values
func validateParameters(f domain.Function, values []interface{}, givenTypes []domain.Datatype) error {
	if len(values) != len(givenTypes) {
		return fmt.Errorf("an internal error has occured! length of the values does not equal the length of the given data types")
	} else if len(values) != len(f.GetParams()) {
		return fmt.Errorf("invalid parameters passed to function %s. number of parameters of parameters passed (%d) does not equal number of parameters of the function (%d)", f.GetName(), len(values), len(f.GetParams()))
	}

	// Checks the parameters, making sure they are the right datatypes
	for i, param := range f.GetParams() {
		givenType := givenTypes[i]
		if param.GetType() != givenType {
			return fmt.Errorf("invalid parameters passed to function %s. type of parameter %d (%s) does not equal the expected parameter type (%s)", f.GetName(), i+1, givenType.GetName(), param.GetType().GetName())
		}
	}

	return nil
}
