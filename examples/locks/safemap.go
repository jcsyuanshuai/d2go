package locks

import "sync"

type SafeMap struct {
	//mux *sync.Mutex    //`sync.Mutex`和`sync.ReMutex`默认值是有效的，所以指向其指针是不必要的
	// sync.Mutex       //不应该将其直接嵌入到结构体中，因为这样会暴露`Mutex`的实现
	mux sync.Mutex

	data map[string]string
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]string),
	}
}

func (m *SafeMap) Get(k string) string {
	m.mux.Lock()
	defer m.mux.Unlock()

	return m.data[k]
}

func (m *SafeMap) Put(k, v string) {
	m.mux.Lock()
	defer m.mux.Unlock() // 使用defer释放锁，defer资源开销小并且似的代码可读性增大

	m.data[k] = v
}
