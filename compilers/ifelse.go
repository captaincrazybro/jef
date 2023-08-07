package compilers

import (
	"fmt"
	"github.com/captaincrazybro/jef/domain"
	"github.com/captaincrazybro/jef/util"
	lu "github.com/captaincrazybro/literalutil"
	"regexp"
	"strings"
)

type elseIfStruct struct {
	ifCond bool
	jef    domain.Jef
}

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
func (iE *ifElse) Run(iter domain.LineIterator) error {
	// Handles the if statement
	ifR1, _ := regexp.Compile("^if +(\\S.*)$")
	if !ifR1.MatchString(iter.Current().Tos()) {
		return fmt.Errorf("invalid if statement! if statement must have a boolean conditional statement after if identifier")
	}

	err, ifCond, ifJef := parseCondStat(ifR1, iter, iE.jef, 1, false)
	if err != nil {
		return err
	}

	// Handles if else statements
	var elseIfz []elseIfStruct
	elifR1, _ := regexp.Compile("^} *(else if|elif).*$")
	elifR2, _ := regexp.Compile("^} *(else if|elif) +(\\S.*)$")
	for elifR1.MatchString(iter.Current().Tos()) {
		// Makes sure there is a potential conditional after the else if
		if !elifR2.MatchString(iter.Current().Tos()) {
			return fmt.Errorf("invalid if statement! if else statement must have a boolean conditional statement after if else identifier")
		}

		err, elifCond, elifJef := parseCondStat(elifR2, iter, iE.jef, 2, false)
		if err != nil {
			return err
		}

		elseIfz = append(elseIfz, elseIfStruct{ifCond: elifCond, jef: elifJef})
	}

	// Handle last else statement
	var elseJef domain.Jef
	elseR1, _ := regexp.Compile("^} *else[ |{]+.*")
	elseR2, _ := regexp.Compile("^} *else$")
	elseR3, _ := regexp.Compile("^} *else *{? *$")
	hasElse := false
	if elseR1.MatchString(iter.Current().Tos()) || elseR2.MatchString(iter.Current().Tos()) {
		// Makes sure there is no conditional (or something else) after the else identifier
		if !elseR3.MatchString(iter.Current().Tos()) {
			return fmt.Errorf("invalid if statement! else statement must not have a boolean conditional statement after else identifier")
		}

		hasElse = true
		err, _, elseJef = parseCondStat(elseR1, iter, iE.jef, 1, true)
		if err != nil {
			return err
		}
	}

	// Checks to make sure there is no invalid expression at the end of the final '}'
	if iter.Current().HasPrefix("}") && iter.Current().ReplaceAll("}", "") != "" {
		return fmt.Errorf("invalid if statement! invalid statement found at the end of closing '}'")
	}

	// Sort through if/else statements
	elseIfTrueIndex := elseIfIsTrue(elseIfz)
	if ifCond {
		err = ifJef.Run()
		if err != nil {
			return err
		}
	} else if elseIfTrueIndex != -1 {
		err = elseIfz[elseIfTrueIndex].jef.Run()
		if err != nil {
			return err
		}
	} else if hasElse {
		err = elseJef.Run()
		if err != nil {
			return err
		}
	}

	return nil
}

// ifElseIsTrue checks to see if one of the else ifs is true
func elseIfIsTrue(elseIfz []elseIfStruct) int {
	for i, elseIf := range elseIfz {
		if elseIf.ifCond {
			return i
		}
	}

	return -1
}

// parseCondStat function to hold the steps used to parse an individual condition statement
func parseCondStat(r1 *regexp.Regexp, iter domain.LineIterator, curJef domain.Jef, capIndex int, isElse bool) (error, bool, domain.Jef) {
	// If the statement is not an else statement, then parses the conditional statement
	ifCond := false
	if !isElse {
		// Parses first conditional statement
		ifCondStr := lu.String(r1.FindStringSubmatch(iter.Current().Tos())[capIndex])
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
				return fmt.Errorf("invalid if statement! could not find a closing '{'"), false, nil
			}
			splitCondStr := iter.Current().Split("{")
			iter.EditCurrent(lu.String(strings.Join(splitCondStr[1:splitCondStr.Len()].Tosa(), "{")))
		}
		ifCondStr = util.TrimWhitespaces(ifCondStr)
		ifCondStr = ifCondStr.TrimPrefix("(")
		ifCondStr = ifCondStr.TrimSuffix(")")

		// Parses the contents of the if condition
		val, err := curJef.GetParserManager().ParseCode(ifCondStr)
		if err != nil {
			return err, false, nil
		}

		ifCond = util.ParseConditionalValue(val, curJef)
		// Handles else statement
	} else {
		_, second := util.SplitStartOfStatement(iter.Current())
		iter.EditCurrent(second)
	}

	// Parses the lines of the if statement
	var ifLines []lu.String
	var err error
	err, ifLines = util.ReadInStatement(iter)

	if err != nil {
		return fmt.Errorf("invalid if statement! %s", err), false, nil
	}

	// Creates new jef instance and runs it if the condition is true
	jef := curJef.New(ifLines)
	return nil, ifCond, jef
}
