package compilers

import (
	"fmt"
	"github.com/captaincrazybro/jef/domain"
	"github.com/captaincrazybro/jef/util"
	lu "github.com/captaincrazybro/literalutil"
	"regexp"
	"strings"
)

// functionDeclaration structure to compile function declarations
type functionDeclaration struct {
	jef domain.Jef
}

func (fD functionDeclaration) GetName() string {
	return functionDeclarationName
}

func (fD functionDeclaration) Check(s lu.String) bool {
	return s.HasPrefix("fun ") || s == "fun"
}

func (fD functionDeclaration) Run(iter domain.LineIterator) error {
	// Handles the if statement
	r1, _ := regexp.Compile("^fun +([a-zA-Z0-9]* +)?(\\S.*)$")
	if !r1.MatchString(iter.Current().Tos()) {
		return fmt.Errorf("invalid function declaration! must have function name, function parameter list, and function executable")
	}

	err, funcName, funcParamsStr, returnType, funcLines := parseFuncStat(r1, iter, fD.jef)
	if err != nil {
		return err
	}

	// Validates the function name
	if fD.jef.GetFunctionManager().GetFunction(funcName.Tos()) != nil {
		return fmt.Errorf("invalid function declaration! function has already been declared")
	}

	// Validates function parameter list
	strParams := splitFuncParamList(funcParamsStr)
	var params []domain.Parameter
	typeParamRegex, _ := regexp.Compile("^([a-zA-Z0-9]*)* +([a-zA-Z][a-zA-Z0-9_]*)")
	noTypeParamRegex, _ := regexp.Compile("^([a-zA-Z][a-zA-Z0-9_]*)")
	for _, strParam := range strParams {
		if strParam != "" {
			var paramType domain.DataType
			var paramName lu.String
			if typeParamRegex.MatchString(strParam.Tos()) {
				// Parses the param type and param value
				paramTypeStr := typeParamRegex.FindStringSubmatch(strParam.Tos())[1]
				paramName = lu.String(typeParamRegex.FindStringSubmatch(strParam.Tos())[2])

				// Parses data type
				paramType = fD.jef.GetDatatypeManager().GetDatatype(paramTypeStr)
				if paramType == nil {
					return fmt.Errorf("invalid function declaration! parameter type does not exist")
				}
			} else if noTypeParamRegex.MatchString(strParam.Tos()) {
				paramName = lu.String(noTypeParamRegex.FindStringSubmatch(strParam.Tos())[1])
				paramType = fD.jef.GetDatatypeManager().GetDatatype("any")
			} else {
				return fmt.Errorf("invalid function declaration! invalid parameter specified")
			}

			params = append(params, fD.jef.GetFunctionManager().CreateParameter(paramName.Tos(), paramType))
		}
	}

	return fD.jef.GetFunctionManager().RegisterFunction(funcName.Tos(), fD.jef, returnType, params, funcLines)
}

// parseFuncStat function to hold the steps used to parse an individual condition statement for functions
func parseFuncStat(r1 *regexp.Regexp, iter domain.LineIterator, curJef domain.Jef) (error, lu.String, lu.String, domain.DataType, []lu.String) {
	// Parses the return type
	var returnType domain.DataType
	if r1.FindStringSubmatch(iter.Current().Tos())[1] != "" {
		returnTypeStr := lu.String(r1.FindStringSubmatch(iter.Current().Tos())[1])
		returnType = curJef.GetDatatypeManager().GetDatatype(util.TrimWhitespaces(returnTypeStr).Tos())
		if returnType == nil {
			return fmt.Errorf("invalid function declaration! return type specified does not exist"), "", "", nil, nil
		}
	} else {
		returnType = curJef.GetDatatypeManager().GetDatatype(domain.AnyDataTypeName)
	}

	funcSubStr := lu.String(r1.FindStringSubmatch(iter.Current().Tos())[2])
	funcSubStr = util.TrimWhitespaces(funcSubStr)
	// Checks the first line for the opening '{'
	first, second := util.SplitStartOfStatement(funcSubStr)
	if first != funcSubStr {
		funcSubStr = first
		iter.EditCurrent(second)
		// Handles if the '{' is on the next line
	} else {
		iter.Next()
		if !iter.Current().HasPrefix("{") {
			return fmt.Errorf("invalid function declaration! could not find a closing '{'"), "", "", nil, nil
		}
		splitCondStr := iter.Current().Split("{")
		iter.EditCurrent(lu.String(strings.Join(splitCondStr[1:splitCondStr.Len()].Tosa(), "{")))
	}
	funcSubStr = util.TrimWhitespaces(funcSubStr)

	// Parses function name and parameter list
	r, _ := regexp.Compile("([a-zA-Z][a-zA-Z0-9]*)\\((.*)\\)")
	if !r.MatchString(funcSubStr.Tos()) {
		return fmt.Errorf("invalid function declaration! invalid function name or parameter list specified"), "", "", nil, nil
	}

	subs := r.FindStringSubmatch(funcSubStr.Tos())
	funcName := lu.String(subs[1])
	funcName = util.TrimWhitespaces(funcName)
	funcParams := lu.String(subs[2])
	funcParams = util.TrimWhitespaces(funcParams)
	funcParams = funcParams.TrimPrefix("(").TrimSuffix(")")

	// Parses the lines of the if statement
	err, funcLines := util.ReadInStatement(iter)

	if err != nil {
		return fmt.Errorf("invalid declaration statement! %s", err), "", "", nil, nil
	}

	return nil, funcName, funcParams, returnType, funcLines
}

// splitFuncParamList splits the function parameter list into sub strings
func splitFuncParamList(s lu.String) []lu.String {
	// Loops through the string, splitting when encountering a "," not in a string
	isQuote := false
	startIndex := 0
	var params []lu.String
	for i, r := range s {
		if r == '"' && (i == 0 || s[i-1] != '\\') {
			isQuote = !isQuote
		} else if !isQuote {
			if r == ',' {
				params = append(params, s[startIndex:i])
				startIndex = i + 1
			}
		}
	}

	// Appends the last param slice to the list of params
	params = append(params, s[startIndex:])

	return params
}
