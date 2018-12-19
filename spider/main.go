package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {

	startTime := time.Now()
	fmt.Println(startTime)
	//fmt.Println("start ")
	//for i:=0;i<1000;i++{
	//	rand.Seed(time.Now().UnixNano())
	//	d := rand.Intn(500)
	//	time.Sleep(time.Millisecond * time.Duration(d))
	//	fmt.Println("i:",i," d:",d)
	//}
	en := NewEngine()
	se := Seed{engin:en}
	se.Run()

	endTime := time.Now()
	fmt.Printf("Multissh finished. Process time %s. Number of tasks is %d", endTime.Sub(startTime), 1000)
}

type Req struct {
	url string
}

type Res struct {
	result string
}


type Seed struct {
	engin *Engine
}

func (s *Seed) Run(){

	go func(){
		// 分发请求
		for i:=0; i< 1000; i++{
			r := Req{url:"Num" + strconv.Itoa(i) + "request"}
			s.engin.in <- r
		}
	}()
	s.engin.run()
}

type Engine struct {
	in chan Req
	out chan Res
}

func (e *Engine) run(){
	CreateWork(e.in, e.out, 1000)
	for{
		select {
			//case i := <- e.in:
			//	go func(){
			//		DoWork(i,e.out)
			//	}()
			case o := <- e.out:
				DoResponse(o)
			default:
				time.Sleep(time.Millisecond*100)
		}
	}
}

func DoResponse(res Res){
	fmt.Println(res.result)
}

func NewEngine() *Engine{
	var result = new(Engine)
	result.in = make(chan Req)
	result.out = make(chan Res)

	return result
}

func CreateWork(in chan Req, out chan Res, count int){
	for i:=0;i<count;i++ {
		go func() {
			for {
				r := <-in
				rand.Seed(time.Now().UnixNano())
				i := rand.Intn(500)
				time.Sleep(time.Millisecond * time.Duration(i))
				out <- Res{result: r.url + "'s response"}
			}
		}()
	}
}

//func DoWork(r Req, och chan Res) {
//	// 模拟请求
//	rand.Seed(time.Now().UnixNano())
//	i := rand.Intn(500)
//	time.Sleep(time.Millisecond * time.Duration(i))
//	och <- Res{result:r.url + "'s response"}
//}