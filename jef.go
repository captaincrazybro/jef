package jef

import (
	"github.com/captaincrazybro/jef/compilers"
	"github.com/captaincrazybro/jef/variable"
	lu "github.com/captaincrazybro/literalutil"
	c "github.com/captaincrazybro/literalutil/console"
)

func init() {
	compilers.Init()
	variable.Init()
	c.SetErrorPrefix("Error: ")
}

// Moto prints the moto for the Jef programming language
func Moto() {
	c.Pln("My name is Jeff!")
}

// Check checks the code for errors and pre registers functions
// TODO: create checking stuff
func Check(s string) {
	//code := lu.String(s)
}

// Run runs the code
func Run(s string) {
	code := lu.String(s)
	lines := code.Split("\n")

	for i := 0; i < lines.Len(); i ++ {
		lineValue := lines[i]
		err := compilers.CompileLine(lineValue, &i)
		if err != nil {
			c.Fln(err)
		}
	}
}