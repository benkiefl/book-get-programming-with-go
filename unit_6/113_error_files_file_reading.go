// Go largely lets you handle errors
// many built in function return an error value
// a value of not nil indicates an error
// important fact: errors are values, think carefully what to do with them
// a large part of programming is handling errors properly(!)

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	// func ioutil.ReadDir(dirname string) ([]fs.FileInfo, error)
	files, err := ioutil.ReadDir(".") // . for the current dir
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

	fmt.Println("##################################")

	// what happens if we give ReadDir a faulty path, like just an empty string
	files, err = ioutil.ReadDir("")
	if err != nil {
		fmt.Println(err) // open : no such file or directory
		os.Exit(1)       // exit status 1
	}

	fmt.Println("execution won't reach this point")
}
