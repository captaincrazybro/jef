package datatypes

import (
	"fmt"
	"github.com/captaincrazybro/jef/domain"
	lu "github.com/captaincrazybro/literalutil"
	"regexp"
)

type Variable struct {
	jef domain.Jef
}

func (vD Variable) GetName() string {
	return "<NO_TYPE>"
}

func (vD Variable) GetVarName() string {
	return "variable"
}

func (vD Variable) Check(s lu.String) bool {
	r, _ := regexp.Compile("^[a-zA-Z][a-zA-Z0-9_]*$")
	return r.MatchString(s.Tos())
}

func (vD Variable) GetValue(s lu.String) (interface{}, error) {
	// Trys to find the variable
	v := vD.jef.GetVariableManager().GetVariable(s.Tos())
	if v == nil {
		return nil, fmt.Errorf("invalid variable called! variable does not exist")
	}

	return v.GetValue(), nil
}
