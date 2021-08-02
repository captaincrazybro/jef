package datatypes

import (
	"fmt"
	lu "github.com/captaincrazybro/literalutil"
)

type String struct {

}

func (sD String) GetName() string {
	return stringDatatypeName
}

func (sD String) Check(s string) bool {
	S := lu.String(s)

	return S.Contains(`"`) && S.HasPrefix(`"`) && S.HasSuffix(`"`)
}

func (sD String) GetValue(s string) (interface{}, error) {
	S := lu.String(s)

	quoteCount := 0
	for i, v := range S {
		if string(v) == `"` && string(S[i - 1]) != `\` {
			quoteCount++
		}
	}
	if quoteCount != 2 {
		return nil, fmt.Errorf("extra `\"` found in provided datatype string; did you mean to put a \"\\\" before the quote?")
	}

	return S.ReplaceAll(`"`, ""), nil
}