package typeparsers

import (
	"fmt"
	"github.com/captaincrazybro/jef/domain"
	"github.com/captaincrazybro/jef/util"
	lu "github.com/captaincrazybro/literalutil"
)

var inequalityOperators = []string{"<", ">", ">=", "<="}

type Inequality struct {
	jef domain.Jef
}

func (iD Inequality) GetName() string {
	return domain.InequalityParserName
}

func (iD Inequality) GetType() domain.DataType {
	return iD.jef.GetDatatypeManager().GetDatatype("bool")
}

func (iD Inequality) Check(s lu.String) bool {
	return util.LineHasOperators(s, inequalityOperators)
}

func (iD Inequality) GetValue(s lu.String) (domain.DataValue, error) {
	parenthesisCount := 0
	isQuote := false

	var operand1, operator, operand2 lu.String

	// Loops through the current string to see if it contains an isolated inequality operator
	for i, r := range s {
		if r == '"' && (i == 0 || s[i-1] != '\\') {
			isQuote = !isQuote
		}

		if !isQuote {
			if r == '(' {
				parenthesisCount++
			} else if r == ')' {
				parenthesisCount--
			} else if parenthesisCount == 0 {
				endIndex := i + 1
				if i+1 != len(s) && s[i+1] != ' ' {
					endIndex++
				}

				if util.StringArrContains(inequalityOperators, string(s[i:endIndex])) {
					operand1 = util.TrimWhitespaces(s[:i])
					operator = util.TrimWhitespaces(s[i:endIndex])
					operand2 = util.TrimWhitespaces(s[endIndex:])
				}
			}
		}
	}

	// Makes sure there is something on both sides of the operator
	if operand1 == "" {
		return nil, fmt.Errorf("invalid inequality expression! nothing found on left side of operator")
	}
	if operand2 == "" {
		return nil, fmt.Errorf("invalid inequality expression! nothing found on right side of operator")
	}

	// Parses both operands
	operand1Val, err := iD.jef.GetParserManager().ParseCode(operand1)
	if err != nil {
		return nil, err
	}
	operand2Val, err := iD.jef.GetParserManager().ParseCode(operand2)
	if err != nil {
		return nil, err
	}

	// Makes sure that the operands are both either an integer or a double
	if operand1Val.GetType().GetName() != domain.IntDataTypeName && operand1Val.GetType().GetName() != domain.DoubleDataTypeName {
		return nil, fmt.Errorf("invalid inequality expression! left expression of operator does not evaluate to int or double")
	}
	if operand2Val.GetType().GetName() != domain.IntDataTypeName && operand2Val.GetType().GetName() != domain.DoubleDataTypeName {
		return nil, fmt.Errorf("invalid inequality expression! right expression of operator does not evaluate to int or double")
	}

	// Evaluates the result of inequality operation
	boolResult := evalInequalityOperator(operand1Val, operator, operand2Val)

	resultData := dataValue{
		typeStruct: iD.jef.GetDatatypeManager().GetDatatype(domain.BooleanDataTypeName),
		value:      boolResult,
	}

	return resultData, nil
}

// evalInequalityOperator evaluates a single inequality operator
func evalInequalityOperator(operand1Val domain.DataValue, operator lu.String, operand2Val domain.DataValue) bool {
	// Converts the values to float64 type
	var val1, val2 float64
	if operand1Val.GetType().GetName() == domain.IntDataTypeName {
		val1 = float64(operand1Val.GetValue().(int))
	} else {
		val1 = operand1Val.GetValue().(float64)
	}
	if operand2Val.GetType().GetName() == domain.IntDataTypeName {
		val2 = float64(operand2Val.GetValue().(int))
	} else {
		val2 = operand2Val.GetValue().(float64)
	}

	// Evaluates the result of the operator depending on the opeartor
	switch operator {
	case "<":
		{
			return val1 < val2
			break
		}
	case ">":
		{
			return val1 > val2
			break
		}
	case "<=":
		{
			return val1 <= val2
			break
		}
	case ">=":
		{
			return val1 >= val2
			break
		}
	}

	return false
}
