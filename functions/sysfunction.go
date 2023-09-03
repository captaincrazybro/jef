package functions

import (
	"github.com/captaincrazybro/jef/domain"
)

type sysFunction struct {
	jef        domain.Jef
	name       string
	returnType domain.DataType
	exec       func(jef domain.Jef) error
	params     []domain.Parameter
}

func (f sysFunction) GetName() string {
	return f.name
}

func (f sysFunction) GetReturnType() domain.DataType {
	return f.returnType
}

func (f sysFunction) GetParams() []domain.Parameter {
	return f.params
}

func (f sysFunction) Run(values []domain.DataValue) (error, domain.DataValue) {
	// Validates the sysFunction values
	err := validateParameters(f, values, f.jef)
	if err != nil {
		return err, nil
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
			err := newJ.GetVariableManager().DeleteVariable(param.GetName())
			if err != nil {
				return err, nil
			}
		}

		newJ.GetVariableManager().RegisterVariable(param.GetName(), dType, val.GetValue())
	}

	err = f.exec(newJ)
	return err, newJ.GetFunctionReturn()
}
