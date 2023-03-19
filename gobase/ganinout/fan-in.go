package main

import (
	`fmt`
	`net/http`
	`sync`
	`time`

	`github.com/gin-gonic/gin`
)

func fanIn(inCh ...<-chan interface{}) <-chan interface{} {
	outCh := make(chan interface{})

	var wg sync.WaitGroup
	for _, ch := range inCh {
		wg.Add(1)
		go func(wg *sync.WaitGroup, in <-chan interface{}) {
			defer wg.Done()
			for val := range in {
				outCh <- val
			}
		}(&wg, ch)
	}

	// 注：wg.Wait() 一定要在goroutine里运行，否则会引起deadlock
	go func() {
		wg.Wait()
		defer close(outCh)
	}()

	return outCh
}

func task1() <-chan interface{} {
	outCh := make(chan interface{})

	go func() {
		defer close(outCh)
		// 假设这个任务很耗时，必须做完才返回
		time.Sleep(time.Second * 2)
		// 任务完成后，往 chan
		outCh <- "Task 1 Finish<br>"
	}()
	return outCh
}

func task2(inCh <-chan interface{}) <-chan interface{} {
	outCh := make(chan interface{})

	go func() {
		defer close(outCh)
		// 耗时5秒
		time.Sleep(time.Second * 5)
		for val := range inCh {
			fmt.Println(val)
			outCh <- "Task 2 Finish<br>"
		}
	}()
	return outCh
}

func main() {
	router := gin.Default()
	router.GET("/show", func(c *gin.Context) {
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusOK)

		ret := task1()
		split1 := task2(ret)
		split2 := task2(ret)
		split3 := task2(ret)
		res := fanIn(split1, split2, split3)
		for val := range res {
			c.Writer.Write([]byte(val.(string)))
			c.Writer.(http.Flusher).Flush()
		}
	})

	router.Run(":8080")
}
