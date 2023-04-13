package util

import (
	lu "github.com/captaincrazybro/literalutil"
	"regexp"
)

func TrimWhitespaces(s lu.String) lu.String {
	paramR, _ := regexp.Compile("^\\s+|\\s+$")
	return lu.String(paramR.ReplaceAllString(s.Tos(), ""))
}
