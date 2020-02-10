package gigasecond

import "time"

// return *t* advanced by a billion seconds
func AddGigasecond(t time.Time) time.Time {
	const giga = 1000 * 1000 * 1000 // can't write 1_000_000_000
	return t.Add(giga * time.Second)
}
