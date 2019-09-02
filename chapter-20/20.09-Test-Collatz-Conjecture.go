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
	server chan int
	wg sync.WaitGroup

	// Maps are not safe for concurrent use. So, synchronization needed
	// when reading and writing for concurrent access.
	safecache struct {
		sync.RWMutex
		m map[int]bool
	}
}

func (this *CollatzChecker) init() {
	this.server = make(chan int, this.number)
	this.safecache.m = make(map[int]bool)

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

	this.safecache.RLock()
	valid, computed := this.safecache.m[number]
	this.safecache.RUnlock()

	if computed { // Already computed
		return valid
	}

	valid = false
	if computing[number] { // Loop
		valid = false
	} else if number & 1 == 1 { // Odd
		computing[number] = true
		valid = this.compute(number*3+1, computing)
	} else { // Even
		computing[number] = true
		valid = this.compute(number >> 1, computing)
	}

	this.safecache.Lock()
	this.safecache.m[number] = valid
	this.safecache.Unlock()

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
	fmt.Println(this.safecache.m)
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
