package datatypes

import (
	"github.com/captaincrazybro/jef/domain"
	"github.com/captaincrazybro/jef/util"
	lu "github.com/captaincrazybro/literalutil"
)

// datatypeManager stores and manages datatypes
type datatypeManager struct {
	datatypes []domain.Datatype
	jef       domain.Jef
}

// New creates a new datatype manager instance
func New(j domain.Jef) domain.DatatypeManager {
	dm := &datatypeManager{datatypes: []domain.Datatype{}, jef: j}
	dm.registerDatatypes()
	return dm
}

// registerDatatypes registers datatypes
func (dm *datatypeManager) registerDatatypes() {
	dm.AddDatatype(Variable{jef: dm.jef})
	dm.AddDatatype(Integer{})
	dm.AddDatatype(Double{})
	dm.AddDatatype(String{})
}

// AddDatatype adds a datatype to the datatype manager
func (dm *datatypeManager) AddDatatype(d domain.Datatype) {
	dm.datatypes = append(dm.datatypes, d)
}

// FindDatatype finds a datatype with a raw code string
func (dm *datatypeManager) FindDatatype(s lu.String) (domain.DataValue, error) {
	// Removes whitespaces around datatype
	s = util.TrimWhitespaces(s)

	for _, v := range dm.datatypes {
		if v.Check(s) {
			_, err := v.GetValue(s)
			if err != nil {
				return nil, err
			}

			return dataValue{typeStruct: v, typeName: v.GetName(), value: v.GetValue(s)}, nil
		}
	}
	return nil, nil
}

// GetDatatype gets a datatype based on the name of the datatype
func (dm *datatypeManager) GetDatatype(name string) domain.Datatype {
	// Loops through the datatypes
	for _, dt := range dm.datatypes {
		if dt.GetName() == name {
			return dt
		}
	}

	return nil
}
