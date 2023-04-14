package functions

import (
	"fmt"
	"github.com/captaincrazybro/jef/domain"
)

// registerFunctions registers system functions
func (fm *functionManager) registerFunctions(j domain.Jef) {
	fm.RegisterFunction("print", nil, []domain.Parameter{parameter{name: "o", dataType: fm.jef.GetDatatypeManager().GetDatatype("any")}}, func(jef domain.Jef) {
		fmt.Println(j.GetVariableManager().GetVariable("o").GetValue())
	})
}
