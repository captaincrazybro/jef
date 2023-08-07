package util

import "github.com/captaincrazybro/jef/domain"

// ParseConditionalValue parses any datatype (from the conditional block) and converts it to a boolean
func ParseConditionalValue(val domain.DataValue, jef domain.Jef) bool {
	// switch statement for the type of the DataValue
	switch val.GetType() {
	// Case for the int type
	case jef.GetDatatypeManager().GetDatatype("boolean"):
		{
			return val.GetValue().(bool)
		}
	case jef.GetDatatypeManager().GetDatatype("int"):
		{
			return val.GetValue() != 0
		}
	// Case for the remaining types
	default:
		{
			return val.GetValue() != nil
		}
	}
}
