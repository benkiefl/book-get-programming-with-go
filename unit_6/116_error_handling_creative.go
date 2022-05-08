// https://blog.golang.org/errors-are-values

// error is only a value and can be handled as such
// whether you return, exit, do nothing and when to do what
// it is up to you and if you know what you're doing it lets you write good code

package main

import (
	"fmt"
	"io"
	"os"
)

type safeWriter struct {
	// io.Writer is an interface wrapping the Write method
	w   io.Writer
	err error // zero value of error type is nil
}

func (sw *safeWriter) writeln(s string) {
	if sw.err != nil {
		return
	}
	// func fmt.Fprintln(w io.Writer, a ...interface{}) (n int, err error)
	_, sw.err = fmt.Fprintln(sw.w, s)
	// returns number of bytes written and error code
}

func main() {

	proverbs("116_proverbs.txt")

}

func proverbs(name string) error {

	f, err := os.Create(name)
	if err != nil {
		return err
	}

	defer f.Close()

	sw := safeWriter{w: f}
	// sw is a struct of type safeWriter with w being the file pointer
	// using the method writeln with fmt.Fprintln(sw.w, "string") we can write to the file
	sw.writeln("Errors are values.")
	sw.writeln("Don't just check errors, handle them gracefully.")
	sw.writeln("Don't panic.")
	sw.writeln("Make the zero value useful.")
	sw.writeln("The bigger the interface, the weaker the abstraction.")
	sw.writeln("interface{} says nothing.")
	sw.writeln("Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.")
	sw.writeln("Documentation is for users.")
	sw.writeln("A little copying is better than a little dependency")
	sw.writeln("Clear is better than clever.")
	sw.writeln("Concurrency is not parallelism.")
	sw.writeln("Don't communicate by sharing memory, share memory by communicating.")
	sw.writeln("Channels orchestrate; mutexes serialize.")
	sw.writeln("\t\t\t\t\t\t\t\t\t\t\t\t--Rob Pike, Go Proverbs")
	sw.writeln("\t\t\t\t\t\t\t\t\t\t\t\t(see go-proverbs.github.io")
	// what happens if a sw.writeln fails?
	// the error will be stored in the sw structure
	// upon checking for nil all subsequent writeln() method calls will fail
	// then the file will be closed (via defer f.Close()) and the error returned

	return sw.err
}
