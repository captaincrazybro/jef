package util

import (
	lu "github.com/captaincrazybro/literalutil"
	"regexp"
)

func TrimWhitespaces(s lu.String) lu.String {
	paramR, _ := regexp.Compile("^\\s+|\\s+$")
	return lu.String(paramR.ReplaceAllString(s.Tos(), ""))
}

func ReadInStatement(iter *LineIterator) (error, []lu.String) {
	// Loops through the lines till the statement has ended (looking for a closing })
	startIndex := iter.Index()
	numOpening := 1
	for iter.Next() && numOpening != 0 {
		if iter.Current().Trim(" ").HasSuffix("{") {
			numOpening++
		} else if iter.Current().Trim(" ").HasSuffix("}") {
			numOpening--
		}
	}

	return nil, iter.Lines()[startIndex:iter.Index()]
}
