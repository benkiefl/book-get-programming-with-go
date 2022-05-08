// err is a special built-in type for errors

package main

import (
	"fmt"
	"os"
)

func main() {
	err := proverbs("114_proverbs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// proverbs, given a file name, creates a file of said name full of Go proverbs ;
// returns nil upon successful completion, error message upon failure
func proverbs(name string) error {
	// func os.Create(name string) (*os.File, error)
	f, err := os.Create(name)
	if err != nil {
		// notice we return without f.Close()
		// that's because: no need to close the file it we had an error creating it
		return err
	}
	// if file already exists it will be over written

	// defer is not strictly needed in my version of the program
	// at this stage we know the file is open
	defer f.Close()
	// defer means f.Close() will be run before _any_ return following
	// so no matter what executional path we take here, the file will be closed properly

	// func fmt.Fprintln(w io.Writer, a ...interface{}) (n int, err error)
	// returns the number of bytes written and any write error encountered
	_, err = fmt.Fprintln(f, "Errors are values.\n"+
		"Don't just check errors, handle them gracefully.\n"+
		"Don't panic.\n"+
		"Make the zero value useful.\n"+
		"The bigger the interface, the weaker the abstraction.\n"+
		"interface{} says nothing.\n"+
		"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.\n"+
		"Documentation is for users.\n"+
		"A little copying is better than a little dependency.\n"+
		"Clear is better than clever.\n"+
		"Concurrency is not parallelism.\n"+
		"Don't communicate by sharing memory, share memory by communicating.\n"+
		"Channels orchestrate; mutexes serialize.\n"+
		"\t\t\t\t\t\t\t\t\t\t\t\t--Rob Pike, Go Proverbs\n"+
		"\t\t\t\t\t\t\t\t\t\t\t\t(see go-proverbs.github.io")
	// about Go proverbs, also see https://www.youtube.com/watch?v=PAAkCSZUG1c
	// I thought above way was the best compromise
	// I hate the formatting with multi-line string literals https://play.golang.org/p/EDAIANoeiET
	// my chosen method at least the indentation doesn't look off (even though with \n and + it's more verbose)

	// commented out here, we used defer f.Close() above
	// f.Close()

	if err != nil {
		return err
	}
	return nil
}
