package typeparsers

import (
	"fmt"
	"github.com/captaincrazybro/jef/domain"
	"github.com/captaincrazybro/jef/util"
	lu "github.com/captaincrazybro/literalutil"
	"math"
	"regexp"
)

var mathOperators = []rune{'^', '*', '/', '+', '-'}
var mathCheckRegex, _ = regexp.Compile("^[^\"].*[\\^+\\-/*].*")

type mathChunk struct {
	isOperator bool
	isDecimal  bool
	strValue   lu.String
	value      interface{}
}

type Math struct {
	jef domain.Jef
}

func (mD Math) GetName() string {
	return domain.MathParserName
}

func (mD Math) GetType() domain.DataType {
	return nil
}

func (mD Math) Check(s lu.String) bool {
	return mathCheckRegex.MatchString(s.Tos())
}

func (mD Math) GetValue(s lu.String) (domain.DataValue, error) {
	err, chunk := evalMath(s, mD.jef)
	if err != nil {
		return nil, err
	}

	if chunk.isDecimal {
		return dataValue{
			typeStruct: mD.jef.GetDatatypeManager().GetDatatype("double"),
			value:      chunk.value.(float64),
		}, nil
	} else {
		return dataValue{
			typeStruct: mD.jef.GetDatatypeManager().GetDatatype("int"),
			value:      chunk.value.(int),
		}, nil
	}
}

// evalMath evaluates the current math statement
func evalMath(s lu.String, jef domain.Jef) (error, mathChunk) {
	// Main part of the function which accounts for the regular math expressions
	if mathCheckRegex.MatchString(s.Tos()) {
		// Divides the expression into different
		err, chunks := splitIntoMathChunks(s)
		if err != nil {
			return err, mathChunk{}
		}

		// Checks to see if there were any enclosing parenthesis around the expression
		for len(chunks) == 1 && s.HasPrefix("(") && s.HasSuffix(")") {
			s = s.TrimPrefix("(").TrimSuffix(")")
			err, chunks = splitIntoMathChunks(s)
			if err != nil {
				return err, mathChunk{}
			}
		}

		// Loops through the chunks and evaluates
		for _, op := range mathOperators {
			for i := 0; i < len(chunks); i++ {
				chunk := chunks[i]
				// Checks if the current chunk is an operator and that it is the current operator being evaluated
				if chunk.isOperator && util.TrimWhitespaces(chunk.strValue) == lu.String(op) {
					firstChunk := chunks[i-1]
					secondChunk := chunks[i+1]

					// Evaluates both sides of the operator
					if firstChunk.value == nil {
						if isSubMathStatement(firstChunk.strValue) {
							err, chunks[i-1] = evalMath(util.TrimWhitespaces(firstChunk.strValue).TrimPrefix("(").TrimSuffix(")"), jef)
							if err != nil {
								return err, mathChunk{}
							}
						} else {
							err, chunks[i-1] = parseMathChunkValue(firstChunk, jef)
							if err != nil {
								return err, mathChunk{}
							}
						}
					}
					if secondChunk.value == nil {
						if isSubMathStatement(secondChunk.strValue) {
							err, chunks[i+1] = evalMath(util.TrimWhitespaces(secondChunk.strValue).TrimPrefix("(").TrimSuffix(")"), jef)
							if err != nil {
								return err, mathChunk{}
							}
						} else {
							err, chunks[i+1] = parseMathChunkValue(secondChunk, jef)
							if err != nil {
								return err, mathChunk{}
							}
						}
					}

					// Evaluates the results of the operation
					resultChunk := evalMathOperator(chunks[i-1], chunk, chunks[i+1])
					subChunks := append(chunks[:i-1], resultChunk)
					chunks = append(subChunks, chunks[i+2:]...)
				}
			}
		}

		return nil, chunks[0]
		// Accounts for the odd chance that there is just a number or something else in the statement
	} else {
		err, lastChunk := parseMathChunkValue(mathChunk{
			isOperator: false,
			isDecimal:  false,
			strValue:   s,
			value:      nil,
		}, jef)
		if err != nil {
			return err, mathChunk{}
		}

		return nil, lastChunk
	}
}

// splitIntoMathChunks divides the math expression into different chunks of operands and mathOperators
func splitIntoMathChunks(line lu.String) (error, []mathChunk) {
	isQuote := false
	parenthesisCount := 0
	startIndex := 0
	var chunks []mathChunk

	for i, s := range line.Tos() {
		if s == '"' && (i == 0 || line[s-1] != '\\') {
			isQuote = !isQuote
		}

		if !isQuote {
			if s == '(' {
				parenthesisCount++
			} else if s == ')' {
				parenthesisCount--
			} else if s == '"' && (i == 0 || line[s-1] != '\\') {
				isQuote = !isQuote
			} else if parenthesisCount == 0 && util.RuneArrContains(mathOperators, s) {
				// Makes sure the first slice is not empty
				firstSlice := util.TrimWhitespaces(line[startIndex:i])
				if firstSlice == "" {
					return fmt.Errorf("invalid math operation! nothing found on one side of operator"), nil
				}
				firstChunk := mathChunk{isOperator: false, isDecimal: false, strValue: firstSlice, value: nil}
				chunks = append(chunks, firstChunk)

				// Gets the operator
				oprChunk := mathChunk{isOperator: true, isDecimal: false, strValue: lu.String(s), value: nil}
				chunks = append(chunks, oprChunk)

				// Sets the start index to the very next index
				startIndex = i + 1
			}
		}
	}

	if parenthesisCount != 0 {
		return fmt.Errorf("invalid math operation! closing parenthesis was not found"), nil
	}

	// Adds the last chunk
	lastChunk := mathChunk{
		isOperator: false,
		isDecimal:  false,
		strValue:   line[startIndex:],
		value:      nil,
	}
	chunks = append(chunks, lastChunk)

	return nil, chunks
}

// evalMathOperator evaluates an operation between two math chunks
func evalMathOperator(firstChunk mathChunk, opChunk mathChunk, lastChunk mathChunk) mathChunk {
	// Casts the interface into a float64
	var val1 float64
	var val2 float64
	if firstChunk.isDecimal {
		val1 = firstChunk.value.(float64)
	} else {
		val1 = float64(firstChunk.value.(int))
	}
	if lastChunk.isDecimal {
		val2 = lastChunk.value.(float64)
	} else {
		val2 = float64(lastChunk.value.(int))
	}
	var result float64

	switch util.TrimWhitespaces(opChunk.strValue) {
	case "^":
		{
			result = math.Pow(val1, val2)
			break
		}
	case "*":
		{
			result = val1 * val2
			break
		}
	case "/":
		{
			result = val1 / val2
			break
		}
	case "+":
		{
			result = val1 + val2
			break
		}
	case "-":
		{
			result = val1 - val2
			break
		}
	}

	if firstChunk.isDecimal || lastChunk.isDecimal {
		return mathChunk{
			isOperator: false,
			isDecimal:  true,
			strValue:   "",
			value:      result,
		}
	} else {
		return mathChunk{
			isOperator: false,
			isDecimal:  false,
			strValue:   "",
			value:      int(result),
		}
	}
}

// parseMathChunkValue parses the value of the current math chunk
func parseMathChunkValue(chunk mathChunk, jef domain.Jef) (error, mathChunk) {
	data, err := jef.GetParserManager().ParseCode(chunk.strValue)
	if err != nil {
		return err, mathChunk{}
	}

	if data.GetType().GetName() != domain.IntDataTypeName && data.GetType().GetName() != domain.DoubleDataTypeName {
		return fmt.Errorf("invalid math operation! math expression does not evaluate to int or double"), mathChunk{}
	}

	chunk.value = data.GetValue()
	chunk.isDecimal = data.GetType().GetName() == "double"

	return nil, chunk
}

// isSubMathStatement checks to see if the current value has parenthesis and is therefore a sub math statement
func isSubMathStatement(s lu.String) bool {
	return util.TrimWhitespaces(s).HasPrefix("(")
}
