package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//原子操作
var x int64
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	//x++
	//lock.Lock()
	//wg.Done()
	//x = x + 1
	atomic.AddInt64(&x, 1)
	//lock.Unlock()
	wg.Done()
}

func main()  {
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go add()
	}
	wg.Wait()
	fmt.Println(x)
	//x = 200
	//ok := atomic.CompareAndSwapInt64(&x, 100, 200)
	//fmt.Println(ok, x)
}
