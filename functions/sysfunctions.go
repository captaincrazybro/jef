package functions

import (
	"fmt"
	"github.com/captaincrazybro/jef/domain"
)

// registerFunctions registers system functions
func (fm *functionManager) registerFunctions(jef domain.Jef) {
	fm.RegisterFunction("print", jef, nil, []domain.Parameter{parameter{name: "o", dataType: fm.jef.GetDatatypeManager().GetDatatype(domain.AnyDataTypeName)}}, func(newJef domain.Jef) {
		fmt.Println(newJef.GetVariableManager().GetVariable("o").GetValue())
	})
}
