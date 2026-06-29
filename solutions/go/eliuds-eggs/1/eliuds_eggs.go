package eliudseggs

import "strconv"

func EggCount(displayValue int) int {
	binary := strconv.FormatInt(int64(displayValue) , 2)

	s := string(binary)

	r:=0

	for _,v := range s {
		if v == '1'{
			r++
		}
	}

	return  r
}
