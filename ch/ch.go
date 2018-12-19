package main

import (
	"fmt"
	"strconv"
	"time"
)

func run(task_id, sleeptime int, ch chan string) {

    time.Sleep(time.Duration(sleeptime) * time.Second)
    ch <- fmt.Sprintf("task id %d , sleep %d second", task_id, sleeptime)
    return
}

func main() {
    input := []int{3, 30,80}
    ch := make(chan string)
    che := make(chan int)
    startTime := time.Now()
    fmt.Println("Multirun start")

	for _, v := range input {
		//go run(i, sleeptime, ch)
		go func(v int){
			ch <- strconv.Itoa(v)
			fmt.Println("ch", v)
		}(v)
	}

    for _, v := range input {
        //go run(i, sleeptime, ch)
        go func(v int){
			che <- v
			fmt.Println("che", v)
		}(v)
    }

    time.Sleep(time.Second)


    endTime := time.Now()
    fmt.Printf("Multissh finished. Process time %s. Number of tasks is %d", endTime.Sub(startTime), len(input))
}