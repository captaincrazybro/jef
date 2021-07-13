package jef

import (
	"github.com/captaincrazybro/jef/compilers"
	"github.com/captaincrazybro/jef/domain"
	"github.com/captaincrazybro/jef/functions"
	"github.com/captaincrazybro/jef/variable"
	lu "github.com/captaincrazybro/literalutil"
	c "github.com/captaincrazybro/literalutil/console"
)

func init() {
	c.SetErrorPrefix("Error: ")
}

type jef struct {
	code      lu.String
	//managers managers.Managers
	compilers domain.CompilerManager
	variables domain.VariableManager
	functions domain.FunctionManager
}

// New creates a new instance of Jef
func New(code string) domain.Jef {
	j := jef {
		code: lu.String(code),
	}

	registerManagers(&j)

	return j
}

// registerManagers registers all the managers
func registerManagers(j *jef) {
	j.compilers = compilers.New(j)
	j.variables = variable.New(j)
	j.functions = functions.New(j)
}

// Moto prints the moto for the Jef programming language
func (_ jef) Moto() {
	c.Pln("My name is Jeff!")
}

// Check checks the code for errors and pre registers functions
// TODO: create checking stuff
func (j jef) Check() {
	//code := lu.String(s)
}

// Run runs the code
func (j jef) Run() {
	code := j.code
	lines := code.Split("\n")

	for i := 0; i < lines.Len(); i ++ {
		lineValue := lines[i]
		err := j.GetCompilerManager().CompileLine(lineValue, &i)
		if err != nil {
			c.Fln(err)
		}
	}
}

func (j jef) GetCompilerManager() domain.CompilerManager {
	return j.compilers
}

func (j jef) GetVariableManager() domain.VariableManager {
	return j.variables
}

func (j jef) GetFunctionManager() domain.FunctionManager {
	return j.functions
}

// GetManagers gets the managers
//func (j jef) GetManagers() managers.Managers {
//	return j.managers
//}