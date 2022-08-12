package lru

type LCache struct {
	Cap   int
	Len   int
	cache map[int]*LinkedNode
	list  *LinkedList
}

func NewLCache(cap int) LCache {
	return LCache{
		Cap:   cap,
		cache: make(map[int]*LinkedNode, cap),
		list:  &LinkedList{},
	}
}

func (l *LCache) Get(key int) int {
	if e, ok := l.cache[key]; ok {
		l.list.MoveToFront(e)
		return e.Value.val
	}
	return pair{}.val
}

func (l *LCache) Put(key, val int) int {
	// 添加新数据时，判断数据是否在cache中存在，存在则更新
	if e, ok := l.cache[key]; ok {
		l.list.MoveToFront(e)
		oldVal := e.Value.val
		e.Value.val = val
		return oldVal
	}
	// 不存在时，判断cache长度是否超出容量，是则删除
	// 不仅删除cache，也要删除列表中
	if l.Len >= l.Cap {
		delete(l.cache, l.list.Back().Value.key)
		l.list.Remove(l.list.Back())
	}
	// 将新数据添加到头节点，并且添加到cache
	l.list.PushToFront(pair{key: key, val: val})
	l.cache[key] = l.list.Front()
	return pair{}.val
}
