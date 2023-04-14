package functions

import "github.com/captaincrazybro/jef/domain"

type parameter struct {
	name     string
	dataType domain.DataType
}

func (p parameter) GetName() string {
	return p.name
}

func (p parameter) GetType() domain.DataType {
	return p.dataType
}
