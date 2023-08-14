package typeparsers

import (
	"fmt"
	"github.com/captaincrazybro/jef/domain"
	"github.com/captaincrazybro/jef/util"
	lu "github.com/captaincrazybro/literalutil"
)

type Equals struct {
	jef domain.Jef
}

func (eD Equals) GetName() string {
	return domain.EqualsParserName
}

func (eD Equals) GetType() domain.DataType {
	return eD.jef.GetDatatypeManager().GetDatatype(domain.BooleanDataTypeName)
}

func (eD Equals) Check(s lu.String) bool {
	return util.LineHasOperators(s, []string{"=="})
}

func (eD Equals) GetValue(s lu.String) (domain.DataValue, error) {
	parenthesisCount := 0
	isQuote := false

	var operand1, operand2 lu.String

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
				if i+1 != len(s) {
					if string(s[i:i+2]) == "==" {
						operand1 = util.TrimWhitespaces(s[:i])
						operand2 = util.TrimWhitespaces(s[i+2:])
					}
				}
			}
		}
	}

	// Makes sure there is something on both sides of the operator
	if operand1 == "" {
		return nil, fmt.Errorf("invalid equality expression! nothing found on left side of operator")
	}
	if operand2 == "" {
		return nil, fmt.Errorf("invalid equality expression! nothing found on right side of operator")
	}

	// Parses both operands
	operand1Val, err := eD.jef.GetParserManager().ParseCode(operand1)
	if err != nil {
		return nil, err
	}
	operand2Val, err := eD.jef.GetParserManager().ParseCode(operand2)
	if err != nil {
		return nil, err
	}

	// Evaluates the equals operator
	boolResult := operand1Val.GetValue() == operand2Val.GetValue()

	resultData := dataValue{
		typeStruct: eD.jef.GetDatatypeManager().GetDatatype(domain.BooleanDataTypeName),
		value:      boolResult,
	}

	return resultData, nil
}
