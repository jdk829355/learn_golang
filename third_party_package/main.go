package main

import (
	"fmt"
	"os"
	"strconv"

	// go get github.com/otiai10/primes로 패키지 다운 받고 import
	// 외부에 공개할 패키지면
	"github.com/otiai10/primes"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Usage:", os.Args[0], "<number>")
	}
	number, err := strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}
	f := primes.Factorize(int64(number))
	fmt.Println("primes:", len(f.Powers()) == 1)
}
