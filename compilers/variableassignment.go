package compilers

import (
	"github.com/captaincrazybro/jef/domain"
	"github.com/captaincrazybro/jef/util"
	lu "github.com/captaincrazybro/literalutil"
)

// variableAssignment structure to compile variables
type variableAssignment struct {
	jef domain.Jef
}

func (v variableAssignment) GetName() string {
	return variableName
}

// TODO: make variableAssignment check more specific
func (v variableAssignment) Check(s lu.String) bool {
	return s.Contains("=") && s.Split("=").Len() >= 2
}

func (v variableAssignment) Run(iter *util.LineIterator) error {
	s := iter.Current()
	varName := s.Split("=")[0].ReplaceAll(" ", "")
	value := s.Split("=")[1]

	// Finds the datatype
	parser, err := v.jef.GetParserManager().ParseCode(value)
	if err != nil {
		return err
	}

	err = v.jef.GetVariableManager().RegisterVariable(varName.Tos(), parser.GetType(), parser.GetValue())
	if err != nil {
		return err
	}

	return nil
}
