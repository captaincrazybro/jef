package compilers

import (
	"fmt"
	lu "github.com/captaincrazybro/literalutil"
)

var compilers []Compiler

// Init registers all compilers
func Init() {
	compilers = []Compiler{}
	AddCompiler(Variable{})
}

// AddCompiler adds a compiler to the compiler list
func AddCompiler(c Compiler) {
	compilers = append(compilers, c)
}

// CompileLine finds the appropriate compiler and runs it
func CompileLine(s lu.String, line *int) error {
	// comment checker
	if s.ReplaceAll(" ", "").HasPrefix("//") {
		return nil
	}

	var compiler Compiler
	for i := 0; i < len(compilers); i++ {
		c := compilers[i]
		if c.Check(s) {
			compiler = c
			break
		}
	}

	// checks if a compiler exists for this line
	if compiler == nil {
		return fmt.Errorf("unexpected line at line number %d, %q", line, s)
	}

	return compiler.Run(s, line)
}

// Compiler compiler interface
type Compiler interface {
	GetName() string
	Check(lu.String) bool
	Run(lu.String, *int) error
}