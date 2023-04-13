package variable

import (
	"fmt"
	"github.com/captaincrazybro/jef/domain"
)

// variableManager structure to store instance of VariableManager
type variableManager struct {
	variables []domain.Variable
	jef       domain.Jef
}

// New creates a new VariableManager instance
func New(j domain.Jef) domain.VariableManager {
	return &variableManager{[]domain.Variable{}, j}
}

// RegisterVariable registers a variable
func (vz *variableManager) RegisterVariable(varName string, varType domain.Datatype, value interface{}) error {
	// TODO: add variable name validation
	if vz.GetVariable(varName) != nil {
		return fmt.Errorf("bad variable initialization, the variable %q has already been declared", varName)
	}

	v := createVariable(varName, varType, value)
	vz.variables = append(vz.variables, v)
	return nil
}

// GetVariable gets a variable
func (vz *variableManager) GetVariable(name string) domain.Variable {
	for i := 0; i < len(vz.variables); i++ {
		variable := vz.variables[i]
		if variable.GetName() == name {
			return variable
		}
	}
	return nil
}

// GetVariables gets the list of variables
func (vz *variableManager) GetVariables() []domain.Variable {
	return vz.variables
}
