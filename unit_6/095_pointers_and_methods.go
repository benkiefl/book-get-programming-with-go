// we can also use pointers in conjunction with methods
// here we have a bday_method() and a bday_function()
// both utilize a pointer to increment the age by 1

package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) bday_method() {
	p.age++
}

func main() {

	p := person{name: "John Doe", age: 23}

	fmt.Println(p.age)

	// when calling a method which utilizes pointers
	// go automatically passes the memory address
	p.bday_method() // (&p).bday_method()

	fmt.Println(p.age)

	bday_function(&p)

	fmt.Println(p.age)

}

func bday_function(p *person) {
	p.age++
}

// structures are frequently passed around with pointers
// but if a method does not need to modify the receiver then it might be best to not use a pointer
// see https://golang.org/doc/faq#methods_on_values_or_pointers
// for example methods of time.Time never use a pointer reciever, they return a new time instead
// see https://golang.org/pkg/time/
