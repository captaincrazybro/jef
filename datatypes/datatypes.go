package datatypes

import "github.com/captaincrazybro/jef/domain"

// datatypeManager stores and manages datatypes
type datatypeManager struct {
	datatypes []domain.Datatype
}

// New creates a new datatype manager instance
func New(j domain.Jef) domain.DatatypeManager {
	dm := datatypeManager{datatypes: []domain.Datatype{}}
	dm.registerDatatypes(j)
	return &dm
}

// registerDatatypes registers datatypes
func (dm *datatypeManager) registerDatatypes(j domain.Jef) {

}

// AddDatatype adds a datatype to the datatype manager
func (dm *datatypeManager) AddDatatype(d domain.Datatype) {
	dm.datatypes = append(dm.datatypes, d)
}

// FindDatatype finds a datatype with the a raw code string
func (dm *datatypeManager) FindDatatype(s string) (domain.Datatype, error) {
	for _, v := range dm.datatypes {
		if v.Check(s) {
			_, err := v.GetValue(s)
			if err != nil {
				return nil, err
			}
			return v, nil
		}
	}
	return nil, nil
}