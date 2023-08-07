package functions

import (
	"github.com/captaincrazybro/jef/domain"
)

type function struct {
	jef        domain.Jef
	name       string
	returnType domain.TypeParser
	exec       func(jef domain.Jef)
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

func (f function) RunExec(values []domain.DataValue) error {
	// Validates the function values
	err := validateParameters(f, values, f.jef)
	if err != nil {
		return err
	}

	// Creates a new jef instance with the same variables, functions etc...
	newJ := f.jef.NewCodeless()
	// Adds the give parameters as variables
	for i, val := range values {
		dType := val.GetType()
		param := f.GetParams()[i]

		// Checks to see if the variable already exists. If it does exist, then it deletes it
		exists := newJ.GetVariableManager().GetVariable(param.GetName()) != nil
		if exists {
			newJ.GetVariableManager().DeleteVariable(param.GetName())
		}

		newJ.GetVariableManager().RegisterVariable(param.GetName(), dType, val.GetValue())
	}

	f.exec(newJ)
	return nil
}
