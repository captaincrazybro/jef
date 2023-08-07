package functions

import (
	"fmt"
	"github.com/captaincrazybro/jef/domain"
)

// registerFunctions registers system functions
func (fm *functionManager) registerFunctions(jef domain.Jef) {
	fm.RegisterFunction("print", jef, nil, []domain.Parameter{parameter{name: "o", dataType: fm.jef.GetDatatypeManager().GetDatatype("any")}}, func(newJef domain.Jef) {
		//fmt.Println("asdf")
		fmt.Println(newJef.GetVariableManager().GetVariable("o").GetValue())
	})
}
