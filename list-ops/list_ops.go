package listops

type binFunc func(int, int) int
type unaryFunc func(int) int
type predFunc func(int) bool
type IntList []int

func (xs IntList) Length() int {
	return 0
}

func (xs IntList) Reverse() IntList {
	return xs
}

func (xs IntList) Append(ys IntList) IntList {
	return xs
}

func (xs IntList) Concat(yss []IntList) IntList {
	return xs
}

func (xs IntList) Filter(f predFunc) IntList {
	return xs
}

func (xs IntList) Map(f unaryFunc) IntList {
	return xs
}

func (xs IntList) Foldl(f binFunc, initial int) int {
	return 0
}

func (xs IntList) Foldr(f binFunc, initial int) int {
	return 0
}
