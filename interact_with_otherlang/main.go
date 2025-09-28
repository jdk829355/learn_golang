package main

// #cgo LDFLAGS: -L. -lcfactorial
// int factorial(int);
import "C"
import "fmt"

func main() {
	fmt.Println(C.factorial(C.int(5)))
}
