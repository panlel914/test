package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(time.Second)
	for i:=0;i<1000;i++{

		go func() {
			<-tick
			fmt.Println("aa")
		}()
	}

	time.Sleep(time.Second*100)
}
