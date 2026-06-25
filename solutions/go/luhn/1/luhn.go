package luhn

import "strings"

func Valid(id string) bool {

	id = strings.ReplaceAll(id , " ", "")

	if len(id) <2 {
		return  false
	}

	sum :=0 
	double := false

	for i := len(id) -1 ; i >= 0 ; i-- {
		if id[i] < '0' || id[i] > '9'{
			return false
		}

		d := int(id[i] - '0')

		if double {
			d*=2
			if (d > 9){
				d-=9
			}
		}

		sum += d
		double =! double
	}

	return  sum % 10 == 0
}
