// interfaces can be used with struct embedding

package main

import (
	"fmt"
	"strings"
)

type starship struct {
	laser
}

type shooter interface {
	shoot() string
}

type laser int

// method shoot on type laser
func (l laser) shoot() string {
	return strings.Repeat("pew ", int(l))
}

func main() {

	s := starship{laser(3)}
	// s is of type starship, which is a struct containing laser

	// laser has a method called shoot which returns "pew " 'l' times
	// due to method forwarding, shoot() can be called from the starship directly

	fmt.Println(s.shoot())

	// since shoot can be called from starship directly
	// starship satisfies the shooter interface
	// therefore type starship can be passed to a function which works with that interface
	shoot_strong(s)

}

// shoot_strong takes anything as an argument which satisfies the shooter interface
func shoot_strong(t shooter) {
	louder := strings.ToUpper(t.shoot())
	fmt.Println(louder)
}
