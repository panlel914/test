package main

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sony/gobreaker"
	"net/http"
	"test/break/mybreak"
	"time"
)

var cb *gobreaker.CircuitBreaker

func init() {
	var st gobreaker.Settings
	st.Name = "HTTP GET"
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		fmt.Println(counts.Requests)
		return counts.Requests >= 3 && failureRatio >= 0.6
	}

	cb = gobreaker.NewCircuitBreaker(st)
}

// Get wraps http.Get in CircuitBreaker.
func Get(url string) ([]byte, error) {
	body, err := cb.Execute(func() (interface{}, error) {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()
		//body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return nil, errors.New("aaa")
	})
	if err != nil {
		return nil, err
	}

	return body.([]byte), nil
}

func main() {
	set := mybreak.MySetting{MaxError:2, Inv:2}
	mybreak := mybreak.NewMyBreak(set)

	fmt.Println("begin")
	for i:=0;i<10;i++{
		if i < 2 {
			v,err := mybreak.Excute(func() (interface{}, error) {
				return "success", nil
			})
			fmt.Println("value is ", v," err", err)
		}else if i < 4{
			v,err := mybreak.Excute(func() (interface{}, error) {
				return "err", errors.New("error")
			})
			fmt.Println("value is ", v," err", err)
		}else if i < 6{
			time.Sleep(time.Second*3)
			v,err := mybreak.Excute(func() (interface{}, error) {
				return "err", errors.New("error")
			})
			fmt.Println("value is ", v," err", err)
		}else{
			v,err := mybreak.Excute(func() (interface{}, error) {
				return "success", nil
			})
			fmt.Println("value is ", v," err", err)
		}



	}
	time.Sleep(time.Second*2)
	v,err := mybreak.Excute(func() (interface{}, error) {
		return "success", nil
	})
	fmt.Println("value is ", v," err", err)

	fmt.Println("all count:", mybreak.Count)
	fmt.Println("end")

	//for i:=0;i<10;i++ {
	//	func(i int){
	//		_, err := Get("http://www.baidu.com/robots.txt")
	//		if err != nil {
	//			// log.Fatal(err)
	//		}
	//
	//		fmt.Println(i)
	//	}(i)
	//}
	//
	//time.Sleep(time.Second*4)
}