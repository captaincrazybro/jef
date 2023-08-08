package typeparsers

import (
	"github.com/captaincrazybro/jef/domain"
	lu "github.com/captaincrazybro/literalutil"
	"regexp"
)

type Math struct {
	jef domain.Jef
}

func (mD Math) GetType() domain.DataType {
	return nil
}

func (mD Math) Check(s lu.String) bool {
	r, _ := regexp.Compile("^[^\"].*[+\\-/*].*")
	return r.MatchString(s.Tos())
}

func (mD Math) GetValue(s lu.String) (domain.DataValue, error) {

}
