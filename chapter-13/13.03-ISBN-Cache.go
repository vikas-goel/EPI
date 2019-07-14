// Create a cache for looking up prices of books identified by their ISBN. You
// implement lookup, insert, and remove methods. Use the Least Recently Used
// (LRU) policy for cache eviction. If an ISBN is already present, insert
// should not change the price, but it should update that entry to be the most
// recently used entry. Lookup should also update that entry to be the most
// recently used entry.

package main

import (
	"fmt"
	"container/list"
	"math/rand"
)

type LRUCache struct {
	size int
	queue list.List
	hash map[int]*list.Element
}

func (this *LRUCache) Lookup(key int) (value int, exists bool) {
	if element, ok := this.hash[key]; ok {
		this.queue.MoveToFront(element)
		value, exists = element.Value.(int), true
	}

	return
}

func (this *LRUCache) Insert(key, value int) int {
	if element, ok := this.hash[key]; ok {
		this.queue.MoveToFront(element)
		return element.Value.(int)
	}

	if this.queue.Len() == this.size {
		last := this.queue.Back()
		this.queue.Remove(last)
		delete(this.hash, last.Value.(int))

		this.size--
	}

	front := this.queue.PushFront(value)
	this.hash[key] = front
	this.size++

	return value
}

func (this *LRUCache) Remove(key int) (existed bool) {
	if element, ok := this.hash[key]; ok {
		this.queue.Remove(element)
		delete(this.hash, key)
		this.size--

		existed = true
	}

	return
}

func NewLRUCache(size int) *LRUCache {
	var cache LRUCache
	cache.size = size
	cache.hash = make(map[int]*list.Element)
	return &cache
}

func main() {
	cache := NewLRUCache(5)

	for i := 0; i < 100; i++ {
		operation := rand.Intn(2)
		isbn := rand.Intn(50) + 1
		if operation == 1 {
			price := rand.Intn(1000) + 1
			cache.Insert(isbn, price)
			fmt.Printf("(insert[%v, %v]) ", isbn, price)
		} else {
			value, exists := cache.Lookup(isbn)
			if exists {
				fmt.Printf("(lookup[%v], %v) ", isbn, value)

				existed := cache.Remove(isbn)
				fmt.Printf("(remove[%v], %v) ", isbn, existed)

				value, exists = cache.Lookup(isbn)
				fmt.Printf("(lookup[%v], %v) ", isbn, exists)
			} else {
				fmt.Printf("(lookup[%v], miss) ", isbn)
			}

		}
	}
}
