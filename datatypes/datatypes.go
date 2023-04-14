package datatypes

import (
	"github.com/captaincrazybro/jef/domain"
)

// datatypeManager stores and manages datatypes
type datatypeManager struct {
	dataTypes []domain.DataType
	jef       domain.Jef
}

// New creates a new datatype manager instance
func New(j domain.Jef) domain.DatatypeManager {
	dm := &datatypeManager{dataTypes: []domain.DataType{}, jef: j}
	dm.registerDataTypes()
	return dm
}

// registerDataTypes registers all the dataTypes
func (dm *datatypeManager) registerDataTypes() {

}

// AddDataType registers a datatype to the list of dataTypes in side the datatype manager
func (dm *datatypeManager) AddDataType(dataType domain.DataType) {
	dm.dataTypes = append(dm.dataTypes, dataType)
}

// GetDatatype gets a datatype based on the name of the datatype
func (dm *datatypeManager) GetDatatype(name string) domain.DataType {
	// Loops through the datatypes
	for _, dt := range dm.dataTypes {
		if dt.GetName() == name {
			return dt
		}
	}

	return nil
}
