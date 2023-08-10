package typeparsers

import (
	"github.com/captaincrazybro/jef/domain"
	lu "github.com/captaincrazybro/literalutil"
)

type Boolean struct {
	jef domain.Jef
}

func (bD Boolean) GetType() domain.DataType {
	return bD.jef.GetDatatypeManager().GetDatatype(domain.BooleanDataTypeName)
}

func (bD Boolean) Check(s lu.String) bool {
	return s.Tos() == "true" || s.Tos() == "false"
}

func (bD Boolean) GetValue(s lu.String) (domain.DataValue, error) {
	boolVal := false
	if s.Tos() == "true" {
		boolVal = true
	}

	return dataValue{value: boolVal, typeStruct: bD.GetType()}, nil
}
