package romannumerals

import (
	"errors"
	"iter"
	"strings"
)

type Values struct{
	arabic int
	roman string
}

var values = []Values {
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400 ,"CD"},
	{100, "C"},
	{90,"XC"},
	{50, "L"},
	{40,"XL"},
	{10, "X"},
	{9 ,"IX"},
	{5, "V"},
	{4 , "IV"},
	{1, "I"},
}


func compositeRomanNumber(input int) iter.Seq[string] {
	return func (yield func(string) bool){
		for _,v := range values{
			for input >= v.arabic{
				if !yield(v.roman){
					return 
				}

				input -=v.arabic
			}
		}
	}
}

func ToRomanNumeral(input int) (string, error) {
	if input > 3999 || input < 1 {
		return "",errors.New("Numbers must be between 1 and 3999")
	}

	var res strings.Builder
	for part := range compositeRomanNumber(input){
		res.WriteString(part)
	}

	return res.String() , nil 
}
