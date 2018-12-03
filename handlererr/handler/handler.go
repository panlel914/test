package main

import (
	"errors"
	"fmt"
	"net/http"
)

func handler (writer http.ResponseWriter, request *http.Request) error{
	return errors.New("peter create a new error")
}

type handlerErr func(writer http.ResponseWriter, request *http.Request) error

func errHelp(handler handlerErr) func (writer http.ResponseWriter, request *http.Request){
	return func(writer http.ResponseWriter, request *http.Request) {
		err :=handler(writer,request)
		fmt.Println(err.Error())
	}
}

func main() {
	fmt.Println("asdf")

	http.HandleFunc("/test/", errHelp(handler))

	http.ListenAndServe(":4567", nil)
}
