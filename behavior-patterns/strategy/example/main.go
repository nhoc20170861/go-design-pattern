package main

func main() {
	// Initialize cache with LFU (Least Frequently Used) eviction algorithm
	lfu := &Lfu{}
	cache := initCache(lfu)

	// Add items to the cache
	cache.add("a", "1")
	cache.add("b", "2")

	// Add another item, which may trigger eviction based on LFU
	cache.add("c", "3")

	// Change eviction algorithm to LRU (Least Recently Used)
	lru := &Lru{}
	cache.setEvictionAlgo(lru)

	// Add another item, which may trigger eviction based on LRU
	cache.add("d", "4")

	// Change eviction algorithm to FIFO (First In First Out)
	fifo := &Fifo{}
	cache.setEvictionAlgo(fifo)

	// Add another item, which may trigger eviction based on FIFO
	cache.add("e", "5")
}