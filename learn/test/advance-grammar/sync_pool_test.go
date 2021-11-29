package advancegrammar_test

import (
	"runtime"
	"sync"
	"testing"
)

/** sync.Pool对象的获取
1. 尝试从私有对象中获取(不是池,仅能缓存一个对象, 协程安全, 所以不需要锁)
2. 私有对象不存在, 从当前Processor共享池中获取
3. 共享池也是空的, 从其他Processor的共享池中获取
4. 所有共享池都是空的, 从指定的New函数中创建对象
*/

/** sycn.Pool的放回
1.先放会到私有对象中
2.私有对象存在, 放在当前Processor共享池
*/

/** sync.Pool声明周期
1.GC会清楚sync.Pool的对象
2.对象缓存有效期是下一次GC之前
*/

/** syncPool使用场景
1.通过复用降低复杂对象的创建和GC代价
2.协程安全,会有锁开销
3.声明周期守GC影响, 不适合做连接池,需要自己管理生命周期的资源池优化
*/

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			t.Logf("create a new object\n")
			return 0
		},
	}
	v := pool.Get().(int)
	t.Logf("get value: %v", v)
	age := 32
	pool.Put(&age)
	// 测试GC情况sync.Pool对象
	runtime.GC()
	v1 := pool.Get().(int)
	t.Logf("get new value: %v", v1)
}

func TestSyncGroupPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			t.Logf("create a new object\n")
			return 100
		},
	}
	age := 10
	pool.Put(age)
	pool.Put(age)
	pool.Put(age)

	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			v := pool.Get().(int)
			t.Logf("get value: %v", v)
		}()
	}
	wg.Wait()
}
