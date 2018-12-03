package main

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sony/gobreaker"
	"net/http"
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
	for i:=0;i<10;i++ {
		func(i int){
			_, err := Get("http://www.baidu.com/robots.txt")
			if err != nil {
				// log.Fatal(err)
			}

			fmt.Println(i)
		}(i)
	}

	time.Sleep(time.Second*4)
}