package main

import (
	"fmt"
	"sync"
)

func worker(id int,c chan int, wg *sync.WaitGroup){

	for n:= range c{
		fmt.Println("id:", id, "number:", n)
		wg.Done()
	}
}

func main() {
	//var ch chan int
	//ch = make(chan int)
	//done := make(chan bool)
	var chs [10]chan int
	var wg sync.WaitGroup

	for i:=0;i<10;i++ {
		chs[i] = make(chan int)
		go worker(i, chs[i], &wg)
		wg.Add(1)
	}

	for i,c:= range chs{
		c<-i+1
	}

	wg.Wait()

}
