package robotname

import (
	"fmt"
	"math/rand"
)

type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		a := 'A' + rand.Intn(26)
		b := 'A' + rand.Intn(26)
		n := rand.Intn(1000)
		r.name = fmt.Sprintf("%c%c%03d", a, b, n)
	}
	return r.name, nil
}
func (r *Robot) Reset() {
	r.name = ""
}
