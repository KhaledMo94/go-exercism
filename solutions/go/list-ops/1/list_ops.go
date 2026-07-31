package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	acc := initial
	for _ , item := range s {
		acc = fn(acc , item)
	}
	return acc
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	acc := initial
	for i:=len(s)-1;i>=0;i--{
		acc = fn(s[i], acc)
	}
	return acc
}

func (s IntList) Filter(fn func(int) bool) IntList {
	var res IntList
	for i,v := range s {
		if fn(s[i]) {
			res = append(res, v)
		}
	}
	return res
}

func (s IntList) Length() int {
	sum :=0 

	for sum = range s{
		sum++
	}

	return sum
}

func (s IntList) Map(fn func(int) int) IntList {
	var res IntList

	for _,v := range s {
		res = append(res, fn(v))
	}

	return res
}

func (s IntList) Reverse() IntList {
	var res IntList

	for i:=len(s)-1;i>=0;i--{
		res = append(res, s[i])
	}

	return res
}

func (s IntList) Append(lst IntList) IntList {
	for i := 0; i < len(lst); i++ {
		s = append(s, lst[i])
	}

	return  s
}

func (s IntList) Concat(lists []IntList) IntList {
	for _,list := range lists{
		for i:=0 ; i<len(list) ; i++{
			s = append(s, list[i])
		}
	}

	return s
}
