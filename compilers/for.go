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
	err, forStr, whileJef := parseLoopStat(forR1, iter, fL.jef)
	if err != nil {
		return err
	}

	// Runs the whileLoop loop and evals the whileCondStr each time
	whileCondData, err := fL.jef.GetParserManager().ParseCode(whileCondStr)
	if err != nil {
		return err
	}

	// Parses whileCondData
	whileCond := util.ParseConditionalValue(whileCondData, fL.jef)

	for whileCond {
		err = whileJef.Run()
		if err != nil {
			return err
		}

		// Reevaluates the whileLoop conditional
		whileCondData, err = w.jef.GetParserManager().ParseCode(whileCondStr)
		if err != nil {
			return err
		}

		// Parses whileCondData
		whileCond = util.ParseConditionalValue(whileCondData, w.jef)
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

	// Checks to see if there are the right amount of expressions in the statement
	if len(strings) != 3 {
		return fmt.Errorf("invalid for statement! invalid number of for loop expressions in for statement: looking for 3"), nil
	}

}
