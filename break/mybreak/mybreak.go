package mybreak

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
	"sync"
	"time"
)

type mybreak struct {
	mutex sync.Mutex
	errCount int
	Count int
	setting MySetting
	status int
	closeTime time.Time
	openTime time.Time
}

func (b *mybreak)clear(){
	b.errCount = 0
	b.status = 0
}

func (b *mybreak) beforeRequest(now time.Time) error{
	b.Count ++
	if b.status == 0{
		return nil
	}else{
		if now.After(b.openTime){
			fmt.Println("reopen")
			b.clear()
			return nil
		}
		return errors.New("Max error count is : " + strconv.Itoa(b.errCount))
	}

	return nil
}

func (b *mybreak) afterRequest(now time.Time){
	if b.errCount >= b.setting.MaxError{
		b.status = 1
		b.closeTime = now
		m, _ := time.ParseDuration(strconv.Itoa(b.setting.Inv)+"s")
		b.openTime = now.Add(m)
	}
}

func (b *mybreak) Excute(req func()(interface{}, error)) (interface{}, error){
	b.mutex.Lock()
	now := time.Now()

	defer b.mutex.Unlock()

	ok := b.beforeRequest(now)
	if ok != nil{
		return nil, ok
	}
	v,err:= req()
	if err == nil {
		return v, nil
	}else{
		b.errCount ++
		b.afterRequest(now)
		return v, err
	}
}

type MySetting struct {
	MaxError int
	Inv int
}

func NewMyBreak (setting MySetting) *mybreak {
	b := new (mybreak)
	b.setting = setting

	return b
}