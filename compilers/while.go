package compilers

import (
	"fmt"
	"github.com/captaincrazybro/jef/domain"
	"github.com/captaincrazybro/jef/util"
	lu "github.com/captaincrazybro/literalutil"
	"regexp"
	"strings"
)

// whileLoop structure for the whileLoop compiler
type whileLoop struct {
	jef domain.Jef
}

// GetName function to return the name of the while compiler
func (w *whileLoop) GetName() string {
	return whileLoopName
}

// Check checks the current line to see if it is an while statement
func (w *whileLoop) Check(s lu.String) bool {
	return s.HasPrefix("while ") || s == "while"
}

// Run runs lines of code included in the while loop
func (w *whileLoop) Run(iter domain.LineIterator) error {
	// Makes sure the whileLoop loop has a conditional statement
	ifR1, _ := regexp.Compile("^while +(\\S.*)$")
	if !ifR1.MatchString(iter.Current().Tos()) {
		return fmt.Errorf("invalid while statement! if statement must have a boolean conditional statement after while identifier")
	}

	// Parses the whileLoop loops
	err, whileCondStr, whileJef := parseLoopStat(ifR1, iter, w.jef, "while")
	if err != nil {
		return err
	}

	// Runs the whileLoop loop and evals the whileCondStr each time
	whileCondData, err := w.jef.GetParserManager().ParseCode(whileCondStr)
	if err != nil {
		return err
	}

	// Parses whileCondData
	whileCond := util.ParseConditionalValue(whileCondData, whileJef)

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
		whileCond = util.ParseConditionalValue(whileCondData, whileJef)
	}

	return nil
}

// parseLoopStat function to hold the steps used to parse an individual condition statement
func parseLoopStat(r1 *regexp.Regexp, iter domain.LineIterator, curJef domain.Jef, loopName string) (error, lu.String, domain.Jef) {
	// Parses first conditional statement
	ifCondStr := lu.String(r1.FindStringSubmatch(iter.Current().Tos())[1])
	ifCondStr = util.TrimWhitespaces(ifCondStr)
	// Checks the first line for the opening '{'
	first, second := util.SplitStartOfStatement(ifCondStr)
	if first != ifCondStr {
		ifCondStr = first
		iter.EditCurrent(second)
		// Handles if the '{' is on the next line
	} else {
		iter.Next()
		if !iter.Current().HasPrefix("{") {
			return fmt.Errorf("invalid %s statement! could not find a closing '{'", loopName), "", nil
		}
		splitCondStr := iter.Current().Split("{")
		iter.EditCurrent(lu.String(strings.Join(splitCondStr[1:splitCondStr.Len()].Tosa(), "{")))
	}
	ifCondStr = util.TrimWhitespaces(ifCondStr)
	ifCondStr = ifCondStr.TrimPrefix("(")
	ifCondStr = ifCondStr.TrimSuffix(")")

	// Parses the lines of the if statement
	err, ifLines := util.ReadInStatement(iter)

	if err != nil {
		return fmt.Errorf("invalid %s statement! %s", loopName, err), "", nil
	}

	// Creates new jef instance and runs it if the condition is true
	jef := curJef.New(ifLines)
	return nil, ifCondStr, jef
}
