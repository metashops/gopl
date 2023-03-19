package main

import (
	"fmt"
	"sync"
	"time"
)

func job1(count int) <-chan int {
	outCh := make(chan int, 2)

	go func() {
		defer close(outCh)
		for i := 0; i < count; i++ {
			time.Sleep(time.Second)
			fmt.Println("job1 finish:", 1)
			outCh <- 1
		}
	}()

	return outCh
}

func job2(inCh <-chan int) <-chan int {
	outCh := make(chan int, 2)

	go func() {
		defer close(outCh)
		for val := range inCh {
			// 耗时2秒
			time.Sleep(time.Second * 2)
			val++
			fmt.Println("job2 finish:", val)
			outCh <- val
		}
	}()

	return outCh
}

func job3(inCh <-chan int) <-chan int {
	outCh := make(chan int, 2)

	go func() {
		defer close(outCh)
		for val := range inCh {
			val++
			fmt.Println("job3 finish:", val)
			outCh <- val
		}
	}()

	return outCh
}

func merge(inCh ...<-chan int) <-chan int {
	outCh := make(chan int, 2)

	var wg sync.WaitGroup
	for _, ch := range inCh {
		wg.Add(1)
		go func(wg *sync.WaitGroup, in <-chan int) {
			defer wg.Done()
			for val := range in {
				outCh <- val
			}
		}(&wg, ch)
	}

	// 重要注意，wg.Wait() 一定要在goroutine里运行，否则会引起deadlock
	go func() {
		wg.Wait()
		close(outCh)
	}()

	return outCh
}

func main() {
	t := time.Now()

	firstResult := job1(10)

	// 拆分成三个job2，即3个goroutine （扇出）
	secondResult1 := job2(firstResult)
	secondResult2 := job2(firstResult)
	secondResult3 := job2(firstResult)
	secondResult4 := job2(firstResult)
	secondResult5 := job2(firstResult)

	// 合并结果(扇入）
	secondResult := merge(secondResult1, secondResult2, secondResult3, secondResult4, secondResult5)

	thirdResult := job3(secondResult)

	for v := range thirdResult {
		fmt.Println(v)
	}

	fmt.Println("all finish")
	fmt.Println("duration:", time.Since(t).String())
}
