
需求：实际业务中，一个组协程需要等待另一组协程完成
![](https://tva1.sinaimg.cn/large/008vxvgGgy1h7sdxxu26vj31ce0hotaf.jpg)

             WaitGroup
             ---------
             | waiter|
             ---------
             |counter|
             ---------
             | sema  |    -------> G ...
             ---------

* waiter 记录有多少个协程在等待运行
* counter 记录有多少个协程正在运行
* sema 记录等待的协程

（1）Wait()
* 如果被等待的协程没有了，直接返回
* 否则，wait 加 1，陷入 sema

（2）Done()
* 被等待协程做完，给 counter -1
* 通过 Add(-1) 实现的

（3）Add
* 当counter减到0，sema 释放出来