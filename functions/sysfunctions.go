package functions

import (
	"github.com/captaincrazybro/jef/domain"
	c "github.com/captaincrazybro/literalutil/console"
)

// registerFunctions registers system functions
func (fm *functionManager) registerFunctions(j domain.Jef) {
	fm.RegisterFunction("print", nil, []domain.Parameter{parameter{name: "o", dataType: fm.jef.GetDatatypeManager().GetDatatype("string")}}, func(jef domain.Jef) {
		c.Pln(j.GetVariableManager().GetVariable("o").GetValue())
	})
}
