package functions

import "github.com/captaincrazybro/jef/domain"

type function struct {
	name string
	funcType string
	exec func(domain.Jef)
	params map[string]domain.Datatype
}

func (f function) GetName() string {
	return f.name
}

func (f function) GetType() string {
	return f.funcType
}

func (f function) GetExec() func(domain.Jef) {
	return f.exec
}

func (f function) GetParams() map[string]domain.Datatype {
	return f.params
}