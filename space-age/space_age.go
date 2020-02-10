package space

import "math"

type Planet string

// how many *planet*-orbits are there in *seconds*
func Age(seconds float64, planet Planet) float64 {
	// in Earth years
	planetOrbitYears := map[Planet]float64 {
		"Mercury": 0.2408467,
		"Venus": 0.61519726,
		"Earth": 1.0,
		"Mars": 1.8808158,
		"Jupiter": 11.862615,
		"Saturn": 29.447498,
		"Uranus": 84.016846,
		"Neptune": 164.79132,
	}

	/*
	a plain
		orbitYears := planetOrbitYears[planet]
	would return 0.0 so we would return +Inf. NaN is better
	*/

	orbitYears, ok := planetOrbitYears[planet]
	if ok {
		return seconds / (orbitYears * 365.25 * 24 * 60 * 60)
	} else {
		return math.NaN()
	}
}
