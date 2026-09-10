package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(fc FodderCalculator, numberOfCows int) (float64, error) {

	totalFodder, errDeterminingAmount := fc.FodderAmount(numberOfCows)
	if errDeterminingAmount != nil {
		return 0, errDeterminingAmount
	}

	fattening, errDeterminingFactor := fc.FatteningFactor()
	if errDeterminingFactor != nil {
		return 0, errDeterminingFactor
	}

	return totalFodder / float64(numberOfCows) * fattening, nil
}

func ValidateInputAndDivideFood(fc FodderCalculator, numberOfCows int) (float64, error) {
	if numberOfCows > 0 {
		return DivideFood(fc, numberOfCows)
	}
	return 0, errors.New("invalid number of cows")
}

func ValidateNumberOfCows(numberOfCows int) error {
	switch {
	case numberOfCows < 0:
		return &InvalideCowsError{
			cowCount: numberOfCows,
			message:  "there are no negative cows",
		}
	case numberOfCows == 0:
		return &InvalideCowsError{
			cowCount: numberOfCows,
			message:  "no cows don't need food",
		}
	default:
		return nil
	}
}

type InvalideCowsError struct {
	cowCount int
	message  string
}

func (e *InvalideCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cowCount, e.message)
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
