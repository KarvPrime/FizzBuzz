package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func Test_fizzBuzz(t *testing.T) {
	fizzBuzz(os.Stdout)
}

func Benchmark_fizzBuzz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fizzBuzz(ioutil.Discard)
	}
}
