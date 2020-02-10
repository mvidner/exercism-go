package listops

type binFunc func(int, int) int
type unaryFunc func(int) int
type predFunc func(int) bool
type IntList []int

func (xs IntList) Length() int {
	n := 0
	for range xs {
		n += 1
	}
	return n
}

func (xs IntList) Reverse() IntList {
	n := len(xs)
	result := make(IntList, n)
	for i, x := range xs {
		result[n-1-i] = x
	}
	return result
}

func (xs IntList) Append(ys IntList) IntList {
	result := xs
	for _, y := range ys {
		result = append(result, y)
	}
	return result
}

func (xs IntList) Concat(yss []IntList) IntList {
	result := xs
	for _, ys := range yss {
		result = result.Append(ys)
	}
	return result
}

func (xs IntList) Filter(f predFunc) IntList {
	// assume the result will be mostly empty
	result := make(IntList, 0)
	// or, mostly full:
	// result := make(IntList, 0, len(xs))
	for _, x := range xs {
		if f(x) {
			result = append(result, x)
		}
	}
	return result
}

func (xs IntList) Map(f unaryFunc) IntList {
	result := make(IntList, 0, len(xs))
	for _, x := range xs {
		result = append(result, f(x))
	}
	return result
}

func (xs IntList) Foldl(f binFunc, initial int) int {
	result := initial
	for _, x := range xs {
		result = f(result, x)
	}
	return result
}

func (xs IntList) Foldr(f binFunc, initial int) int {
	result := initial
	n := len(xs)
	for i, _ := range xs {
		result = f(xs[n-1-i], result)
	}
	return result
}
