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
	GetParserManager() ParserManager
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
	GetType() DataType
	GetValue() interface{}
	GetName() string
}

// Function interface to store a function
type Function interface {
	GetName() string
	GetReturnType() TypeParser
	GetExec() func(Jef)
	GetParams() []Parameter
	RunExec([]DataValue, Jef) error
}

// Parameter interface to store a parameter
type Parameter interface {
	GetName() string
	GetType() DataType
}

// TypeParser interface to store a TypeParser
type TypeParser interface {
	GetType() DataType
	Check(lu.String) bool
	GetValue(lu.String) (DataValue, error)
}

// DataType interface to store a datatype
type DataType interface {
	GetName() string
	// TODO: Implement a method to get relevant typeparsers for the given datatype
}

// DataValue interface to store a data value
type DataValue interface {
	GetType() DataType
	GetValue() interface{}
}

// CompilerManager interface to store compilerManager instance
type CompilerManager interface {
	AddCompiler(Compiler)
	CompileLine(lu.String, *int) error
}

// VariableManager interface to store instance of variableManager
type VariableManager interface {
	RegisterVariable(string, DataType, interface{}) error
	GetVariable(string) Variable
	GetVariables() []Variable
}

// FunctionManager interface to store instance of functionManager
type FunctionManager interface {
	RegisterFunction(string, TypeParser, []Parameter, func(Jef)) error
	GetFunction(string) Function
}

// DatatypeManager interface to store instance of datatypeManager
type DatatypeManager interface {
	AddDataType(DataType)
	GetDatatype(string) DataType
}

// ParserManager interface to store instance of parserManager
type ParserManager interface {
	AddParser(TypeParser)
	ParseCode(lu.String) (DataValue, error)
}
