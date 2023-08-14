package typeparsers

import (
	"fmt"
	"github.com/captaincrazybro/jef/domain"
	"github.com/captaincrazybro/jef/util"
	lu "github.com/captaincrazybro/literalutil"
	"regexp"
)

type Increment struct {
	jef domain.Jef
}

func (i Increment) GetName() string {
	return domain.IncrementParserName
}

func (i Increment) GetType() domain.DataType {
	return nil
}

func (i Increment) Check(s lu.String) bool {
	r1, _ := regexp.Compile("^\\+\\+\\S*$")
	r2, _ := regexp.Compile("^\\S*\\+\\+$")
	return r1.MatchString(s.Tos()) || r2.MatchString(s.Tos())
}

func (i Increment) GetValue(s lu.String) (domain.DataValue, error) {
	s = util.TrimWhitespaces(s)
	// Parses variable name
	r1, _ := regexp.Compile("^\\+\\+(\\S*)$")
	r2, _ := regexp.Compile("^(\\S*)\\+\\+$")
	var varName string
	if r1.MatchString(s.Tos()) {
		varName = r1.FindStringSubmatch(s.Tos())[1]
	} else {
		varName = r2.FindStringSubmatch(s.Tos())[1]
	}

	// Validates variable name
	if varName == "" {
		return nil, fmt.Errorf("invalid increment statement! no variable specified")
	}

	variable := i.jef.GetVariableManager().GetVariable(varName)
	if variable == nil {
		return nil, fmt.Errorf("invalid increment statement! variable does not exist")
	}

	if variable.GetType().GetName() != domain.IntDataTypeName && variable.GetType().GetName() != domain.DoubleDataTypeName {
		return nil, fmt.Errorf("invalid increment statement! variable type must be int or double")
	}

	// If post-increment, creates dataValue instance before increment
	var val dataValue
	if r2.MatchString(s.Tos()) {
		val = dataValue{
			typeStruct: variable.GetType(),
			value:      variable.GetValue(),
		}
	}

	var result interface{}
	// Increments the variable
	if variable.GetType().GetName() == domain.IntDataTypeName {
		result = variable.GetValue().(int) + 1
	} else if r2.MatchString(s.Tos()) {
		varName = r2.FindStringSubmatch(s.Tos())[1]
	} else {
		return nil, fmt.Errorf("invalid increment statement! variable name is invalid")
	}

	// Updates the variable value
	variable.UpdateValue(result)

	// if pre-increment, creates dataValue instance after increment
	if r1.MatchString(s.Tos()) {
		val = dataValue{
			typeStruct: variable.GetType(),
			value:      variable.GetValue(),
		}
	}

	return val, nil
}
