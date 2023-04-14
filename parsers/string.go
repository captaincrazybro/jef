package parsers

import (
	"fmt"
	"github.com/captaincrazybro/jef/domain"
	lu "github.com/captaincrazybro/literalutil"
	"regexp"
	"strings"
)

type String struct {
	jef domain.Jef
}

func (sD String) GetType() domain.DataType {
	return sD.jef.GetDatatypeManager().GetDatatype("string")
}

func (sD String) Check(s lu.String) bool {
	r, _ := regexp.Compile("^\".*\"$")
	return r.MatchString(s.Tos())
}

func (sD String) GetValue(s lu.String) (interface{}, error) {
	r, _ := regexp.Compile("^\"(.*)\"$")
	str := r.FindStringSubmatch(s.Tos())[1]

	// Accounts for backslash plus a quote
	for i, s := range str {
		if s == '"' {
			if i == 0 || str[i-1] == '\\' {
				return nil, fmt.Errorf("invalid string! additional quote detected in string literal")
			}
		}
	}
	strings.ReplaceAll(str, "\\\"", "\"")

	return str, nil
}
