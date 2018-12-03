package mycache

import (
	"fmt"
	"time"
)
import "sync"

var cacheMap map[string]CacheItem = make(map[string]CacheItem)
var lock sync.RWMutex

type CacheItem struct {
	value interface{}
	key string
	time time.Time
	exprited time.Time
	expritedFun func(*CacheItem)
}

func Add(key string, value interface{}, duration time.Duration, exfun func(*CacheItem)){
	lock.Lock()

	defer lock.Unlock()

	item := CacheItem{}

	item.value = value
	item.key = key
	item.time = time.Now()
	item.exprited = time.Now().Add(duration)
	item.expritedFun = exfun
	cacheMap[key] = item
}

func Del(key string){
	lock.Lock()
	defer lock.Unlock()

	delete(cacheMap, key)
}

func ClearCache(){
	lock.Lock()
	defer lock.Unlock()
	small := 1 * time.Second

	for k,v := range cacheMap{
		temp := time.Now().Sub(v.exprited)
		fmt.Println(temp)
		if temp > 0{
			delete(cacheMap, k)
			if v.expritedFun != nil {
				v.expritedFun(&v)
			}
		}else{
			if small == 0 || temp < small {
				small = temp
			}
		}
	}

	time.AfterFunc(small, func(){
		go ClearCache()
	})
}
