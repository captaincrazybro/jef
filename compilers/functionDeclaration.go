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
	r1, _ := regexp.Compile("^fun +(\\S.*)$")
	if !r1.MatchString(iter.Current().Tos()) {
		return fmt.Errorf("invalid function declaration! must have function name, function parameter list, and function executable")
	}

	err, funcName, funcParamsStr, funcJef := parseFuncStat(r1, iter, fD.jef)
	if err != nil {
		return err
	}

	// Validates the function name
	if fD.jef.GetFunctionManager().GetFunction(funcName.Tos()) != nil {
		return fmt.Errorf("invalid function declaration! function has already been declared")
	}

	// Validates function parameter list
	strParams := splitFuncParamList(funcParamsStr)

}

// parseFuncStat function to hold the steps used to parse an individual condition statement for functions
func parseFuncStat(r1 *regexp.Regexp, iter domain.LineIterator, curJef domain.Jef) (error, lu.String, lu.String, domain.Jef) {
	// Parses first conditional statement
	funcSubStr := lu.String(r1.FindStringSubmatch(iter.Current().Tos())[1])
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
			return fmt.Errorf("invalid function declaration! could not find a closing '{'"), "", "", nil
		}
		splitCondStr := iter.Current().Split("{")
		iter.EditCurrent(lu.String(strings.Join(splitCondStr[1:splitCondStr.Len()].Tosa(), "{")))
	}
	funcSubStr = util.TrimWhitespaces(funcSubStr)

	// Parses function name and parameter list
	r, _ := regexp.Compile("([a-zA-Z][a-zA-Z0-9]*)\\((.*)\\)")
	if !r.MatchString(funcSubStr.Tos()) {
		return fmt.Errorf("invalid function declaration! invalid function name or parameter list specified"), "", "", nil
	}

	subs := r.FindStringSubmatch(funcSubStr.Tos())
	funcName := lu.String(subs[1])
	funcName = util.TrimWhitespaces(funcName)
	funcParams := lu.String(subs[2])
	funcParams = util.TrimWhitespaces(funcParams)
	funcParams = funcParams.TrimPrefix("(").TrimSuffix(")")

	// Parses the lines of the if statement
	err, ifLines := util.ReadInStatement(iter)

	if err != nil {
		return fmt.Errorf("invalid declaration statement! %s", err), "", "", nil
	}

	// Creates new jef instance and runs it if the condition is true
	jef := curJef.New(ifLines)
	return nil, funcName, funcParams, jef
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
