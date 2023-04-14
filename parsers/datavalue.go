package parsers

import "github.com/captaincrazybro/jef/domain"

type dataValue struct {
	typeStruct domain.DataType
	value      interface{}
}

// GetType returns the domain.TypeParser instance
func (dV dataValue) GetType() domain.DataType {
	return dV.typeStruct
}

// GetValue returns the value of the
func (dV dataValue) GetValue() interface{} {
	return dV.value
}
