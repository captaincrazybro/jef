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
}

// Compiler compilers interface
type Compiler interface {
	GetName() string
	Check(lu.String) bool
	Run(lu.String, *int) error
}

// Variable interface to store a variable
type Variable interface {
	GetType() string
	GetValue() interface{}
	GetName() string
}

// Function interface to store a function
type Function interface {
	GetName() string
	GetType() string
	GetExec() func(Jef)
}

// Datatype interface to store a datatype
type Datatype interface {
	GetName() string
	Check()
}

// CompilerManager interface to store compilerManager instance
type CompilerManager interface {
	AddCompiler(Compiler)
	CompileLine(lu.String, *int) error
}

// VariableManager interface to store instance of variableManager
type VariableManager interface {
	RegisterVariable(string, string, interface{}) error
	GetVariable(string) Variable
}

// FunctionManager interface to store instance of functionManager
type FunctionManager interface {
	RegisterFunction(string, string, func(Jef)) error
	GetFunction(string) Function
}

// DatatypeManager interface to store instance of datatypeManager
type DatatypeManager interface {
	AddDatatype(Datatype)
	FindDatatype(string) (Datatype, error)
}