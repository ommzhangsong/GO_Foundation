package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Task struct {
	f func() error
}

func NewTask(f func() error) *Task {
	return &Task{f: f}
}

type Pool struct {
	RunningWorkers int64      //运行worker的数量
	Capacity       int64      //所有worker数量
	JobCH          chan *Task // 任务队列，放到chan中
	sync.Mutex                // 互斥锁，保护并发读写
}

func NewPool(capacity int64, tasksum int) *Pool {
	return &Pool{
		RunningWorkers: 0,
		Capacity:       capacity,
		JobCH:          make(chan *Task, tasksum),
	}
}

/*
GetCap
原子操作不是普通的变量读写，它会让 CPU 执行特定的内存屏障指令，这个指令会 “绕过缓存”，强制和主内存做数据同步。
也就是多协程的时候，可能读到旧值，不同的cpu核可能读到不同的值。
不是有mesi吗？MESI 协议确实是 CPU 缓存一致性协议，目标是让多核心的缓存数据保持一致，但它解决的是 “最终一致性”，而非 “即时一致性”。
MESI 的广播、失效标记、重新加载都是需要时间的（总线通信、缓存行状态切换），哪怕是纳秒级延迟，在并发场景下也会出问题
atomic.StoreInt64：强制 Core1 立刻广播失效通知，且等待其他核心确认 “缓存行已失效” 后，再继续执行后续指令；
atomic.LoadInt64：强制 Core2 先检查缓存行状态，若未同步则等待，直到拿到主内存最新值。
*/
func (p *Pool) GetCap() int64 {
	return atomic.LoadInt64(&p.Capacity) //是对p.Cap取地址，原子操作要的是地址
}
func (p *Pool) GetRunningWorkers() int64 {
	return atomic.LoadInt64(&p.RunningWorkers)
}
func (p *Pool) IncRunningWorkers() {
	atomic.AddInt64(&p.RunningWorkers, 1)
}
func (p *Pool) DecRunningWorkers() {
	atomic.AddInt64(&p.RunningWorkers, -1)
}
func (p *Pool) run() {
	// run 是既执行task。它又做了创建一个go协程，以及增加一个worker
	// 每一次执行run都会创建一个协程进去看jobch中有无可执行的任务，所以每一次
	// addtask的时候，就会创建一个协程
	p.IncRunningWorkers()
	go func() {
		defer func() {
			p.DecRunningWorkers()
		}()
		// 用一个协程，解决多个任务，减少协程的创建与销毁
		for task := range p.JobCH {
			err := task.f()
			if err != nil {
				return
			}
		}

	}()
}
func (p *Pool) AddTask(task *Task) {
	p.Lock()
	defer p.Unlock()
	if p.GetRunningWorkers() < p.Capacity {
		p.run()
	}
	p.JobCH <- task
}

// 入队顺序：task0 → task1 → task2 → ... → task9；
// 执行顺序可能是：go1 拿到 task0、go2 拿到 task1、go3 拿到 task2、go1 拿到 task3、go2 拿到 task4...（也可能是其他组合）。
func main() {
	pool := NewPool(3, 10)
	for i := 0; i < 10; i++ {
		pool.AddTask(NewTask(func() error {
			fmt.Printf("task %v\n", i)
			return nil
		}))
	}
	close(pool.JobCH)
	time.Sleep(1e9)
}

/*
✅ run() 的核心是创建 worker 协程（go1、go2、go3），这些协程是真正执行任务的主体；
✅ AddTask() 是非阻塞的（只要 JobCH 有缓冲且没满），所以循环 10 次能快速把 10 个 task 加入队列；
✅ 3 个 worker 协程会并发争抢队列中的 task，最终每个 worker 都会分到一部分 task 执行；
✅ 10 个 task 被加入队列的速度极快（几乎同时），因为 AddTask() 只是往缓冲通道里写数据，耗时微秒级。
*/
