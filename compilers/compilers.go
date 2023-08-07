package compilers

import (
	"fmt"
	"github.com/captaincrazybro/jef/domain"
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
	cz.AddCompiler(&ifElse{j})
	cz.AddCompiler(functioncalls{j})
}

// AddCompiler adds a compilers to the compilers list
func (cz *compilerManager) AddCompiler(c domain.Compiler) {
	cz.compilers = append(cz.compilers, c)
}

// CompileLine finds the appropriate compilers and runs it
func (cz *compilerManager) CompileLine(iter domain.LineIterator) error {
	s := iter.Current()
	// comment checker
	if s.ReplaceAll(" ", "").HasPrefix("//") {
		return nil
	}

	// Checks if empty lines lu.String, line *int
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
		return fmt.Errorf("unexpected line at line number %d, %q", iter.Index()+1, s)
	}

	return compiler.Run(iter)
}
