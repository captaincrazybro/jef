package compilers

import (
	"github.com/captaincrazybro/jef/domain"
	lu "github.com/captaincrazybro/literalutil"
)

// for structure for the for compiler
type forLoop struct {
	jef domain.Jef
}

// GetName function to return the name of the forLoop compiler
func (fL *forLoop) GetName() string {
	return forLoopName
}

// Check checks the current line to see if it is a for statement
func (fL *forLoop) Check(s lu.String) bool {
	return s.HasPrefix("for ") || s == "for"
}

//
