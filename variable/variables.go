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

// Copy creates a copy of the current variableManager instance
func (vz *variableManager) Copy(newJ domain.Jef) domain.VariableManager {
	newVz := &variableManager{variables: vz.variables, jef: newJ}
	return newVz
}

// RegisterVariable registers a variable, except if the variable already exists and then returns true
func (vz *variableManager) RegisterVariable(varName string, varType domain.DataType, value interface{}) bool {
	// TODO: add variable name validation
	if vz.GetVariable(varName) != nil {
		return true
	}

	v := createVariable(varName, varType, value)
	vz.variables = append(vz.variables, v)
	return false
}

// UpdateVariable updates the value of a variable
func (vz *variableManager) UpdateVariable(varName string, newType domain.DataType, newValue interface{}) error {
	v := vz.GetVariable(varName)
	// Makes sure the variable exists
	if v == nil {
		return fmt.Errorf("internal error! tried to update variable that does not exist")
	}

	// Validates the type
	if newType != v.GetType() {
		return fmt.Errorf("invalid variable reassignment! new type does not match original type of variable")
	}

	v.UpdateValue(newValue)
	return nil
}

// DeleteVariable deletes a variable
func (vz *variableManager) DeleteVariable(varName string) error {
	if vz.GetVariable(varName) == nil {
		return fmt.Errorf("invalid variable deletion! variable does not exist")
	}

	// Finds the index of the variable name
	isFound := false
	i := 0
	for i < len(vz.variables) && !isFound {
		if vz.variables[i].GetName() == varName {
			isFound = true
		} else {
			i++
		}
	}

	vz.variables[i] = vz.variables[len(vz.variables)-1]
	vz.variables = vz.variables[:len(vz.variables)-1]
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
