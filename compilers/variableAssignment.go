package compilers

import (
	"fmt"
	"github.com/captaincrazybro/jef/domain"
	lu "github.com/captaincrazybro/literalutil"
	"regexp"
)

// variableAssignment structure to compile variables
type variableAssignment struct {
	jef domain.Jef
}

func (v variableAssignment) GetName() string {
	return variableName
}

func (v variableAssignment) Check(s lu.String) bool {
	reg, _ := regexp.Compile("^\\S* +\\S* *= *.*$")
	return reg.MatchString(s.Tos())
}

func (v variableAssignment) Run(iter domain.LineIterator) error {
	s := iter.Current()

	// Parses in possible type, variable name, and variable value
	typeReg, _ := regexp.Compile("^([a-zA-Z0-9]*) +([a-zA-Z][a-zA-Z0-9_]*) *= *(.*)$")
	noTypeReg, _ := regexp.Compile("^([a-zA-Z][a-zA-Z0-9_]*) *= *(.*)$")
	specifiesType := typeReg.MatchString(s.Tos())

	var varType domain.DataType
	var typeName, varName, value lu.String
	if specifiesType {
		groups := typeReg.FindStringSubmatch(s.Tos())
		typeName = lu.String(groups[1])
		varName = lu.String(groups[2])
		value = lu.String(groups[3])

		// Makes sure the type exists
		varType = v.jef.GetDatatypeManager().GetDatatype(typeName.Tos())
		if varType == nil {
			return fmt.Errorf("invalid variable assignment! type does not exist")
		}

		exists := v.jef.GetVariableManager().GetVariable(varName.Tos())
		if exists != nil {
			return fmt.Errorf("invalid variable assignment! variable has already been assigned a type")
		}
	} else if noTypeReg.MatchString(s.Tos()) {
		groups := noTypeReg.FindStringSubmatch(s.Tos())
		varName = lu.String(groups[1])
		value = lu.String(groups[2])
	} else {
		return fmt.Errorf("invalid variable assignment! variable type or variable name is invalid")
	}

	// Finds the datatype
	parser, err := v.jef.GetParserManager().ParseCode(value)
	if err != nil {
		return err
	}

	// If the type was specified, checks to see if the value's type matches the
	if specifiesType && parser.GetType() != varType {
		return fmt.Errorf("invalid variable assignment! specified variable type does not match the type of the assigned value")
	}

	// trys to register a new variable
	// if the variable already exists, then it updates it instead
	exists := v.jef.GetVariableManager().RegisterVariable(varName.Tos(), parser.GetType(), parser.GetValue())
	if exists {
		err = v.jef.GetVariableManager().UpdateVariable(varName.Tos(), parser.GetType(), parser.GetValue())
		if err != nil {
			return err
		}
	}

	return nil
}
