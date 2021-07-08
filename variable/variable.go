package variable

var variables []variable

// Init initializes the variable array
func Init() {
	variables = []variable{}
}

// RegisterVariable registers a variable
func RegisterVariable(varType string, value interface{}) {
	v := createVariable(varType, value)
	variables = append(variables, v)
}

// createVariable creates a variable structure
func createVariable(varType string, value interface{}) variable {
	return variable{
		variableType: varType,
		value:        value,
	}
}

type variable struct {
	variableType string
	value interface{}
}

func (v variable) GetVariableType() string {
	return v.variableType
}

func (v variable) GetValue() interface{} {
	return v.value
}