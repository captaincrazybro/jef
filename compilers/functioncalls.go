package compilers

import (
	"fmt"
	"github.com/captaincrazybro/jef/domain"
	lu "github.com/captaincrazybro/literalutil"
	"regexp"
	"strings"
)

// variable structure to compile variables
type functioncalls struct {
	jef domain.Jef
}

func (v functioncalls) GetName() string {
	return functionCallName
}

func (v functioncalls) Check(s lu.String) bool {
	// Validates the function call without the parameters
	r, _ := regexp.Compile("[a-zA-Z][a-zA-Z0-9]*\\((.*)\\)")
	return r.MatchString(s.Tos())
}

func (v functioncalls) Run(s lu.String, line *int) error {
	// Gets the function name and parameter details
	r, _ := regexp.Compile("([a-zA-Z][a-zA-Z0-9]*)\\((.*)\\)")
	subz := r.FindStringSubmatch(s.Tos())
	name := subz[1]

	// Checks to see if the function exists
	f := v.jef.GetFunctionManager().GetFunction(name)
	if f == nil {
		return fmt.Errorf("invalid function name! there is no function named \"%s\"", name)
	}

	// Parses parameters
	rawParams := subz[2]
	params := strings.Split(rawParams, ",")
	var paramTypes []domain.TypeParser
	var paramValues []interface{}
	for _, param := range params {
		// Finds the datatype based on the value passed in
		dt, err := v.jef.GetDatatypeManager().FindDatatype(lu.String(param))
		if err != nil {
			return err
		}

		// Parses the datatype value
		paramTypes = append(paramTypes, dt)
		value, err := dt.GetValue(lu.String(param))
		if err != nil {
			return err
		}

		paramValues = append(paramValues, value)
	}

	// Runs the function
	err := f.RunExec(paramValues, paramTypes, v.jef)
	return err
}
