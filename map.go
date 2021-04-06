package algorithm

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
实现阻塞读且并发安全的map
思路
1、阻塞使用chan
2、并发安全使用sync.RWMutex
*/

type MyMap interface {
	Set(key string, val interface{})                   //存入key /val，如果该key读取的goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回
	Get(key string, timeout time.Duration) interface{} //读取一个key，如果key不存在阻塞，等待key存在或者超时
}

type Key struct {
	c       chan bool // 用于唤醒获取该key阻塞的goroutine
	value   interface{}
	isExist bool // true表示已存在，false表示未存在，有goroutine阻塞
}

type Map struct {
	m   map[string]*Key
	RWm sync.RWMutex
}

func NewMap() *Map {
	return &Map{
		m:   make(map[string]*Key),
		RWm: sync.RWMutex{},
	}
}

func (m *Map) Set(key string, val interface{}) {
	m.RWm.Lock()
	defer m.RWm.Unlock()

	v, ok := m.m[key]
	if !ok || v.isExist { // 该key不存在于map中
		tempKey := Key{
			value:   val,
			isExist: true,
		}
		m.m[key] = &tempKey
		return
	}

	if !v.isExist { // 该key已经存在，有goroutine阻塞
		v.isExist = true
		v.value = val
		close(v.c)
	}
}

func (m *Map) Get(key string, t time.Duration) interface{} {
	m.RWm.RLock()
	defer m.RWm.RUnlock()

	val, ok := m.m[key]
	if ok && val.isExist {
		return val.value
	}

	if !ok { // 不存在该key，阻塞或超时
		tempKey := Key{
			c:     make(chan bool),
			value: val,
		}
		m.m[key] = &tempKey

		go func(t time.Duration) {
			time.Sleep(t)
			close(tempKey.c)
		}(t)

		<-tempKey.c
		if tempKey.isExist {
			return val.value
		}
		return nil
	}

	return nil
}

// ==================
/*
场景：在一个高并发的web服务器中，要限制IP的频繁访问。现模拟100个IP同时并发访问服务器，每个IP要重复访问1000次。
每个IP一分钟之内只能访问一次。修改以下代码完成该过程，要求能成功输出 success:100
思路
1、子线程：（定时器）定期检查map，删除超时的key，WithCancel避免孤儿线程
2、go中map并发不安全，需要加锁
3、atomic.AddInt64对临界值操作
*/

type Ban struct {
	visitIPs map[string]time.Time
	lock     sync.Mutex
}

func NewBan(ctx context.Context) *Ban {
	o := &Ban{visitIPs: make(map[string]time.Time)}
	go func() {
		timer := time.NewTimer(time.Minute * 1)
		for {
			select {
			case <-timer.C: // 定时器
				o.lock.Lock()
				for k, v := range o.visitIPs {
					if time.Since(v) >= time.Minute*1 {
						delete(o.visitIPs, k)
					}
				}
				o.lock.Unlock()
				timer.Reset(time.Minute * 1)
			case <-ctx.Done(): // 子线程退出
				return
			}
		}
	}()
	return o
}
func (o *Ban) visit(ip string) bool {
	o.lock.Lock()
	defer o.lock.Unlock()
	if _, ok := o.visitIPs[ip]; ok {
		return true
	}
	o.visitIPs[ip] = time.Now()
	return false
}
func Filter() {
	success := int64(0)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ban := NewBan(ctx)

	wait := &sync.WaitGroup{}

	wait.Add(1000 * 100)
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			go func(j int) {
				defer wait.Done()
				ip := fmt.Sprintf("192.168.1.%d", j)
				if !ban.visit(ip) {
					atomic.AddInt64(&success, 1)
				}
			}(j)
		}

	}
	wait.Wait()

	fmt.Println("success:", success)
}
