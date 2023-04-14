package parsers

import (
	"github.com/captaincrazybro/jef/domain"
	lu "github.com/captaincrazybro/literalutil"
	"regexp"
	"strconv"
)

type Integer struct {
	jef domain.Jef
}

func (iD Integer) GetType() domain.DataType {
	return iD.jef.GetDatatypeManager().GetDatatype("int")
}

func (iD Integer) Check(s lu.String) bool {
	r, _ := regexp.Compile("^-?\\d*$")
	return r.MatchString(s.Tos())
}

func (iD Integer) GetValue(s lu.String) (domain.DataValue, error) {
	r, _ := regexp.Compile("^-?\\d*$")
	parsedString := r.FindString(s.Tos())
	value, err := strconv.ParseFloat(parsedString, 32)
	if err != nil {
		return nil, err
	}

	return dataValue{value: value, typeStruct: iD.GetType()}, nil
}
