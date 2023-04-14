package functions

import "github.com/captaincrazybro/jef/domain"

type parameter struct {
	name     string
	dataType domain.TypeParser
}

func (p parameter) GetName() string {
	return p.name
}

func (p parameter) GetType() domain.TypeParser {
	return p.dataType
}
