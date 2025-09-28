package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(1 * time.Millisecond)
			fmt.Println(i)
		}(i)
	}
	time.Sleep(100 * time.Millisecond)
}
