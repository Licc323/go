package main

import (
	"fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup
var notifyCh = make(chan struct{}, 5)

func worker(id int, jobs <-chan int, result chan<- int)  {
	for j := range jobs {
		//fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		result <- j * 2
		notifyCh <- struct{}{}
	}
}

func main()  {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	//50个任务
	go func() {
		for j := 1; j <= 50; j++ {
			jobs <- j
		}
		close(jobs)
	}()
	// 开启3个 goroutine
	//wg.Add(3)
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	wg.Wait()

	go func() {
		for i := 0; i < 50; i++ {
			<-notifyCh
		}
		close(results)
	}()
	//输出结果
		for x := range results {
			fmt.Println(x)
		}
	}
