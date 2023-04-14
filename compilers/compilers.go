package compilers

import (
	"fmt"
	"github.com/captaincrazybro/jef/domain"
	"github.com/captaincrazybro/jef/util"
	lu "github.com/captaincrazybro/literalutil"
)

// compilerManager structure to store compilerManager instance
type compilerManager struct {
	compilers []domain.Compiler
}

// New creates a new CompilerManager instance
func New(j domain.Jef) domain.CompilerManager {
	cz := compilerManager{}
	cz.registerCompilers(j)
	return &cz
}

// Init registers all compilerManager
func (cz *compilerManager) registerCompilers(j domain.Jef) {
	cz.AddCompiler(variableAssignment{j})
	cz.AddCompiler(functioncalls{j})
}

// AddCompiler adds a compilers to the compilers list
func (cz *compilerManager) AddCompiler(c domain.Compiler) {
	cz.compilers = append(cz.compilers, c)
}

// CompileLine finds the appropriate compilers and runs it
func (cz compilerManager) CompileLine(s lu.String, line *int) error {
	// Trims the whitespaces around
	s = util.TrimWhitespaces(s)
	// comment checker
	if s.ReplaceAll(" ", "").HasPrefix("//") {
		return nil
	}

	// Checks if empty line
	if s == "" {
		return nil
	}

	var compiler domain.Compiler
	for i := 0; i < len(cz.compilers); i++ {
		c := cz.compilers[i]
		if c.Check(s) {
			compiler = c
			break
		}
	}

	// checks if a compilers exists for this line
	if compiler == nil {
		return fmt.Errorf("unexpected line at line number %d, %q", *line+1, s)
	}

	return compiler.Run(s, line)
}
