package datatypes

import (
	lu "github.com/captaincrazybro/literalutil"
	"regexp"
	"strconv"
)

type Double struct {
}

func (dD Double) GetName() string {
	return doubleDatatypeName
}

func (dD Double) GetVarName() string {
	return doubleDatatypeName
}

func (dD Double) Check(s lu.String) bool {
	r, _ := regexp.Compile("^-?\\d*\\.?\\d*$")
	return r.MatchString(s.Tos())
}

func (dD Double) GetValue(s lu.String) (interface{}, error) {
	r, _ := regexp.Compile("^-?\\d*\\.?\\d*$")
	parsedString := r.FindString(s.Tos())
	return strconv.Atoi(parsedString)
}
