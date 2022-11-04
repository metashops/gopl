### 1、锁拷贝的问题：
* 拷贝可能导致锁的死锁问题
* 使用 vet 工具可以检查锁的拷贝问题，也可以检查可能的bug


```go
m := sync.Mutex
m.Lock()

// 执行业务逻辑

n := m  // 锁的复制
m.Unlock()

n.Lock() // n 一直是被锁住的




```

下面程序是否隐藏锁的问题？
```go
type Person struct {
	mu     sync.RWMutex
	salary int
	level  int
}

func (p *Person) promote() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.salary += 1000
	p.level += 1
	fmt.Printf("salary: %d, level: %d\n", p.salary, p.level)
}
func main() {
p := Person{salary: 1000, level: 1}
p2 := p // 这里相当拷贝整个 p 的所有字段，包括锁

```

隐秘的锁，不一定能发现，可以使用 Go 提供的 race
使用：go vet main.go

```shell
> go vet main.go                                                                               
# command-line-arguments
./main.go:32:8: assignment copies lock value to p2: command-line-arguments.Person contains sync.RWMutex
./main.go:33:14: call of fmt.Println copies lock value: command-line-arguments.Person contains sync.RWMutex
```

### 2、race 竞争检测

使用race命令：go build -race main.go
在执行，如win执行： ./main.exe