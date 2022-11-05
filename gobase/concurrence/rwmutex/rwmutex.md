
加写锁原理：
(1)在没有读协程的情况下：

     -----          ---------------
    |  W  |  --->  |   |   |   |   |
     -----          ---------------

     -----------             ---------------          
    | writerSem |    --->   |   |   |   |   |           
     -----------             ---------------

     -----------             ---------------          
    | readerSem |    --->   |   |   |   |   |           
     -----------             ---------------

   readerCount = - rwmutexMaxReaders
   readerWait  = 0

(2)有读协程的情况下:


     -----          ---------------
    |  W  |  --->  |   |   |   |   |
     -----          ---------------

     -----------             ---------------          
    | writerSem |    --->   | G |   |   |   |           
     -----------             ---------------

     -----------             ---------------          
    | readerSem |    --->   |   |   |   |   |           
     -----------             ---------------

   readerCount = 3 - rwmutexMaxReaders
   readerWait  = 3

总结：
* 先加mutex写锁，若已经被加写锁会阻塞等待
* 将readercount变为负数，阻塞读锁的获取
* 计算需要等待多少个读协程释放
* 如果需要等待读协程释放，陷入writerSem

          
### 解写锁        
* 将 readerCount 变为正值，允许读锁的获取


### 加读锁
 * 检查readerCount是否大于0，大于说明没有写锁干扰，那么直接将readerCount 加 1。（来一个读协程就+1）
 * 如果加读锁时readerCount小于0，说明被加了写锁干扰，陷入readerSem

### 解读锁
* 给readerCount 减 1
* 如readerCount正值，解锁成功
* 如readerCount负值，有写锁在排队；如自己时readerwait最后一个，唤醒写协程