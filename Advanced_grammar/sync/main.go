package main

import (
	"fmt"
	"sync"
)

var num int

// ❗init 函数自动执行，程序一启动就执行，也是执行一次，但是控制不了我们执行的时间
// 也一样，不管包导入多少次，就放内存里了，只会执行一次
// 初始化包级变量，然后执行init后执行main，可以有多个init，包被导入的时候（import）它就执行了
func init() {
	fmt.Println("我是init函数")

}

var wg sync.WaitGroup

/*
sync once是单例模式里的懒汉模式，也就是延迟初始化，只有在调用do的时候，它才初始化创建
once变量不是延迟，延迟的是oncedo后面的函数。
延迟是相对于这个开局就声明创建完成而说的。正常函数也能做到延迟，但做不到只执行一次
因为go有多线程，所以我们只需要初始化一次
*/
var once sync.Once
var mu sync.Mutex

// 加mutex互斥锁，所有操作都是写，用互斥锁否则就用读写锁
func goroutine(i int) {
	defer wg.Done()
	mu.Lock()
	num = num + 1
	mu.Unlock()
}

type Config struct {
	app string
}

var instance *Config

// InitConfig 配置文件的加载

func InitConfig() *Config {
	once.Do(func() {
		instance = &Config{"apple"}
	})
	return instance
}
func main() {
	wg.Add(1000)
	//如果最后add多了，会死锁
	for i := range 1000 {
		go goroutine(i)
	}
	wg.Wait()
	fmt.Println(num)
}

/*
❗带锁结构的变量赋值时，锁的状态也会赋值，所以复制锁的时候，lock的锁任然lock
❗sunc.Mapsync.Map 的设计和普通 map 不一样，它的 key 和 value 类型都是 interface{}（空接口）
*/
