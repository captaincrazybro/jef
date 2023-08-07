package functions

import (
	"fmt"
	"github.com/captaincrazybro/jef/domain"
)

// registerFunctions registers system functions
func (fm *functionManager) registerFunctions(jef domain.Jef) {
	fm.RegisterFunction("print", jef, nil, []domain.Parameter{parameter{name: "o", dataType: fm.jef.GetDatatypeManager().GetDatatype("any")}}, func() {
		fmt.Println(jef.GetVariableManager().GetVariable("o").GetValue())
	})
}
