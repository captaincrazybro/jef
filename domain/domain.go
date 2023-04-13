package domain

import (
	lu "github.com/captaincrazybro/literalutil"
)

// TODO: add documentation for interfaces
// TODO: change lu.String parameters to type string

// Jef interface to store instances of jef
type Jef interface {
	Moto()
	Check()
	Run()
	GetCompilerManager() CompilerManager
	GetVariableManager() VariableManager
	GetFunctionManager() FunctionManager
	GetDatatypeManager() DatatypeManager
	New(code string) Jef
	NewCodeless() Jef
}

// Compiler compilers interface
type Compiler interface {
	GetName() string
	Check(lu.String) bool
	Run(lu.String, *int) error
}

// Variable interface to store a variable
type Variable interface {
	GetType() Datatype
	GetValue() interface{}
	GetName() string
}

// Function interface to store a function
type Function interface {
	GetName() string
	GetReturnType() Datatype
	GetExec() func(Jef)
	GetParams() []Parameter
	RunExec([]interface{}, []Datatype, Jef) error
}

// Parameter interface to store a parameter
type Parameter interface {
	GetName() string
	GetType() Datatype
}

// Datatype interface to store a Datatype
type Datatype interface {
	GetName() string
	GetVarName() string
	Check(lu.String) bool
	GetValue(lu.String) (interface{}, error)
}

// DataValue interface to store a data value
type DataValue interface {
	GetType() Datatype
	GetTypeName() string
	GetValue() interface{}
}

// CompilerManager interface to store compilerManager instance
type CompilerManager interface {
	AddCompiler(Compiler)
	CompileLine(lu.String, *int) error
}

// VariableManager interface to store instance of variableManager
type VariableManager interface {
	RegisterVariable(string, Datatype, interface{}) error
	GetVariable(string) Variable
	GetVariables() []Variable
}

// FunctionManager interface to store instance of functionManager
type FunctionManager interface {
	RegisterFunction(string, Datatype, []Parameter, func(Jef)) error
	GetFunction(string) Function
}

// DatatypeManager interface to store instance of datatypeManager
type DatatypeManager interface {
	AddDatatype(Datatype)
	FindDatatype(lu.String) (DataValue, error)
	GetDatatype(string) Datatype
}
