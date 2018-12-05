package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age int
}

func (u User) Print(){
	fmt.Println(u)
}

func (u User) PrintByArg(name string){
	fmt.Println(name)
}

func main() {
	//var num float64 = 3.1415926
	//fmt.Println("type:", reflect.TypeOf(num))
	//fmt.Println("type:", reflect.ValueOf(num))

	u := User{"peter",10}
	fmt.Println("type:", reflect.TypeOf(u))
	fmt.Println("type:", reflect.ValueOf(u))

	//v := reflect.ValueOf(num)
	// t := reflect.TypeOf(num)
	fmt.Println("--------------------------")
	fmt.Println("print value and method")
	print(u)
	fmt.Println("--------------------------")
	fmt.Println("set value")
	var num float64 = 3.1415926
	fmt.Println("old value : ", num)
	v := reflect.ValueOf(&num)
	newValue := v.Elem()
	newValue.SetFloat(35)
	fmt.Println("new value : ", num)

	fmt.Println()
	fmt.Println("-----------------------------")
	fmt.Println("call method")
	method(u)
}

func method (input interface{}){
	v := reflect.ValueOf(input)

	m := v.MethodByName("PrintByArg")
	args := []reflect.Value{reflect.ValueOf("wudebao")}
	m.Call(args)

	m2 :=v.MethodByName("Print")
	m2.Call(nil)
}

func set(input *interface{}){
	v := reflect.ValueOf(input)
	newValue := v.Elem()

	newValue.SetInt(35)

}

func print(input interface{}){
	t := reflect.TypeOf(input)
	v := reflect.ValueOf(input)

	for i:=0;i<t.NumField();i++{
		f :=t.Field(i)
		v := v.Field(i).Interface()
		fmt.Println("field name:",f.Name," field type:", f.Type," value: ", v)
	}

	for i:=0;i<t.NumMethod();i++{
		f := t.Method(i)
		fmt.Println("method name: ", f.Name ,"  method type:  ", f.Type)
	}
}
