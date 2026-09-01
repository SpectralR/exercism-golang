package advanced

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
  message string
  cows int
}

func (e *InvalidCowsError) Error() string {
  return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.message)
}


// TODO: define the 'DivideFood' function
func DivideFood(calc FodderCalculator, nbCow int) (float64, error) {
	totalFood, errTotal := calc.FodderAmount(nbCow)
	if (errTotal != nil){
		return 0, errTotal
	}

	factor, errFactor := calc.FatteningFactor()
	if (errFactor != nil){
		return 0, errFactor
	}

	return (totalFood / float64(nbCow)) * factor, nil
}


// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(calc FodderCalculator, nbCow int) (float64, error) {
	if (nbCow <= 0) {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(calc, nbCow)
}
// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(nbCows int) error {
	if (nbCows == 0) {
		return &InvalidCowsError{
			message: "no cows don't need food",
			cows: nbCows,
		}
	} else if (nbCows < 0) {
		return &InvalidCowsError{
			message: "there are no negative cows",
			cows: nbCows,
		}
	}
	return nil
}