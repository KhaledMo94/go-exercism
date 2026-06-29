package isbnverifier

import (
	"math"
	"strings"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn , "-","")

	if len(isbn) != 10 {
		return false
	}

	r :=0
	factor := 10

	for i , num := range isbn {
		if num == 'X' && i == 9 {
			r += 10 * factor
		}else{
			r += int(num - '0') * factor
		}
		factor--
	}

	return  math.Mod(float64(r),11) == 0
}
