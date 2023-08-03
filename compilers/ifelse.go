package compilers

import (
	"fmt"
	"github.com/captaincrazybro/jef/domain"
	"github.com/captaincrazybro/jef/util"
	lu "github.com/captaincrazybro/literalutil"
	"regexp"
)

// ifElse structure for the ifElse compiler
type ifElse struct {
	jef domain.Jef
}

// GetName function to return the name of the ifElse compiler
func (iE *ifElse) GetName() string {
	return ifElseName
}

// Check checks the current line to see if it is an if statement
func (iE *ifElse) Check(s lu.String) bool {
	return s.HasPrefix("if ") || s == "if"
}

// Run runs lines of code included in the if statement
func (iE *ifElse) Run(iter *util.LineIterator) error {
	firstS := iter.Current()

	// Compiles the first line
	l1R, _ := regexp.Compile("^if(.*){$")
	l1R2, _ := regexp.Compile("^if(.*){}$")
	if !l1R.MatchString(firstS.Tos()) {
		if !l1R2.MatchString(firstS.Tos()) {
			return fmt.Errorf("invalid if statement! if statement must contain an opening '{'")
		}
	}

	// Parses first conditional statement
	ifCondStr := lu.String(l1R.FindString(firstS.Tos()))
	ifCondStr = util.TrimWhitespaces(ifCondStr)
	ifCondStr.TrimPrefix("(")
	ifCondStr.TrimSuffix(")")

	// Parses the contents of the if condition
	val, err := iE.jef.GetParserManager().ParseCode(ifCondStr)
	if err != nil {
		return err
	}

	//
	ifCond := parseConditionalValue(val, iE.jef)

	// Parses the lines of the if statement
	ifLines := []lu.String{}
	if !l1R2.MatchString(firstS.Tos()) {
		_, ifLines = util.ReadInStatement(iter)
	}
	ifJef := iE.jef.New(ifLines)
	if ifCond {
		ifJef.Run()
	}

	// TODO: Implement if else and else statements

	return nil
}

// parseConditionalValue parses any datatype (from the conditional block) and converts it to a boolean
func parseConditionalValue(val domain.DataValue, jef domain.Jef) bool {
	// switch statement for the type of the DataValue
	switch val.GetType() {
	// Case for the int type
	case jef.GetDatatypeManager().GetDatatype("int"):
		{
			return val.GetValue() != 0
		}
	// Case for the remaining types
	default:
		{
			return val.GetValue() != nil
		}
	}
}
