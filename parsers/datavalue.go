package parsers

import "github.com/captaincrazybro/jef/domain"

type dataValue struct {
	typeStruct domain.TypeParser
	typeName   string
	value      interface{}
}

// GetType returns the domain.TypeParser instance
func (dV dataValue) GetType() domain.TypeParser {
	return dV.typeStruct
}

// GetTypeName returns the name of the type (this is sometimes different than the domain.TypeParser#GetName() method value
func (dV dataValue) GetTypeName() string {
	return dV.typeName
}

// GetValue returns the value of the
func (dV dataValue) GetValue() interface{} {
	return dV.value
}
