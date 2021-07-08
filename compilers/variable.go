package compilers

import (
	"github.com/captaincrazybro/jef/variable"
	lu "github.com/captaincrazybro/literalutil"
	c "github.com/captaincrazybro/literalutil/console"
)

// Variable structure to compile variables
type Variable struct {}

func (v Variable) GetName() string {
	return variableName
}

// TODO: make variable check more specific
func (v Variable) Check(s lu.String) bool {
	return s.Contains("=") && s.Split("=").Len() >= 2
}

func (v Variable) Run(s lu.String, line *int) error {
	varName := s.Split("=")[0].ReplaceAll(" ", "")
	value := s.Split("=")[1].TrimPrefix(" ")
	// TODO: variable name validation
	variable.RegisterVariable(varName.Tos(), value)

	c.Plnf("%s and %v", varName, value)

	return nil
}