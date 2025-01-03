package main

import "fmt"

type Cache struct {
	storage      map[string]string
	evictionAlgo EvictionAlgo
	capacity     int
	maxCapacity  int
}

func initCache(e EvictionAlgo) *Cache {
	return &Cache{
		storage:      make(map[string]string),
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *Cache) setEvictionAlgo(e EvictionAlgo) {
	c.evictionAlgo = e
}

func (c *Cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.storage[key] = value
	c.capacity++
	fmt.Println("Added key", key, "value", c.get(key), "capacity", c.capacity)
	fmt.Println("Storage", c.storage)
}

func (c *Cache) get(key string) string {
	return c.storage[key]
}

func (c *Cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}
