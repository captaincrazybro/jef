package functions

import (
	"fmt"
	"github.com/captaincrazybro/jef/domain"
	lu "github.com/captaincrazybro/literalutil"
)

// functionManager structure to store functions
type functionManager struct {
	functions []domain.Function
	jef       domain.Jef
}

// New creates a new instance of functionManager
func New(j domain.Jef) domain.FunctionManager {
	fm := &functionManager{functions: []domain.Function{}, jef: j}
	fm.registerFunctions(j)
	return fm
}

// Copy creates a copy of the instance of functionManager
func (fm *functionManager) Copy(newJ domain.Jef) domain.FunctionManager {
	newFm := &functionManager{functions: fm.functions, jef: newJ}
	return newFm
}

// RegisterFunction registers a regular function
func (fm *functionManager) RegisterFunction(name string, jef domain.Jef, returnType domain.DataType, params []domain.Parameter, lines []lu.String) error {
	if fm.GetFunction(name) != nil {
		return fmt.Errorf("bad function declaration! function %q has already been declared", name)
	}

	fun := function{
		name:       name,
		jef:        jef,
		returnType: returnType,
		lines:      lines,
		params:     params,
	}
	fm.functions = append(fm.functions, fun)
	return nil
}

// registerSysFunction registers a new sysFunction
func (fm *functionManager) registerSysFunction(name string, jef domain.Jef, returnType domain.DataType, params []domain.Parameter, exec func(domain.Jef) error) {
	if fm.GetFunction(name) != nil {
		fmt.Printf("bad sysFunction declaration, sysFunction %q has already been declared\n", name)
	}

	fun := sysFunction{
		name:       name,
		jef:        jef,
		returnType: returnType,
		exec:       exec,
		params:     params,
	}
	fm.functions = append(fm.functions, fun)
}

// GetFunction gets a sysFunction
func (fm *functionManager) GetFunction(name string) domain.Function {
	for i := 0; i < len(fm.functions); i++ {
		function := fm.functions[i]
		if function.GetName() == name {
			return function
		}
	}
	return nil
}

// CreateParameter creates a parameter for a sysFunction
func (fm *functionManager) CreateParameter(name string, paramType domain.DataType) domain.Parameter {
	return parameter{name, paramType}
}

// validateParameters validates the parameters with given parameter values
func validateParameters(f domain.Function, values []domain.DataValue, j domain.Jef) error {
	if len(values) != len(f.GetParams()) {
		return fmt.Errorf("invalid parameters passed to sysFunction %s. number of parameters of parameters passed (%d) does not equal number of parameters of the sysFunction (%d)", f.GetName(), len(values), len(f.GetParams()))
	}

	// Checks the parameters, making sure they are the right datatypes
	for i, param := range f.GetParams() {
		givenType := values[i].GetType()
		if param.GetType() != j.GetDatatypeManager().GetDatatype("any") && param.GetType() != givenType {
			return fmt.Errorf("invalid parameters passed to sysFunction %s. type of parameter %d (%s) does not equal the expected parameter type (%s)", f.GetName(), i+1, givenType.GetName(), param.GetType().GetName())
		}
	}

	return nil
}
