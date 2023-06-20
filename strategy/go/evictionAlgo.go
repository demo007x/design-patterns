package main

// 策略接口
type EvictionAlgo interface {
	evict(c *Cache)
}
