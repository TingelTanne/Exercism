package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	var counter int = 0

	//first Step: Check if n is positiv
	if n <= 0 {
		return -1, errors.New("The number can't be less then 1!")
	}

	//counter loop to determin the number of itterations
	//needed to reach 1
	for ; n != 1; counter++ {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}
	}

	return counter, nil
}
