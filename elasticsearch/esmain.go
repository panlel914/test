package main

import (
	"fmt"
	"gopkg.in/olivere/elastic.v3"
)

func main() {
	Save()
}

type student struct {
	name string
	age int
	class string
}

func Save(){
	client,err := elastic.NewClient(
		elastic.SetSniff(false))

	if err != nil{
		panic("client is nil")
	}
	//stu1 := student{
	//	name:"peter",
	//	age:30,
	//	class:"class1",
	//}
	resp,err :=client.Index().
		Index("testes").
		Type("test").
		BodyJson(`{"name":"brian","age":4,"class":"small6"}`).
		Do()

	if err != nil{
		panic("error")
	}

	fmt.Println(resp)
}
