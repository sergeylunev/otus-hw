package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	Cache // Remove me after realization.

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
	} else {
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
		lc.queue.PushFront(ci)
		li := lc.queue.Front()
		lc.items[key] = li

		return false
	}
}

func (lc *lruCache) Get(key Key) (interface{}, bool) {
	if v, ok := lc.items[key]; !ok {
		return nil, false
	} else {
		val := v.Value.(cacheItem)
		lc.queue.MoveToFront(v)

		return val.value, ok
	}
}

func (lc *lruCache) Clear() {
	for k, v := range lc.items {
		lc.queue.Remove(v)
		delete(lc.items, k)
	}
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
