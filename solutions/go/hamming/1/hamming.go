package hamming

import "errors"

func Distance(a, b string) (int, error) {
	var differences int = 0
	if len(a) != len(b) {
		return -1, errors.New("Can't compare sequences with defferent lenghths!")
	} else {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				differences++
			}
		}
	}
	return differences, nil
}
