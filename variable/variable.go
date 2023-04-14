package variable

import "github.com/captaincrazybro/jef/domain"

// variable structure to store a Variable
type variable struct {
	variableType domain.TypeParser
	value interface{}
	name string
}

// createVariable creates a variable structure
func createVariable(name string, varType domain.TypeParser, value interface{}) variable {
	return variable{
		variableType: varType,
		value:        value,
		name: name,
	}
}

// GetType gets the variable type
func (v variable) GetType() domain.TypeParser {
	return v.variableType
}

// GetValue gets the variable value
func (v variable) GetValue() interface{} {
	return v.value
}

// GetName gets the variable name
func (v variable) GetName() string {
	return v.name
}