package util

// RuneArrContains checks to see if the given array of runes contains the given rune
func RuneArrContains(rz []rune, r1 rune) bool {
	for _, r2 := range rz {
		if r1 == r2 {
			return true
		}
	}

	return false
}

// StringArrContains checks to see if the given array of strings contains the given string
func StringArrContains(sz []string, s1 string) bool {
	for _, s2 := range sz {
		if s1 == s2 {
			return true
		}
	}

	return false
}
