package typeparsers

import (
	"fmt"
	"github.com/captaincrazybro/jef/domain"
	"github.com/captaincrazybro/jef/util"
	lu "github.com/captaincrazybro/literalutil"
)

type parserManager struct {
	jef     domain.Jef
	parsers []domain.TypeParser
}

// New creates a new datatype manager instance
func New(j domain.Jef) domain.ParserManager {
	dm := &parserManager{parsers: []domain.TypeParser{}, jef: j}
	dm.registerParsers()
	return dm
}

// registerParsers registers all the parsers
func (pm *parserManager) registerParsers() {
	pm.AddParser(Integer{jef: pm.jef})
	pm.AddParser(Double{jef: pm.jef})
	pm.AddParser(String{jef: pm.jef})
	pm.AddParser(Boolean{jef: pm.jef})
	pm.AddParser(Variable{jef: pm.jef})
}

// AddParser adds a parser to the datatype manager
func (pm *parserManager) AddParser(d domain.TypeParser) {
	pm.parsers = append(pm.parsers, d)
}

// ParseCode finds the appropriate parser for the given code string
func (pm *parserManager) ParseCode(s lu.String) (domain.DataValue, error) {
	// Removes whitespaces around datatype
	s = util.TrimWhitespaces(s)

	for _, v := range pm.parsers {
		if v.Check(s) {
			dV, err := v.GetValue(s)
			if err != nil {
				return nil, err
			}

			return dV, nil
		}
	}

	return nil, fmt.Errorf("code not recognized! type does not exist")
}
