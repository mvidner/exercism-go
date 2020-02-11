package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ChanFrequency(ch chan FreqMap, s string) {
	ch <- Frequency(s)
}

func UpdateMap(m FreqMap, increment FreqMap) {
	for k, v := range increment {
		m[k] += v
	}
}

func ConcurrentFrequency(ss []string) FreqMap {
	m := FreqMap{}
	ch := make(chan FreqMap)
	// start
	for _, s := range ss {
		go ChanFrequency(ch, s)
	}
	// collect
	for i := 0; i < len(ss); i++ {
		mm := <-ch
		UpdateMap(m, mm)
	}
	return m
}
