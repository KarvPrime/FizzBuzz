package main

import (
	"fmt"
	"io"
)

/*
"Write a program that prints the numbers from 1 to 100.
But for multiples of three print “Fizz” instead of the number
and for the multiples of five print “Buzz”.
For numbers which are multiples of both three and five print “FizzBuzz”."
*/

func fizzBuzz(w io.Writer) {
	var err error

	three := 1
	five := 1

	for i := 1; i < 101; i++ {
		if three == 3 && five == 5 {
			_, err = fmt.Fprintln(w, "FizzBuzz")
			three = 1
			five = 1
		} else if three == 3 {
			_, err = fmt.Fprintln(w, "Fizz")
			three = 1
			five++
		} else if five == 5 {
			_, err = fmt.Fprintln(w, "Buzz")
			five = 1
			three++
		} else {
			_, err = fmt.Fprintln(w, i)
			three++
			five++
		}

		if err != nil {
			panic(err)
		}
	}
}
