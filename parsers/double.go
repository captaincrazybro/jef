package parsers

import (
	"github.com/captaincrazybro/jef/domain"
	lu "github.com/captaincrazybro/literalutil"
	"regexp"
	"strconv"
)

type Double struct {
	jef domain.Jef
}

func (dD Double) GetType() domain.DataType {
	return dD.jef.GetDatatypeManager().GetDatatype("double")
}

func (dD Double) Check(s lu.String) bool {
	r, _ := regexp.Compile("^-?\\d*\\.?\\d*$")
	return r.MatchString(s.Tos())
}

func (dD Double) GetValue(s lu.String) (domain.DataValue, error) {
	r, _ := regexp.Compile("^-?\\d*\\.?\\d*$")
	parsedString := r.FindString(s.Tos())
	value, err := strconv.Atoi(parsedString)
	if err != nil {
		return nil, err
	}

	return dataValue{typeStruct: dD.GetType(), value: value}, nil
}
