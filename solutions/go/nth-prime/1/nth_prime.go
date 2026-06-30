package nthprime

import (
	"errors"

)

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n <=0 {
		return 0,errors.New("n must be positive")
	}

	number:=0

	for i:=2 ; ; i++ {
		is_prime :=true
		for j:= int(i/2) ; j > 1 ; j-- {
			if i%j == 0 {
				is_prime = false
				break
			}
		}

		if is_prime {
			number++
		}

		if number == n {
			return i , nil
		}
	}
}
