package main

import "fmt"

func main() {

	var boss *string

	jane := "Jane Doe"
	john := "John Doe"

	boss = &jane

	fmt.Printf("your boss is %s\n", *boss)

	boss = &john

	fmt.Printf("your boss is %s\n", *boss)

	// our pointer boss can always point to the right person
	// if we change our variables, our pointer continues to work with the new values
	// since boss pointing to john means boss pointing to whatever is stored in john

	john = "John Bob"

	fmt.Println(john, *boss)

	// we can also use pointers to change the underlying values of course

	*boss = "John Qux"

	fmt.Printf("we renamed %s, and appointed him boss: %s\n", john, *boss)

	// as you might be able to see even at this early stage: liberal use of pointers makes programs confusing
	// passing pointers around in functions and etc means that a variable could be manipulated in a lot of places
	// might lead to cases of cumbersome debugging

	// we can have multiple pointers point to the same variable, even compare pointers using ==

	p1, p2 := &jane, &john

	fmt.Println(p1 == p2)

	p1, p2 = &jane, &jane

	fmt.Println(p1 == p2)

	// if comparing strings, the content is important, not the memory address

	str1, str2 := "hello", "hello"

	fmt.Println("string comparison:", str1 == str2)

}
