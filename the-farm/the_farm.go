package thefarm

import (
	"errors"
	"fmt"
)

type SillyNephewError struct {
	cows int
}

func (s *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", s.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodderAmount, fodderAmountError := weightFodder.FodderAmount()

	switch {
	case fodderAmountError == ErrScaleMalfunction && fodderAmount > 0:
		return (fodderAmount * 2) / float64(cows), nil
	case fodderAmount < 0 && (fodderAmountError == nil || fodderAmountError == ErrScaleMalfunction):
		return 0, errors.New("negative fodder")
	case cows == 0:
		return 0, errors.New("division by zero")
	case cows < 0:
		return 0, &SillyNephewError{cows: cows}
	case fodderAmountError != nil:
		return 0, fodderAmountError
	default:
		return fodderAmount / float64(cows), nil
	}
}
