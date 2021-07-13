// Package managers DEPRECATED
package managers

import (
	"github.com/captaincrazybro/jef/domain"
)

// Managers interface to store managers
type Managers interface {
	GetCompilerManager() domain.CompilerManager
	GetVariableManager() domain.VariableManager
}

// managers structure to store managers
type managers struct {
	compilerManager domain.CompilerManager
	variableManager domain.VariableManager
}

func (m managers) GetCompilerManager() domain.CompilerManager {
	return m.compilerManager
}

func (m managers) GetVariableManager() domain.VariableManager {
	return m.variableManager
}

//// New creates a new manager
//func New(j domain.Jef) Managers {
//	mz := managers{
//		compilerManager: compilers.New(j),
//		variableManager: variable.New(),
//	}
//	return mz
//}