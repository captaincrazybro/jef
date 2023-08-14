package jef

import (
	"fmt"
	"github.com/captaincrazybro/jef/compilers"
	"github.com/captaincrazybro/jef/datatypes"
	"github.com/captaincrazybro/jef/domain"
	"github.com/captaincrazybro/jef/functions"
	"github.com/captaincrazybro/jef/typeparsers"
	"github.com/captaincrazybro/jef/util"
	"github.com/captaincrazybro/jef/variable"
	lu "github.com/captaincrazybro/literalutil"
	c "github.com/captaincrazybro/literalutil/console"
)

func init() {
	c.SetErrorPrefix("Error: ")
}

type jef struct {
	lines     []lu.String
	isSubJef  bool
	compilers domain.CompilerManager
	variables domain.VariableManager
	functions domain.FunctionManager
	dataTypes domain.DatatypeManager
	parsers   domain.ParserManager
}

// NewFromCode creates a new instance of Jef
func NewFromCode(code string) domain.Jef {
	lines := lu.String(code).Split("\n")
	return New(lines)
}

// New creates a new instance of jef with a list of lines
func New(lines []lu.String) domain.Jef {
	j := jef{
		lines: lines,
	}

	registerManagers(&j)

	return &j
}

// registerManagers registers all the managers
func registerManagers(j *jef) {
	j.compilers = compilers.New(j)
	j.variables = variable.New(j)
	j.dataTypes = datatypes.New(j)
	j.parsers = typeparsers.New(j)
	j.functions = functions.New(j)
}

// Moto prints the moto for the Jef programming language
func (_ *jef) Moto() {
	c.Pln("My name is Jeff!")
}

// Check checks the code for errors and pre registers functions
// TODO: create checking stuff
func (j *jef) Check() {
	//code := lu.String(s)
}

// Run runs the code
func (j *jef) Run() error {
	iter := util.LineIterator{}
	iter.New(j.lines)

	for iter.Next() {
		err := j.GetCompilerManager().CompileLine(&iter)
		if err != nil {
			if j.isSubJef {
				return fmt.Errorf("%s", err)
			} else {
				return fmt.Errorf("%s - line: %d", err, iter.Index()+1)
			}
		}
	}

	return nil
}

func (j *jef) GetCompilerManager() domain.CompilerManager {
	return j.compilers
}

func (j *jef) GetVariableManager() domain.VariableManager {
	return j.variables
}

func (j *jef) GetFunctionManager() domain.FunctionManager {
	return j.functions
}

func (j *jef) GetDatatypeManager() domain.DatatypeManager {
	return j.dataTypes
}

func (j *jef) GetParserManager() domain.ParserManager {
	return j.parsers
}

// NewFromCode creates a new instance of Jef based on an existing
func (j *jef) NewFromCode(code string) domain.Jef {
	lines := lu.String(code).Split("\n")
	return j.New(lines)
}

// New creates a new instance of Jef based on a list of lines
func (j *jef) New(lines []lu.String) domain.Jef {
	newJef := &jef{
		lines:     lines,
		isSubJef:  true,
		compilers: j.GetCompilerManager(),
		dataTypes: j.GetDatatypeManager(),
		parsers:   j.GetParserManager(),
	}

	newJef.functions = j.GetFunctionManager().Copy(newJef)
	newJef.variables = j.GetVariableManager().Copy(newJef)

	return newJef
}

// NewCodeless creates a new instance of Jef based on an existing, without code
// This is only used for system functions that do not need to run code
func (j *jef) NewCodeless() domain.Jef {
	newJef := &jef{
		isSubJef:  true,
		compilers: j.GetCompilerManager(),
		dataTypes: j.GetDatatypeManager(),
		parsers:   j.GetParserManager(),
	}

	newJef.functions = j.GetFunctionManager().Copy(newJef)
	newJef.variables = j.GetVariableManager().Copy(newJef)

	return newJef
}
