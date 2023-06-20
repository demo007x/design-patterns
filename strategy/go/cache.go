package main

type Cache struct {
	storage      map[string]string
	evictionAlgo EvictionAlgo
	capacity     int
	macCapacity  int
}

// 初始化的时候将策略注入到 cache 中
func initCache(e EvictionAlgo) *Cache {
	storage := make(map[string]string)
	return &Cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		macCapacity:  2,
	}
}

// 动态修改策略
func (c *Cache) setEvictionAlgo(e EvictionAlgo) {
	c.evictionAlgo = e
}

func (c *Cache) add(key, value string) {
	// 如果缓存中的容量等于了最大容量，则需要执行策略来移除 s
	if c.capacity == c.macCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *Cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

func (c *Cache) get(key string) {
	delete(c.storage, key)
}
