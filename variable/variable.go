package variable

// variable structure to store a Variable
type variable struct {
	variableType string
	value interface{}
	name string
}

// createVariable creates a variable structure
func createVariable(name string, varType string, value interface{}) variable {
	return variable{
		variableType: varType,
		value:        value,
		name: name,
	}
}

// GetType gets the variable type
func (v variable) GetType() string {
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