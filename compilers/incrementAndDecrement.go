package compilers

import (
	"github.com/captaincrazybro/jef/domain"
	lu "github.com/captaincrazybro/literalutil"
)

// incrementAndDecrement structure to compile variables
type incrementAndDecrement struct {
	jef domain.Jef
}

func (i incrementAndDecrement) GetName() string {
	return incrementAndDecrementName
}

func (i incrementAndDecrement) Check(s lu.String) bool {
	return i.jef.GetParserManager().GetParser(domain.IncrementParserName).Check(s) || i.jef.GetParserManager().GetParser(domain.DecrementParserName).Check(s)
}

func (i incrementAndDecrement) Run(iter domain.LineIterator) error {
	s := iter.Current()
	// Runs the type parser for either increment or decrement and ignores the dataValue
	if i.jef.GetParserManager().GetParser(domain.IncrementParserName).Check(s) {
		_, err := i.jef.GetParserManager().GetParser(domain.IncrementParserName).GetValue(s)
		return err
	} else {
		_, err := i.jef.GetParserManager().GetParser(domain.DecrementParserName).GetValue(s)
		return err
	}
}
