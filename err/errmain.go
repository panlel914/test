package rpcdemo

import "fmt"

func main() {
	fmt.Println("1111")
	defer func() {
		if r := recover(); r != nil{
			fmt.Println("recovered from ",r)
		}
	}()
	panic("333")
	fmt.Println("2222")
}
