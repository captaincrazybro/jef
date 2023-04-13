package datatypes

import "github.com/captaincrazybro/jef/domain"

type dataValue struct {
	typeStruct domain.Datatype
	typeName   string
	value      interface{}
}

// GetType returns the domain.Datatype instance
func (dV dataValue) GetType() domain.Datatype {
	return dV.typeStruct
}

// GetTypeName returns the name of the type (this is sometimes different than the domain.Datatype#GetName() method value
func (dV dataValue) GetTypeName() string {
	return dV.typeName
}

// GetValue returns the value of the
func (dV dataValue) GetValue() interface{} {
	return dV.value
}
