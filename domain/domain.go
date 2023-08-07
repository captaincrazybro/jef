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
	Run() error
	GetCompilerManager() CompilerManager
	GetVariableManager() VariableManager
	GetFunctionManager() FunctionManager
	GetDatatypeManager() DatatypeManager
	GetParserManager() ParserManager
	New([]lu.String) Jef
	NewFromCode(string) Jef
	NewCodeless() Jef
}

// Compiler compilers interface
type Compiler interface {
	GetName() string
	Check(lu.String) bool
	Run(LineIterator) error
}

// Variable interface to store a variable
type Variable interface {
	GetType() DataType
	GetValue() interface{}
	GetName() string
	UpdateValue(interface{})
}

// Function interface to store a function
type Function interface {
	GetName() string
	GetReturnType() TypeParser
	GetExec() func()
	GetParams() []Parameter
	RunExec([]DataValue) error
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
	CompileLine(LineIterator) error
}

// VariableManager interface to store instance of variableManager
type VariableManager interface {
	RegisterVariable(string, DataType, interface{}) bool
	UpdateVariable(string, DataType, interface{}) error
	DeleteVariable(string) error
	GetVariable(string) Variable
	GetVariables() []Variable
	Copy(newJ Jef) VariableManager
}

// FunctionManager interface to store instance of functionManager
type FunctionManager interface {
	RegisterFunction(string, Jef, TypeParser, []Parameter, func()) error
	GetFunction(string) Function
	Copy(newJ Jef) FunctionManager
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

// LineIterator interface to store instance of LineBreak in utils
type LineIterator interface {
	New(newLines []lu.String)
	Next() bool
	Current() lu.String
	Index() int
	Lines() []lu.String
	GoToLine(i int)
	EditCurrent(s lu.String)
}
