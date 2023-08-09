package util

import (
	"fmt"
	"github.com/captaincrazybro/jef/domain"
	lu "github.com/captaincrazybro/literalutil"
	"regexp"
)

func TrimWhitespaces(s lu.String) lu.String {
	paramR, _ := regexp.Compile("^\\s+|\\s+$")
	return lu.String(paramR.ReplaceAllString(s.Tos(), ""))
}

// ReadInStatement reads in a statement for an if, for... etc... statement
func ReadInStatement(iter domain.LineIterator) (error, []lu.String) {
	// Loops through the lines till the statement has ended (looking for a closing })
	if iter.Current() == "" {
		iter.Next()
	}
	startIndex := iter.Index()
	var endIndex int
	numOpening := 1
	for numOpening != 0 {
		if iter.Current().HasPrefix("}") || iter.Current().HasSuffix("}") {
			numOpening--
		} else if iter.Current().HasSuffix("{") {
			numOpening++
		}

		if numOpening != 0 {
			if !iter.Next() {
				return fmt.Errorf("no closing '}' found"), nil
			}
		} else {
			// Removes trailing '}' if necessary
			endIndex = iter.Index()
			if iter.Current() != "}" && !iter.Current().HasPrefix("}") && iter.Current().HasSuffix("}") {
				iter.EditCurrent(iter.Current().TrimSuffix("}"))
				endIndex++
			}
		}
	}

	statLines := iter.Lines()[startIndex:endIndex]
	return nil, statLines
}

// SplitStartOfStatement splits the first part of the statement which potentially contains the conditional and first identifier from the trail of the statement
func SplitStartOfStatement(line lu.String) (lu.String, lu.String) {
	isQuote := false
	for i, ch := range line {
		if ch == '"' && line[i-1] != '\\' {
			isQuote = !isQuote
			// If the current character is an '{' and
		} else if ch == '{' && !isQuote {
			return lu.String(line.Tos()[0:i]), lu.String(line.Tos()[i+1 : len(line)])
		}
	}

	return line, ""
}
