package functions

import "github.com/captaincrazybro/jef/domain"

type function struct {
	name       string
	returnType domain.TypeParser
	exec       func(domain.Jef)
	params     []domain.Parameter
}

func (f function) GetName() string {
	return f.name
}

func (f function) GetReturnType() domain.TypeParser {
	return f.returnType
}

func (f function) GetExec() func(domain.Jef) {
	return f.exec
}

func (f function) GetParams() []domain.Parameter {
	return f.params
}

func (f function) RunExec(values []domain.DataValue, j domain.Jef) error {
	// Validates the function values
	err := validateParameters(f, values, j)
	if err != nil {
		return err
	}

	// Creates a new jef instance with the same variables, functions etc...
	newJ := j.NewCodeless()
	// Adds the give parameters as variables
	for i, val := range values {
		dType := val.GetType()
		param := f.GetParams()[i]
		err := newJ.GetVariableManager().RegisterVariable(param.GetName(), dType, val.GetValue())
		if err != nil {
			return err
		}
	}

	f.exec(newJ)
	return nil
}
