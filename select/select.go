package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func(){
		i:=0
		for{
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()

	return out
}

func main() {
	c1, c2:= generator(), generator()

	for {
		select {
			case n:= <- c1:
				fmt.Println("c1 n:", n)
			case n:= <- c2:
				fmt.Println("c2 n:", n)

		}
	}
}
