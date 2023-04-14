package compilers

import "github.com/captaincrazybro/jef/domain"

// ifElse structure for the ifElse compiler
type ifElse struct {
	jef domain.Jef
}

// GetName function to return the name of the ifElse compiler
func (iE *ifElse) GetName() string {
	return ifElseName
}
