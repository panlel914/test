package main

import (
	"feg/fegLib/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"test/mycache"
	"time"
)

type Student struct {
	name string
	age int
}

func (s *Student)SetName(name string){
	s.name = name
}

func main() {
	//fmt.Println("Starting test...")
	//ml := NewMapList()
	//var a, b, c Keyer
	//a = &Elements{"Alice"}
	//b = &Elements{"Conrad"}
	//c = &Elements{"bbbbb"}
	//ml.Push(a)
	//ml.Push(b)
	//ml.Push(c)
	//cb := func(data Keyer) {
	//	fmt.Println(ml.dataMap[data.GetKey()].Value.(*Elements).value)
	//}
	//fmt.Println("Print elements in the order of pushing:")
	//ml.Walk(cb)
	//fmt.Printf("Size of MapList: %d \n", ml.Size())
	//ml.Remove(b)
	//fmt.Println("After removing b:")
	//ml.Walk(cb)
	//fmt.Printf("Size of MapList: %d \n", ml.Size())

	//aaa,_ := GetEveryDay("2018-09-25", "2018-09-26")

	//fmt.Println(aaa)
	//var aaa float64 = 1.323
	//
	//fmt.Println("type: ", reflect.TypeOf(aaa))
	//fmt.Println("value: ", reflect.ValueOf(aaa))

	//stu := new(Student)
	//stu.name = "Peter"
	//stu.age = 32
	//fmt.Println("type: ", reflect.TypeOf(stu))
	//fmt.Println("value: ", reflect.ValueOf(stu))

	//pointer := reflect.ValueOf(&aaa)
	//value := reflect.ValueOf(aaa)

	//convertPointer := pointer.Interface().(*float64)
	//caluePointer := value.Interface().(float64)

	//fmt.Println(convertPointer)
	//fmt.Println(caluePointer)


	//bbb := reflect.ValueOf(&stu)
	//typeaaa := reflect.TypeOf(stu)
	//
	//for i := 0; i < typeaaa.NumMethod(); i++ {
	//	m := typeaaa.Method(i)
	//	fmt.Printf("%s: %v\n", m.Name, m.Type)
	//}

	//user := User{1, "Allen.Wu", 25}
	//getValue := reflect.ValueOf(user)
	//methodValue := getValue.MethodByName("ReflectCallFuncHasArgs")
	//args := []reflect.Value{reflect.ValueOf("wudebao"), reflect.ValueOf(30)}
	//methodValue.Call(args)

	//student := Student{name:"ppp", age: 30}
	//value := reflect.ValueOf(&student)
	//method := value.MethodByName("SetName")
	//args2 := []reflect.Value{reflect.ValueOf("111")}
	//method.Call(args2)
	//fmt.Println(student)


	//methodValue := bbb.MethodByName("SetName")
	//args := []reflect.Value{reflect.ValueOf("wokao")}
	//methodValue.Call(args)

	//ccc :=bbb.Elem()
	//ccc.Interface().(*Student).name = "Peter2"
	//println(ccc.CanSet())
	//println(stu.name)

	//fmt.Println("aaa=",aaa)
	// fmt.Println(GetEveryDayStr("2018-10-01",-30))
	// ws消息测试
	// ws.StartWS()
	//u := User{Name:"peter", Age:30, Id:1}
	//uf := UserF{U:u}
	//fmt.Println(uf)
	//ChangeZhi(&uf,"lind")
	//fmt.Println(uf)
	var array1 [5]int
	array2 := [3]int{1,2,3}
	array3 := [...]int{1,2,3,4,3,5,6}

	for _,num:= range array3{
		fmt.Println(num)
	}


	fmt.Println(array1, array2, array3)
}

type UserF struct {
	U User
}

func ChangeZhi(u *UserF, name string){
	u.U.Name = name
}

func ChangeUserFName(u UserF, name string){
	u.U.Name = name
}

func ChangeUserName(user User, name string){
	user.Name = name
}

func GetNumber(op func(int, int) int, a,b int)(int){
	return op(a,b)
}

func GetEveryDayStr(start string ,number int)(list []string, err error) {
	var(
		startTime time.Time
		endTime time.Time
	)
	if number > 0 {
		startTime, err = time.Parse("2006-01-02", start)
		endTime = startTime.AddDate(0, 0, number)
	}else{
		endTime, err = time.Parse("2006-01-02", start)
		startTime = endTime.AddDate(0, 0, number)
	}

	if err != nil {
		return
	}

	for {
		if startTime.Equal(endTime) {
			break
		}
		str := utils.GetStringFromTime(startTime)
		list = append(list, str)
		startTime = startTime.AddDate(0, 0, 1)
	}

	return
}

func GetDateByString(str string) (date string, ok bool) {
	retTime, err := time.Parse("2006-01-02", str)
	if err == nil{
		date = retTime.Format("2006-01-02")
	}
	ok = err == nil

	return
}

type User struct {
	    Id   int
	    Name string
	    Age  int
	}

func (u User) ReflectCallFuncHasArgs(name string, age int) {
	fmt.Println("ReflectCallFuncHasArgs name: ", name, ", age:", age, "and origal User.Name:", u.Name)
}

func (u User) ReflectCallFuncNoArgs() {
	fmt.Println("ReflectCallFuncNoArgs")
}

type AnchorIncomeInfo struct {
	InComeList map[string] float64
	TotalInCome float64
}

func GetEveryDay(start string ,end string)(info AnchorIncomeInfo, err error) {
	startTime, err := time.Parse("2006-01-02", start)
	endTime, err := time.Parse("2006-01-02", end)
	if err != nil {
		return
	}

	info.InComeList = make(map[string]float64)
	info.TotalInCome = 0

	for {
		str := utils.GetStringFromTime(startTime)
		info.InComeList[str] = 0
		if startTime.Equal(endTime) {
			break
		}
		startTime = startTime.AddDate(0, 0, 1)
	}

	return
}

//func main(){
//	r := &log2.ReadFromFile{}
//	w := &log2.WriteToFile{}
//
//	log := &log2.LogProcess{
//		Rp: make(chan string),
//		Pw: make(chan string),
//		ReaderObj:r,
//		WriterOjb:w,
//	}
//
//	go log.ReaderObj.Read(log.Rp)
//	go log.ProcessLog()
//	go log.WriterOjb.Write(log.Pw)
//
//	time.Sleep(1 * time.Second)
//	//aa := Student{
//	//	name: "Peter",
//	//	age: 18,
//	//}
//	//fmt.Println(aa)
//	//str := GetSqlitStr("1,2,4,5,6,8,a,b,c,d,e,10")
//	//
//	//fmt.Println(str)
//	//
//	//cache := cache2go.Cache("myCache")
//	//
//	//// We will put a new item in the cache. It will expire after
//	//// not being accessed via Value(key) for more than 5 seconds.
//	//val := myStruct{"This is a test!", []byte{}}
//	//cache.Add("someKey", 5*time.Second, &val)
//
//	//mycache.ClearCache()
//	//
//	//mycache.Add("peter", "pan", 10 * time.Second, Delete)
//	//
//	//time.Sleep(20 * time.Second)
//	//
//	//select {}
//
//	//t1 := time.Now()
//	//
//	//time.Sleep(time.Second)
//	//
//	//t2 := time.Now()
//	//
//	//fmt.Println(t2.Sub(t1))
//	//
//	//fmt.Println(t1)
//	//fmt.Println(t2)
//
//	//fmt.Println("aaa")
//	//
//	//dur := 2 * time.Second
//	//time.AfterFunc(dur, func(){
//	//	go print()
//	//})
//	//
//	//
//	//
//	//time.Sleep(5 * time.Second)
//	//fmt.Println("bbb")
//	//var ch chan int = make(chan int, 100)
//	//
//	//go loop(ch)
//	////loop()
//	////time.Sleep(time.Second)
//	//
//	//for i:=0;i<10;i++{
//	//	fmt.Printf("%d ",<-ch)
//	//	ch <- i
//	//}
//
//	// 开启5个routine
//	//for i := 0; i < 5; i++ {
//	//	go foo(i)
//	//}
//	//
//	//// 取出信道中的数据
//	//for i := 0; i < 5; i++ {
//	//	fmt.Print(<- ch)
//	//}
//
//	//异步任务
//	//router:=gin.Default()
//	////router.Use(MiddleWare())  //使用validate()中间件身份验证
//	//router.GET("/async", func(context *gin.Context) {
//	//	ccp := context.Copy()
//	//
//	//	go func(){
//	//		time.Sleep(5 * time.Second)
//	//		fmt.Println(ccp.Request.URL.Path)
//	//	}()
//	//	context.JSON(http.StatusOK, "ok")
//	//})
//	//
//	//router.GET("/sync", func(context *gin.Context) {
//	//	time.Sleep(5)
//	//	context.JSON(http.StatusOK, "ok")
//	//})
//	//
//	//router.GET("/",Test)
//	//router.GET("/test", MiddleWare(), Test2)
//	//router.Run(":9999")  //localhost:8989/
//	//
//
//}


func inArray(need interface{}, needArr []interface{}) bool {
	for _,v := range needArr{
		if need == v{
			return true
		}
	}
	return false
}

func Test2(ctx *gin.Context){
	ctx.JSON(http.StatusOK,"test2")
}

func Test(ctx *gin.Context){
	fmt.Println(ctx.Get("request"))
	ctx.JSON(http.StatusOK, "ok")
}

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before middleware")
		c.Set("request", "clinet_request")
		c.Next()
		fmt.Println("after middleware")
	}
}

func GetTimeByString(str string)(retTime time.Time, ok bool){
	retTime, err := time.Parse("2006-01-02", str)
	ok = err == nil

	return
}

var ch chan int = make(chan int)
func foo(id int) { //id: 这个routine的标号
	ch <- id
}


func loop(ch chan int){
	for i:=0;i<10;i++{
		fmt.Printf("%d ",i)
		ch <- i
	}
}

func Delete (item *mycache.CacheItem){
	fmt.Println(item)
}

func print(){
	fmt.Println("cccc")
}

func GetSqlitStr(str string)(splitStr string){
	tempList := strings.Split(str,",")
	var list []int
	for _,value := range tempList{
		id, err := strconv.Atoi(value)
		if err == nil {
			list = append(list, id)
		}
	}

	splitStr = strings.Replace(strings.Trim(fmt.Sprint(list), "[]"), " ", ",", -1)

	return
}
