package util

func RuneArrContains(rz []rune, r1 rune) bool {
	for _, r2 := range rz {
		if r1 == r2 {
			return true
		}
	}

	return false
}
