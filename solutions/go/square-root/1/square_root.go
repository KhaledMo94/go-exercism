package squareroot

import "errors"

func SquareRoot(x int) (int, error) {
	if x < 0 {
		return  0 , errors.New("must be positive integer")
	}
	if x == 0 {
		return 0 , nil
	}

	guess := x

	for {
		next := (guess + x/guess) / 2

		diff := next - guess
		if diff < 0 {
			diff = -diff
		}

		if float64(diff) < 10e-6 {
			return next , nil
		}

		guess = next
	}
}
