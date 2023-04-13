package compilers

import (
	"github.com/captaincrazybro/jef/domain"
	"github.com/captaincrazybro/jef/util"
	lu "github.com/captaincrazybro/literalutil"
)

// variable structure to compile variables
type variable struct {
	jef domain.Jef
}

func (v variable) GetName() string {
	return variableName
}

// TODO: make variable check more specific
func (v variable) Check(s lu.String) bool {
	return s.Contains("=") && s.Split("=").Len() >= 2
}

func (v variable) Run(s lu.String, line *int) error {
	varName := s.Split("=")[0].ReplaceAll(" ", "")
	value := s.Split("=")[1]
	value = util.TrimWhitespaces(value)

	// Finds the datatype
	dataType, _ := v.jef.GetDatatypeManager().FindDatatype(s)

	err := v.jef.GetVariableManager().RegisterVariable(varName.Tos(), dataType, value)
	if err != nil {
		return err
	}

	return nil
}
