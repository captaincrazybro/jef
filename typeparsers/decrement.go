package typeparsers

import (
	"fmt"
	"github.com/captaincrazybro/jef/domain"
	"github.com/captaincrazybro/jef/util"
	lu "github.com/captaincrazybro/literalutil"
	"regexp"
)

type Decrement struct {
	jef domain.Jef
}

func (d Decrement) GetName() string {
	return domain.DecrementParserName
}

func (d Decrement) GetType() domain.DataType {
	return nil
}

func (d Decrement) Check(s lu.String) bool {
	r1, _ := regexp.Compile("^--\\S*$")
	r2, _ := regexp.Compile("^\\S*--$")
	return r1.MatchString(s.Tos()) || r2.MatchString(s.Tos())
}

func (d Decrement) GetValue(s lu.String) (domain.DataValue, error) {
	s = util.TrimWhitespaces(s)
	// Parses variable name
	r1, _ := regexp.Compile("^--([a-zA-Z][a-zA-Z0-9_]*)$")
	r2, _ := regexp.Compile("^([a-zA-Z][a-zA-Z0-9_]*)--$")
	var varName string
	if r1.MatchString(s.Tos()) {
		varName = r1.FindStringSubmatch(s.Tos())[1]
	} else if r2.MatchString(s.Tos()) {
		varName = r2.FindStringSubmatch(s.Tos())[1]
	} else {
		return nil, fmt.Errorf("invalid decrement statement! variable name is invalid")
	}

	// Validates variable name
	if varName == "" {
		return nil, fmt.Errorf("invalid decrement statement! no variable specified")
	}

	variable := d.jef.GetVariableManager().GetVariable(varName)
	if variable == nil {
		return nil, fmt.Errorf("invalid decrement statement! variable does not exist")
	}

	if variable.GetType().GetName() != domain.IntDataTypeName && variable.GetType().GetName() != domain.DoubleDataTypeName {
		return nil, fmt.Errorf("invalid decrement statement! variable type must be int or double")
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
		result = variable.GetValue().(int) - 1
	} else {
		result = variable.GetValue().(float64) - 1.0
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
