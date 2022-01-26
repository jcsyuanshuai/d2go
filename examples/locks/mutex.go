package locks

import "sync"

func bad() {
	mu := new(sync.Mutex)
	mu.Lock()
	defer mu.Unlock()
	//...
}

func good() {
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()
	//...
}

type SMap struct {
	//mu *sync.Mutex
	// sync.Mutex
	mu sync.Mutex

	data map[string]string
}

func NewSMap() *SMap {
	return &SMap{
		data: make(map[string]string),
	}
}

func (m *SMap) Get(k string) string {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.data[k]
}
