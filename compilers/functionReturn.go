package compilers

import (
	"fmt"
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
	reg, _ := regexp.Compile("^return +.*")
	return reg.MatchString(s.Tos())
}

func (fR functionReturn) Run(iter domain.LineIterator) error {
	// Checks to see if this current jef instance is a function
	if !fR.jef.IsFunction() {
		return fmt.Errorf("invalid return statement! can't return from a non-function statement")
	}

	reg, _ := regexp.Compile("^return +(.*)")
	returnStr := lu.String(reg.FindStringSubmatch(iter.Current().Tos())[1])

	// parses the return string
	returnData, err := fR.jef.GetParserManager().ParseCode(returnStr)
	if err != nil {
		return err
	}

	fR.jef.SetFunctionReturn(returnData)

	return nil
}
