package compilers

import (
	"fmt"
	"github.com/captaincrazybro/jef/domain"
	lu "github.com/captaincrazybro/literalutil"
	"regexp"
)

// mathAssignment structure to compile variables
type mathAssignment struct {
	jef domain.Jef
}

func (mA mathAssignment) GetName() string {
	return variableName
}

func (mA mathAssignment) Check(s lu.String) bool {
	reg, _ := regexp.Compile("^\\S* *(\\+=|-=|\\*=|/=).*$")
	return reg.MatchString(s.Tos())
}

func (mA mathAssignment) Run(iter domain.LineIterator) error {
	s := iter.Current()
	// Parses the varName and ending expression
	reg, _ := regexp.Compile("^([a-zA-Z][a-zA-Z0-9_]*) *(\\+=|-=|\\*=|/=) *(.*)$")
	if !reg.MatchString(s.Tos()) {
		return fmt.Errorf("invalid math assignment statement! variable has invalid name")
	}

	subs := reg.FindStringSubmatch(s.Tos())
	varName := subs[1]
	op := subs[2]
	expr := subs[3]

	// Validates the variable
	if varName == "" {
		return fmt.Errorf("invalid math assignment statement! no variable specified")
	}

	variable := mA.jef.GetVariableManager().GetVariable(varName)
	if variable == nil {
		return fmt.Errorf("invalid math assignment statement! variable does not exist")
	}

	if variable.GetType().GetName() != domain.IntDataTypeName && variable.GetType().GetName() != domain.DoubleDataTypeName {
		return fmt.Errorf("invalid math assignment statement! variable type must be int or double")
	}

	// Parses expression
	exprData, err := mA.jef.GetParserManager().ParseCode(lu.String(expr))
	if err != nil {
		return err
	}

	// Makes sure that the evaluated expression type is int or double
	if exprData.GetType().GetName() != domain.IntDataTypeName && exprData.GetType().GetName() != domain.DoubleParserName {
		return fmt.Errorf("invalid math assignment statement! operation expression must be of type int or double")
	}

	performMathAssignmentOperation(variable, op, exprData)
	return nil
}

// performMathAssignmentOperation performs the math assignment operation based on the string operator and variable and expression types
func performMathAssignmentOperation(variable domain.Variable, op string, exprData domain.DataValue) {
	// Converts the variable and expression value appropriately
	var val1 float64
	var val2 float64
	if variable.GetType().GetName() == domain.IntDataTypeName {
		val1 = float64(variable.GetValue().(int))
	} else {
		val1 = variable.GetValue().(float64)
	}
	if exprData.GetType().GetName() == domain.IntDataTypeName {
		val2 = float64(exprData.GetValue().(int))
	} else {
		val2 = exprData.GetValue().(float64)
	}

	// Computes the operation
	var result float64
	switch op {
	case "+=":
		{
			result = val1 + val2
			break
		}
	case "-=":
		{
			result = val1 - val2
			break
		}
	case "*=":
		{
			result = val1 * val2
			break
		}
	case "/=":
		{
			result = val1 / val2
			break
		}
	}

	// Updates the variable value based on the type of the variable
	if variable.GetType().GetName() == domain.DoubleDataTypeName {
		variable.UpdateValue(result)
	} else {
		variable.UpdateValue(int(result))
	}
}
