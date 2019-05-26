// Design a multi-threaded program for checking the Collatz conjecture for
// every integer in [1, U] where U is an input to your program. To keep your
// program from overloading the system, you should not have more than n threads
// running at a time.

package main

import (
	"fmt"
	"sync"
)

type CollatzChecker struct {
	number, cores int
	cache map[int]bool
	server chan int
	wg sync.WaitGroup
}

func (this *CollatzChecker) init() {
	this.cache = make(map[int]bool)
	this.server = make(chan int, this.number)

	for i := 0; i < this.cores; i++ {
		this.wg.Add(1)
		go func() {
			defer this.wg.Done()
			this.checker()
		}()
	}
}

func (this *CollatzChecker) checker() {
	for num := range this.server {
		computing := make(map[int]bool)
		this.compute(num, computing)
	}
}

func (this *CollatzChecker) compute(number int, computing map[int]bool) bool {
	if number == 1 {
		return true
	}

	// Maps are not safe for concurrent use. So, synchronization needed
	// when reading and writing for concurrent access.
	var safecache = struct {
		sync.RWMutex
		m map[int]bool
	}{m: this.cache}

	safecache.RLock()
	valid, computed := safecache.m[number]
	safecache.RUnlock()

	if computed { // Already computed
		return valid
	}

	valid := false
	if computing[number] { // Loop
		valid = false
	} else if number & 1 == 1 { // Odd
		valid = this.compute(number*3+1, computing)
	} else { // Even
		valid = this.compute(number >> 1, computing)
	}

	safecache.Lock()
	safecache.m[number] = valid
	safecache.Unlock()

	return valid
}

func (this *CollatzChecker) Check() {
	for i := 2; i <= this.number; i++ {
		this.server <- i
	}

	close(this.server)

	// Wait for all the go routines to exit.
	this.wg.Wait()
}

func (this *CollatzChecker) Print() {
	fmt.Println(this.cache)
}

func NewCollatzChecker(number, cores int) *CollatzChecker {
	this := &CollatzChecker{number: number, cores: cores}

	this.init()

	return this
}

func main() {
	resource := NewCollatzChecker(100, 8)
	resource.Check()
	resource.Print()
}
