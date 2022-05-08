// pointer is a variable that points to the address of another variable
// pointers are a form of indirection

package main

import "fmt"

func main() {

	// go uses the well established C-syntax for pointer

	// address operator &
	var v int = 5
	fmt.Println(&v) // memory address of variable v

	p := &v
	fmt.Println(v, p)

	// the oppositve of & is *
	// called dereferencing (provide the value a memory address refers to)
	fmt.Println(v, *p)

	// *&v is basically same as simply using v
	// * dereference
	// & memory address of v
	fmt.Println(*&v)

	// p is a pointer of type *int
	fmt.Printf("pointer p is of type %T\n", p)
	fmt.Println()

	country := "Russia"
	pointer_to_country := &country

	fmt.Printf("country %s stored at address %v\n", country, pointer_to_country)
	fmt.Printf("dereferncing the pointer gives us our country %s\n", *pointer_to_country)
	fmt.Printf("memory address of our pointer is %v\n", &pointer_to_country)
	fmt.Printf("our pointer here is of type %T\n", pointer_to_country)

	// a *int pointer can't be used as a *string ... pointer
	// I can make pointer p from above point to x (since p is a pointer of type *int)
	var x int = 5
	p = &x
	// I can't however make pointer p (of type *int) point to country (type string)
	// p = &country ; cannot use &country (value of type *string) as *int value in assignment

	// Pointers can be used where any type may be used
	// 		variable declarations
	//		function parameters
	//		return types
	//		structure field types
	//	et cetera

}
