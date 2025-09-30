package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(1 * time.Millisecond)
			fmt.Println(i)
		}()
	}
	time.Sleep(100 * time.Millisecond)
}
