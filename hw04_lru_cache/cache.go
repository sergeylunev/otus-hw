package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (lc *lruCache) Set(key Key, value interface{}) bool {
	if v, ok := lc.items[key]; ok {
		v.Value = cacheItem{
			key:   key,
			value: value,
		}
		return true
	}

	if len(lc.items) == lc.capacity {
		back := lc.queue.Back()
		backKey := back.Value.(cacheItem).key
		delete(lc.items, backKey)
		lc.queue.Remove(back)
	}

	ci := cacheItem{
		key:   key,
		value: value,
	}
	lc.items[key] = lc.queue.PushFront(ci)

	return false
}

func (lc *lruCache) Get(key Key) (interface{}, bool) {
	v, ok := lc.items[key]

	if !ok {
		return nil, false
	}

	val := v.Value.(cacheItem)
	lc.queue.MoveToFront(v)

	return val.value, ok
}

func (lc *lruCache) Clear() {
	lc.items = make(map[Key]*ListItem, lc.capacity)
	lc.queue = NewList()
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
