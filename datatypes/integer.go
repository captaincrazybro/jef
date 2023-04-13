package datatypes

import (
	lu "github.com/captaincrazybro/literalutil"
	"regexp"
	"strconv"
)

type Integer struct {
}

func (iD Integer) GetName() string {
	return integerDatatypeName
}

func (iD Integer) GetVarName() string {
	return integerDatatypeName
}

func (iD Integer) Check(s lu.String) bool {
	r, _ := regexp.Compile("^-?\\d*$")
	return r.MatchString(s.Tos())
}

func (iD Integer) GetValue(s lu.String) (interface{}, error) {
	r, _ := regexp.Compile("^-?\\d*$")
	parsedString := r.FindString(s.Tos())
	return strconv.Atoi(parsedString)
}
