package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	if number < 1 || number > 64{
		return 0 , errors.New("Number Out OF Range")
	}
	return  uint64(math.Pow(2.0 , float64(number - 1))) , nil
}

func Total() uint64 {
	var sum uint64
	for i := 1; i < 65; i++ {
		v , _ := Square(i)	
		sum = sum + v
	}
	return sum
}
