package robotname

import (
	"errors"
	"fmt"
	"math/rand"
)

type Robot struct {
	name string
}

var allRobotNames map[string]bool

func RandomName() string {
	a := 'A' + rand.Intn(26)
	b := 'A' + rand.Intn(26)
	n := rand.Intn(1000)
	return fmt.Sprintf("%c%c%03d", a, b, n)
}

func UniqueName() (string, error) {
	if allRobotNames == nil {
		allRobotNames = make(map[string]bool)
	}
	if len(allRobotNames) == 26 * 26 * 10 * 10 * 10 {
		return "", errors.New("There are no other robot names left")
	}
	for {
		n := RandomName()
		known := allRobotNames[n]
		if !known {
			allRobotNames[n] = true
			return n, nil
		}
	}
}

func (r *Robot) Name() (string, error) {
	var err error
	if r.name == "" {
		n, err := UniqueName()
		if err == nil {
			r.name = n
		}

	}
	return r.name, err
}

// make *r* forget its name, next r.Name() will give it a new one
func (r *Robot) Reset() {
	r.name = ""
}
