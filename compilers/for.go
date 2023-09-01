package compilers

import (
	"fmt"
	"github.com/captaincrazybro/jef/domain"
	"github.com/captaincrazybro/jef/util"
	lu "github.com/captaincrazybro/literalutil"
	"regexp"
)

// for structure for the for compiler
type forLoop struct {
	jef domain.Jef
}

// GetName function to return the name of the forLoop compiler
func (fL *forLoop) GetName() string {
	return forLoopName
}

// Check checks the current line to see if it is a for statement
func (fL *forLoop) Check(s lu.String) bool {
	return s.HasPrefix("for ") || s == "for"
}

// Run runs lines of code included in the for loop
func (fL *forLoop) Run(iter domain.LineIterator) error {
	// Makes sure the whileLoop loop has a conditional statement
	forR1, _ := regexp.Compile("^for +(\\S.*)$")
	if !forR1.MatchString(iter.Current().Tos()) {
		return fmt.Errorf("invalid for statement! if statement must have a boolean conditional statement after for identifier")
	}

	// Parses the whileLoop loops
	err, forStr, forJef := parseLoopStat(forR1, iter, fL.jef, "for")
	if err != nil {
		return err
	}

	// Runs the whileLoop loop and evals the whileCondStr each time
	err, forExprz := splitForStatement(forStr)
	if err != nil {
		return err
	}

	// Runs the initial expression
	initExprLine := util.LineIterator{}
	initExprLine.New([]lu.String{forExprz[0]})
	initExprLine.Next()
	err = forJef.GetCompilerManager().CompileLine(&initExprLine)
	if err != nil {
		return err
	}

	// Parses the loop conditional statement
	condData, err := forJef.GetParserManager().ParseCode(forExprz[1])
	if err != nil {
		return err
	}

	cond := util.ParseConditionalValue(condData, forJef)

	// Prepares the in-loop expression
	inLoopExprLine := util.LineIterator{}
	inLoopExprLine.New([]lu.String{forExprz[2]})
	inLoopExprLine.Next()

	for cond {
		err = forJef.Run()
		if err != nil {
			return err
		}

		// Runs the in-loop expression
		err = forJef.GetCompilerManager().CompileLine(&inLoopExprLine)
		if err != nil {
			return err
		}

		// Reevaluates the for loop conditional
		condData, err = forJef.GetParserManager().ParseCode(forExprz[1])
		if err != nil {
			return err
		}

		cond = util.ParseConditionalValue(condData, forJef)
	}

	return nil
}

// splitForStatement splits the expression for the for statement into the three statements
func splitForStatement(s lu.String) (error, []lu.String) {
	// Loops through the string finding the semicolons
	isQuote := false
	startIndex := 0
	var strings []lu.String
	for i, r := range s {
		if r == '"' && (i == 0 || s[i-1] == '\\') {
			isQuote = !isQuote
		} else if !isQuote {
			if r == ';' {
				strings = append(strings, s[startIndex:i])
				startIndex = i + 1
			}
		}
	}

	// Appends the last expression slice
	strings = append(strings, s[startIndex:])

	// Checks to see if there are the right amount of expressions in the statement
	if len(strings) != 3 {
		return fmt.Errorf("invalid for statement! invalid number of for loop expressions in for statement: looking for 3"), nil
	}

	return nil, strings
}
