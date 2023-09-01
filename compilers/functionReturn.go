package compilers

import (
	"github.com/captaincrazybro/jef/domain"
	lu "github.com/captaincrazybro/literalutil"
	"regexp"
)

// mathAssignment structure to compile variables
type functionReturn struct {
	jef domain.Jef
}

func (fR functionReturn) GetName() string {
	return functionReturnName
}

func (fR functionReturn) Check(s lu.String) bool {
	reg, _ := regexp.Compile("^\\Sreturn +.*")
	return reg.MatchString(s.Tos())
}

func (fR functionReturn) Run(iter domain.LineIterator) error {
	return nil
}
